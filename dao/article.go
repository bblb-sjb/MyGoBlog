package dao

import (
	"log"
	"myWeb/models"
)

func GetPostArticlePage(page, limit int) ([]models.Post, error) {
	row, err := DB.Query("select * from blog_post limit ?,?", (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var posts []models.Post
	for row.Next() {
		var post models.Post
		err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CountGetAllPost() int {
	var count int
	err := DB.QueryRow("SELECT COUNT(1) FROM blog_post").Scan(&count)
	if err != nil {
		log.Printf("查询文章总数失败: %v", err)
		return 0
	}
	return count
}
