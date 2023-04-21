package main

import (
	"Test2/initializers"
	"Test2/pkg/routes"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func init() {
	//initializers.LoadEnvVar()
	initializers.Connect()
	initializers.SyncDB()
}

func main() {
	router := gin.Default()
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views/templates",
		Extension: ".html",
	})
	routes.Router(router)
}
