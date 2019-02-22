package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"time"

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

func initTemplateManager(router *gin.Engine) {
	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})
}

func main() {
	debug := gin.Mode() == gin.DebugMode

	svSettings := readSettings()

	//Make new gin router
	router := gin.Default()

	db, err := initDatabse(svSettings, debug)
	if err != nil {
		log.Println("Database error:", err)
	} else {
		router.Use(databaseMiddleware(db))
		migrateModels(db)
	}

	if debug {
		router.Use(cors.Default())
	}

	RouteAPI(router)

	listenAt := fmt.Sprintf("%s:%s", *svSettings.Host, *svSettings.Port)
	if err := endless.ListenAndServe(listenAt, router); err != nil {
		log.Fatalln(err)
	}
}
