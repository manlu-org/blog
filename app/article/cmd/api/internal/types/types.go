// Code generated by taoctl. DO NOT EDIT.
package types

type PageReq struct {
	PageSize int64 `form:"page_size"`
	Page     int64 `form:"page"`
}

type Meta struct {
	PageSize int64 `json:"page_size"`
	Page     int64 `json:"page":"page"`
	Total    int64 `json:"total"`
}

type Article struct {
	Id      int64  `json:"id"`
	Titicle string `json:"titicle"`
	Content string `json:"content"`
}

type GetArticleListResp struct {
	List interface{} `json:"list"`
	Meta *Meta       `json:"meta"`
}

type ShowArticleReq struct {
	Id int64 `path:"id"`
}

type ShowArticleResp struct {
	Article *Article `json:"article"`
}

type CreateArticleReq struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

type CreateArticleResp struct {
	Success bool `json:"success"`
}

type DeleteArticleReq struct {
	Id int64 `path:"id"`
}

type DeleteArticleResp struct {
	Success bool `json:"success"`
}
