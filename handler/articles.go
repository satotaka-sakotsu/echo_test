package handler

import (
	"net/http"
	"strconv"

	"example.com/echo_test/model"
	"github.com/labstack/echo/v4"
)

func AddArticle(c echo.Context) error {
	article := new(model.Article)
	if err := c.Bind(article); err != nil {
		return err
	}

	if article.Title == "" || article.Content == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	article.UID = uid
	model.CreateArticle(article)

	return c.JSON(http.StatusCreated, article)
}

func GetArticles(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	articles := model.FindArticles(&model.Article{UID: uid})
	return c.JSON(http.StatusOK, articles)
}

func GetArticle(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	article := model.FindArticle(&model.Article{ID: articleID})
	return c.JSON(http.StatusOK, article)
}

func DeleteArticle(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteArticle(&model.Article{ID: articleID,
		UID: uid}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateArticle(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	article := model.FindArticle(&model.Article{ID: articleID, UID: uid})
	if article.ID == 0 {
		return echo.ErrNotFound
	}

	a := new(model.Article)
	if err := c.Bind(a); err != nil {
		return err
	}
	if a.Title == "" || a.Content == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	article.Title = a.Title
	article.Content = a.Content
	if err := model.UpdateArticle(&article); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}
