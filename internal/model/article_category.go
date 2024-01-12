package model

type ArticleCategory struct {
	ArticleID  int `json:"article_id"`
	CategoryID int `json:"category_id"`

	Article  *Article  `json:"article"`
	Category *Category `json:"category"`
	Auditable
}

func NewArticleCategory(
	articleID int,
	categoryID int,
	createdBy string,
) *ArticleCategory {
	return &ArticleCategory{
		ArticleID:  articleID,
		CategoryID: categoryID,
		Auditable:  NewAuditable(createdBy),
	}
}
