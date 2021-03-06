package main

import (
	"example.com/echo_test/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "public/assets")

	e.File("/", "public/index.html")            // GET /
	e.File("/signup", "public/signup.html")     // GET /signup
	e.POST("/signup", handler.Signup)                // POST /signup
	e.File("/login", "public/login.html")       // GET /login
	e.POST("/login", handler.Login)                  // POST /login
	e.File("/todos", "public/todos.html")       // GET /todos
	e.File("/articles", "public/articles.html") // GET /articles
	e.File("/articles/:id", "public/article_detail.html") // GET /articles/:id

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config)) // /api 下はJWTの認証が必要

	api.GET("/todos", handler.GetTodos)                 // GET /api/todos
	api.POST("/todos", handler.AddTodo)                 // POST /api/todos
	api.DELETE("/todos/:id", handler.DeleteTodo)        // DELETE /api/todos/:id
	api.PUT("/todos/:id/completed", handler.UpdateTodo) // PUT /api/todos/:id/completed

	api.GET("/articles", handler.GetArticles)          // GET /api/articles
	api.GET("/articles/:id", handler.GetArticle)       // GET /api/articles/:id
	api.POST("/articles", handler.AddArticle)          // POST /api/articles
	api.DELETE("/articles/:id", handler.DeleteArticle) // DELETE /api/articles/:id
	api.PUT("/articles/:id", handler.UpdateArticle)    // PUT /api/articles/:id

	return e
}
