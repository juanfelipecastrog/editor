package controller

import (
	"github.com/gin-gonic/gin"
	"gold/internal/db"
	"gold/internal/model"
	"net/http"
)

type ArticleController interface {
	Home(c *gin.Context)
	Create(c *gin.Context)
	AllArticles(c *gin.Context)
	ArticleShow(c *gin.Context)
}

type articleController struct {
}

func NewArticleController() ArticleController {
	return &articleController{}
}

func (ac *articleController) Home(c *gin.Context) {
	articleID := c.Param("id")
	var article model.Article
	db.Connection().First(&article, articleID)

	c.HTML(http.StatusOK, "template/home/index", gin.H{
		"NAME":  "Home Page",
		"TITLE": article.Title,
		"BODY":  article.Body,
	})
}

func (ac *articleController) Create(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	article := model.Article{Title: body.Title, Body: body.Body}
	result := db.Connection().Create(&article)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"article": article,
	})
}

func (ac *articleController) AllArticles(c *gin.Context) {
	var articles []model.Article
	db.Connection().Find(&articles)

	c.JSON(200, gin.H{
		"article": articles,
	})
}

func (ac *articleController) ArticleShow(c *gin.Context) {
	id := c.Param("id")
	var article model.Article
	db.Connection().First(&article, id)

	c.JSON(200, gin.H{
		"article": article,
	})
}
