package service

import (
	"log"
	"myWeb/config"
	"myWeb/dao"
	"myWeb/models"
)

func FindPostByPigeonhole() *models.PigeholeRes {

	posts, err := dao.GetAllPost()
	if err != nil {
		log.Printf("GetAllPost查询归档失败: %v", err)
		return nil
	}
	//按月份归档
	//map[string]Post
	archive := make(map[string][]models.Post)
	for _, post := range posts {
		month := post.CreateAt.Format("2006-01")
		archive[month] = append(archive[month], post)
	}
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Printf("GetAllCategory查询分类失败: %v", err)
		return nil
	}

	return &models.PigeholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		archive,
	}
}
