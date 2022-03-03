package repository

import (
	"context"
	"gorm.io/gorm"
	"handle-big-post-request/models"
)

type postDataRepository struct {
	Conn *gorm.DB
}

func NewPostDataRepository(Conn *gorm.DB) PostDataRepository {
	return &postDataRepository{Conn}
}

func (r *postDataRepository) BatchInsert(ctx context.Context, postData *[]models.PostSubmit) error {
	rs := r.Conn.WithContext(ctx).Create(postData)
	return rs.Error
}
