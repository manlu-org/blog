// Code generated by taoctl. DO NOT EDIT.
package handler

import (
	"net/http"

	article "backend-learning/blog/app/article/cmd/api/internal/handler/article"
	"backend-learning/blog/app/article/cmd/api/internal/svc"

	"manlu.org/tao/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/article/list",
				Handler: article.GetArticleListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/article/:id",
				Handler: article.ShowArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/article/create",
				Handler: article.CreateArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/article/:id/del",
				Handler: article.DeleteArticleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}