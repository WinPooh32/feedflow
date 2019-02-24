package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/WinPooh32/feedflow/api"
	"github.com/WinPooh32/feedflow/model"
	"github.com/WinPooh32/feedflow/web"

	gintemplate "github.com/foolin/gin-template"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type settings struct {
	Port *string
	Host *string

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

	s.DbHost = flag.String("dbhost", "localhost", "listening database server ip")
	s.DbPort = flag.String("dbport", "5432", "listening database port")
	s.DbDriver = flag.String("dbdriver", "postgres", "The database diver")
	s.DbUser = flag.String("dbuser", "", "The database username")
	s.DbPassword = flag.String("dbpass", "", "The database user password")
	s.DbName = flag.String("dbname", "", "The database name")
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
	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  listPartials("views", "partials", ".html"),
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})
}

func routeStatic(router *gin.Engine, prefix string) {
	router.Static(prefix, "./assets")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/bundle.js", "./frontend/dist/bundle.js")
}

func initRouter(router *gin.Engine, svSettings settings, debug bool) *gin.Engine {
	db, err := initDatabse(svSettings, debug)

	if err != nil {
		log.Println("Database error:", err)
	} else {
		router.Use(databaseMiddleware(db))
		model.MigrateModels(db)
	}

	if debug {
		router.Use(cors.Default())
	}

	initTemplateManager(router)
	routeStatic(router, "/assets")
	web.RouteWeb(router)
	api.RouteAPI(router)

	return router
}

func main() {
	debug := gin.Mode() == gin.DebugMode
	svSettings := readSettings()

	//Make new gin router
	router := initRouter(gin.Default(), svSettings, debug)

	listenAt := fmt.Sprintf("%s:%s", *svSettings.Host, *svSettings.Port)
	if err := endless.ListenAndServe(listenAt, router); err != nil {
		log.Fatalln(err)
	}
}
