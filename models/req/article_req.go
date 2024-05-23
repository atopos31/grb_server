package req

type Article struct {
	Uuid       uint32   `json:"uuid"` //uuid 在更新文章时需要
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	CoverImage string   `json:"cover_image"`                // 文章封面
	CategoryID uint     `json:"category_id"`                // 分类ID
	TagNames   []string `json:"tags"`                       // 标签
	Top        uint8    `json:"top" binding:"oneof=0 1"`    // 是否置顶
	Status     uint8    `json:"status" binding:"oneof=0 1"` // 文章状态
}

type ArticleList struct {
	PageSize int `form:"page_size" binding:"gte=1,lte=20"`
	PageNum  int `form:"page_num" binding:"gte=1"`
}
