package service

import (
	"html/template"
	"log"
	"myWeb/config"
	"myWeb/dao"
	"myWeb/models"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pid)
	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postRes, nil
}

func Writing() (wr models.WriteRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Printf("查询分类异常: %v", err)
		return
	}
	wr.Categorys = categorys
	return
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func GetPostById(pid int) (*models.Post, error) {
	return dao.GetPostById(pid)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func SearchPost(val string) ([]models.SearchResp, error) {
	return dao.SearchPost(val)
}
