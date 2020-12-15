package model

import "fmt"

type Article struct {
	UID       int    `json:"uid"`
	ID        int    `json:"id" gorm:"praimaly_key"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type Articles []Article

func CreateArticle(article *Article) {
	db.Create(article)
}

func FindArticles(a *Article) Articles {
	var articles Articles
	db.Where(a).Find(&articles)
	return articles
}

func FindArticle(a *Article) Article {
	var article Article
	db.Where(a).First(&article)
	return article
}

func DeleteArticle(a *Article) error {
	if rows := db.Where(a).Delete(&Article{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Article (%v) to delete", a)
	}
	return nil
}

func UpdateArticle(a *Article) error {
	rows := db.Model(a).Update(map[string]interface{}{
		"title": a.Title,
		"content": a.Content,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Article (%v) to update", a)
	}
	return nil
}
