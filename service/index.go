package service

import (
	"html/template"
	"log"
	"myWeb/config"
	"myWeb/dao"
	"myWeb/models"
)

func GetAllIndexInfo(slug string, page, limit int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("查询分类异常")
		return nil, err
	}
	var posts []models.Post

	if slug == "" {
		posts, err = dao.GetPostArticlePage(page, limit)
		if err != nil {
			log.Println("查询文章异常")
			return nil, err
		}
	} else {
		posts, err = dao.GetPostBySlug(slug, page, limit)
		if err != nil {
			log.Println("查询分类异常")
			return nil, err
		}
	}

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	var total int
	if slug == "" {
		total = dao.CountGetAllPost()
	} else {
		total = dao.CountGetAllPostBySlug(slug)
	}

	pagesCount := (total-1)/limit + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}

	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,          //文章
		total,              //文章总数
		page,               //当前页
		pages,              //页码,两页就是[]int{1,2}
		page != pagesCount, //是否有下一页
	}
	return hr, nil
}
