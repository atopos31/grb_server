package req

type Article struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	CoverImage string   `json:"cover_image"`                    // 文章封面
	CategoryID uint     `json:"category_id" binding:"required"` // 分类ID
	TagNames   []string `json:"tags"`                           // 标签
	Top        uint8    `json:"top" binding:"oneof=0 1"`        // 是否置顶
	Status     uint8    `json:"status" binding:"oneof=0 1"`     // 文章状态
	CreatedAt  int64    `json:"created_at"`                     // 发布时间
}

type ArticleList struct {
	PageSize   int    `form:"page_size" binding:"gte=1,lte=20"`
	PageNum    int    `form:"page_num" binding:"gte=1"`
	TitleLike  string `form:"title_like"`
	CategoryID int    `form:"cate_id"`
}

type ArticleSertion struct {
	Key   string `json:"key" binding:"oneof=top status"` // 用于更新文章状态
	Value uint8  `json:"value" binding:"oneof=0 1"`
}
