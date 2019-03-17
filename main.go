package main

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/WinPooh32/feedflow/api"
	"github.com/WinPooh32/feedflow/database"
	"github.com/WinPooh32/feedflow/model"
	"github.com/WinPooh32/feedflow/web"
	"github.com/jessevdk/go-flags"

	gintemplate "github.com/WinPooh32/gin-template"
	"github.com/WinPooh32/gzip"
	"github.com/cnjack/throttle"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/go-session/redis"
	"github.com/go-session/session"
)

type options struct {
	Config string `short:"c" long:"config" default:"" no-ini:"true"`

	Verbose bool `short:"v" long:"verbose" deafault:"false"`

	Limit       uint64 `short:"l" long:"limit"         default:"500" description:"Request throttle"`
	LimitWithin uint64 `          long:"limit-within"  default:"5"`

	Port string `short:"p"   long:"port"  default:"8080"       description:"listening port"`
	Host string `short:"h"   long:"host"  default:"localhost"  description:"listening server ip"`
	Ssl  string `            long:"ssl"   default:""           description:"cert;private"`
	Gzip bool   `            long:"gzip"                       description:"gzip compression"`

	DbHost     string `long:"dbhost"   default:"localhost" description:"listening database server ip"`
	DbPort     string `long:"dbport"   default:"5432"      description:"listening database port"`
	DbDriver   string `long:"dbdriver" default:"sqlite3"   description:"The database diver"`
	DbUser     string `long:"dbuser"   default:""          description:"The database username"`
	DbPassword string `long:"dbpass"   default:""          description:"The database user password"`
	DbName     string `long:"dbname"   default:"storage"   description:"The database name"`
	DbSsl      bool   `long:"dbssl"                        description:"The database ssl enabled"`
}

func readSettings() options {
	opts := options{}
	parser := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash)

	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		log.Fatalln(err)
	}

	if len(opts.Config) > 0 {
		// Parse an ini file
		iniParser := flags.NewIniParser(parser)
		if err := iniParser.ParseFile(opts.Config); err != nil {
			log.Println("error parsing ini file: ", opts.Config)
			log.Fatalln(err)
		}
	}

	return opts
}

func listPartials(viewsPath, partialsPath, fileExtension string) []string {
	partials := make([]string, 0)
	walkPath := viewsPath + "/" + partialsPath

	err := filepath.Walk(walkPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				trimmed := strings.TrimLeft(path, viewsPath+"/")
				trimmed = strings.TrimSuffix(trimmed, fileExtension)

				partials = append(partials, trimmed)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	log.Println("Listed partials:", partials)
	return partials
}

func initTemplateManager(router *gin.Engine) {

	config := gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  listPartials("views", "partials", ".html"),
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: false,
	}

	//new template engine
	router.HTMLRender = gintemplate.New(config)
}

func routeStatic(router *gin.Engine, prefix string) {
	router.Static(prefix, "./assets")
	router.Static("js", "./frontend/dist/js")
	router.Static("css", "./frontend/dist/css")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
}

func initGoSession() (store session.ManagerStore) {
	store = redis.NewRedisStore(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       15,
		PoolSize: runtime.NumCPU(),
	})

	// store = gormstore.MustStoreWithDB(db, "go-session", 600)

	return store
}

func initRouter(router *gin.Engine, opts options) (*gin.Engine, func()) {
	log.Println("Initialize gin router...")

	verbose := opts.Verbose

	db, err := database.Init(database.Credential{
		Driver:   opts.DbDriver,
		Host:     opts.DbHost,
		Port:     opts.DbPort,
		User:     opts.DbUser,
		Database: opts.DbName,
		Password: opts.DbPassword,
		Ssl:      opts.DbSsl,
	}, verbose)

	//setup middlewares
	router.Use(gin.Recovery())

	if verbose {
		router.Use(gin.Logger())
	}

	router.Use(cors.Default()) //FIXME for debug purpose

	router.Use(throttle.Policy(&throttle.Quota{
		Limit:  opts.Limit,
		Within: time.Second * time.Duration(opts.LimitWithin),
	}))

	if opts.Gzip {
		router.Use(gzip.Gzip(gzip.BestSpeed))
	}

	if err != nil {
		log.Println("Database error:", err)
	} else {
		router.Use(database.NewMiddleware(db))
		model.MigrateModels(db)
	}

	sessionStore := initGoSession()
	sessionExpireOpt := session.SetExpired(24 * 60 * 60) // 24 hours
	sessionStoreOpt := session.SetStore(sessionStore)
	router.Use(ginsession.New(sessionStoreOpt, sessionExpireOpt))

	//setup templates
	initTemplateManager(router)

	//setup routes
	routeStatic(router, "/assets")
	web.RouteWeb(router)
	api.RouteAPI(router)

	return router, func() {
		recovery := func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in shutdown", r)
			}
		}

		log.Println("Server shutdown!")

		func() {
			defer recovery()
			log.Println("Session storage shutdown!")
			if err := sessionStore.Close(); err != nil {
				log.Println("Session storage closing error:", err)
			}
		}()

		func() {
			defer recovery()
			log.Println("Databse shutdown!")
			if err := db.Close(); err != nil {
				log.Println("Databas closing error:", err)
			}
		}()
	}
}

func httpsRedirect() {
	redirect := func(w http.ResponseWriter, req *http.Request) {
		// remove/add not default ports from req.Host
		target := "https://" + req.Host + req.URL.Path
		if len(req.URL.RawQuery) > 0 {
			target += "?" + req.URL.RawQuery
		}
		// log.Printf("redirect to: %s", target)
		http.Redirect(w, req, target,
			// see @andreiavrammsd comment: often 307 > 301
			http.StatusTemporaryRedirect)
	}

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
}

func main() {
	defer os.Exit(0)

	if err := writePidFile("feedflow.pid"); err != nil {
		log.Fatalln(err)
	}

	opts := readSettings()

	if opts.Verbose {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	//Make new gin router
	router, onShutdown := initRouter(gin.New(), opts)

	listenAt := fmt.Sprintf("%s:%s", opts.Host, opts.Port)

	srv := endless.NewServer(listenAt, router)

	//Set hook for all signals
	hookableSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGTSTP,
	}

	for _, sig := range hookableSignals {
		srv.SignalHooks[endless.PRE_SIGNAL][sig] = append(srv.SignalHooks[endless.PRE_SIGNAL][sig], onShutdown)
	}

	//Start the http server
	var err error
	if len(opts.Ssl) == 0 {
		err = srv.ListenAndServe()
	} else {
		//Redirect http to https
		httpsRedirect()

		tlsConf := &tls.Config{}
		tlsConf.NextProtos = []string{"h2", "http/1.1"}
		srv.Server.TLSConfig = tlsConf

		keys := strings.Split(opts.Ssl, ";")
		err = srv.ListenAndServeTLS(keys[0], keys[1])
	}

	if err != nil {
		log.Fatalln(err)
	}
}
