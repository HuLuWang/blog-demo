package data

import (
	"blog-demo/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	log *log.Helper
}

// NewArticleRepo
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	ps, err := ar.data.db.
}