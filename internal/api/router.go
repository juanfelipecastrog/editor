package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gold/config"
	"gold/internal/controller"
	"gold/internal/html"
	"gold/internal/static"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	static.LoadStatic(router)
	html.LoadHtml(router)

	server := &Server{
		router: router,
	}

	server.registerRoutes()

	return server
}

func (s *Server) registerRoutes() {
	articleController := controller.NewArticleController()

	web := s.router.Group("/web")
	web.GET("/:id", articleController.Home)

	apiRoute := s.router.Group("/api/articles")
	apiRoute.POST("/", articleController.Create)
	apiRoute.GET("/", articleController.AllArticles)
	apiRoute.GET("/:id", articleController.ArticleShow)
}

func (s *Server) Run() {
	s.router.Run(fmt.Sprintf("%s:%s", config.Configs.Server.Host, config.Configs.Server.Port)) // listen and serve on 0.0.0.0:8080
}
