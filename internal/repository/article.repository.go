package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"monopc-starter/internal/model" // Import the model package
	"time"
)

type ArticleRepositoryUseCase interface {
	Create(ctx context.Context, article *model.Article) error
	Update(ctx context.Context, article *model.Article) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Article, error)
	GetById(ctx context.Context, id int) (*model.Article, error)
}

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepositoryUseCase {
	return &ArticleRepository{db}
}

func (r *ArticleRepository) Create(ctx context.Context, article *model.Article) error {
	if err := r.db.WithContext(ctx).Model(&model.Article{}).Create(&article).Error; err != nil {
		return errors.Wrap(err, "error creating article")
	}
	return nil
}

func (r *ArticleRepository) Update(ctx context.Context, article *model.Article) error {
	oldTime := article.UpdatedAt
	newTime := time.Now()

	var txnError error

	txnError = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sourceModel := new(model.Article)
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(sourceModel, article.ID).Error; err != nil {
			log.Println("[ArticleRepository-Update]", err)
			return err
		}

		updates := sourceModel.MapUpdateFrom(article)
		if len(*updates) > 0 {
			(*updates)["updated_at"] = newTime

			if err := tx.Model(&model.Article{}).Where("id = ?", article.ID).UpdateColumns(updates).Error; err != nil {
				log.Println("[ArticleRepository-Update]", err)
				return err
			}
		}

		return nil
	})

	if txnError != nil {
		article.UpdatedAt = oldTime
		return txnError
	}

	article.UpdatedAt = newTime
	return nil
}

func (r *ArticleRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Model(&model.Article{}).Delete(&model.Article{}, id).Error; err != nil {
		return errors.Wrap(err, "error deleting article")
	}
	return nil
}

func (repo *ArticleRepository) GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Article, error) {
	var articles []*model.Article

	q := repo.db.WithContext(ctx).Model(&model.Article{})

	if query != "" {
		q = q.Where("title LIKE ?", "%"+query+"%")
	}

	if sort != "" && order != "" {
		q = q.Order(sort + " " + order)
	} else {
		q = q.Order("created_at DESC")
	}

	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}

	if err := q.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ArticleRepository) GetById(ctx context.Context, id int) (*model.Article, error) {
	var article model.Article
	if err := r.db.WithContext(ctx).Model(&model.Article{}).First(&article, id).Error; err != nil {
		return nil, errors.Wrap(err, "error getting article by id")
	}
	return &article, nil
}
