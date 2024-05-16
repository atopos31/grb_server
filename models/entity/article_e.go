package entity

import (
	"gorm.io/gorm"
)

// 文章
type Article struct {
	gorm.Model
	Uuid       string `gorm:"type:varchar(255);not null;index;"` // 	文章uuid
	Title      string `gorm:"type:varchar(255);not null"`        //     文章标题
	Summary    string `gorm:"type:varchar(255);"`                // 	文章摘要
	Content    string `gorm:"type:longtext;not null"`            // 	文章内容
	CoverImage string `gorm:"type:varchar(255);"`                // 	文章封面
	Top        uint8  `gorm:"type:int(1);default:0"`             //	    是否置顶 0：否 1：是
	Status     uint8  `gorm:"type:int(1);default:0"`             // 	文章状态 0：未发布 1：已发布
	Views      int    `gorm:"type:int(11);default:0"`            // 	文章浏览量
	CategoryID uint   // 	文章分类ID
	Category   Category
	Tags       []Tag `gorm:"many2many:articles_tags;"` // 	文章标签
}

// 文章分类
type Category struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255);not null;unique"`
	Articles []Article `gorm:"foreignKey:CategoryID"`
}

// 文章标签
type Tag struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255);not null;unique"`
	Articles []Article `gorm:"many2many:articles_tags;"`
}