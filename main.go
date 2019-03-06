package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"html/template"
	"log"
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

	gintemplate "github.com/WinPooh32/gin-template"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/go-session/redis"
	"github.com/go-session/session"
)

type settings struct {
	Port *string
	Host *string
	Ssl  *string

	DbHost     *string
	DbPort     *string
	DbDriver   *string
	DbUser     *string
	DbPassword *string
	DbName     *string
	DbSsl      *bool
}

func readSettings() settings {
	var s settings

	s.Host = flag.String("host", "localhost", "listening server ip")
	s.Port = flag.String("port", "8080", "listening port")
	s.Ssl = flag.String("ssl", "", "cert;private")

	s.DbHost = flag.String("dbhost", "localhost", "listening database server ip")
	s.DbPort = flag.String("dbport", "5432", "listening database port")
	s.DbDriver = flag.String("dbdriver", "sqlite3", "The database diver")
	s.DbUser = flag.String("dbuser", "", "The database username")
	s.DbPassword = flag.String("dbpass", "", "The database user password")
	s.DbName = flag.String("dbname", "storage", "The database name")
	s.DbSsl = flag.Bool("dbssl", false, "The database ssl enabled")

	flag.Parse()
	return s
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

func initRouter(router *gin.Engine, svSettings settings, debug bool) (*gin.Engine, func()) {
	log.Println("Initialize gin router...")

	db, err := database.Init(database.Credential{
		Driver:   *svSettings.DbDriver,
		Host:     *svSettings.DbHost,
		Port:     *svSettings.DbPort,
		User:     *svSettings.DbUser,
		Database: *svSettings.DbName,
		Password: *svSettings.DbPassword,
		Ssl:      *svSettings.DbSsl,
	}, debug)

	//setup middlewares
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

	if debug {
		router.Use(cors.Default())
	}

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

func main() {
	debug := gin.Mode() == gin.DebugMode
	svSettings := readSettings()

	//Make new gin router
	router, onShutdown := initRouter(gin.Default(), svSettings, debug)

	listenAt := fmt.Sprintf("%s:%s", *svSettings.Host, *svSettings.Port)

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
	if len(*svSettings.Ssl) == 0 {
		err = srv.ListenAndServe()
	} else {
		tlsConf := &tls.Config{}
		tlsConf.NextProtos = []string{"h2", "http/1.1"}
		srv.Server.TLSConfig = tlsConf

		keys := strings.Split(*svSettings.Ssl, ";")
		err = srv.ListenAndServeTLS(keys[0], keys[1])
	}

	if err != nil {
		log.Fatalln(err)
	}
}
