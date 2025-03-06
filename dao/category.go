package dao

import (
	"database/sql"
	"log"
	"myWeb/models"
)

func GetAllCategory(page, limit int) ([]models.Category, error) {
	// 计算偏移量
	offset := (page - 1) * limit

	// 执行查询
	rows, err := DB.Query("SELECT cid, name, create_at, update_at FROM blog_category LIMIT ?,?", limit, offset)
	if err != nil {
		log.Printf("查询分类异常: %v", err)
		return nil, err
	}
	defer rows.Close()

	// 解析数据
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Printf("解析分类数据异常: %v", err)
			return nil, err
		}
		categories = append(categories, category)
	}

	// 检查迭代时是否出错
	if err = rows.Err(); err != nil {
		log.Printf("查询分类迭代异常: %v", err)
		return nil, err
	}

	return categories, nil
}

func GetCategoryNameById(cid int) string {
	var name string
	err := DB.QueryRow("SELECT name FROM blog_category WHERE cid = ?", cid).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("未找到分类 ID %d 的名称", cid)
			return ""
		}
		log.Printf("查询分类名称异常：%v", err)
		return ""
	}
	return name
}
