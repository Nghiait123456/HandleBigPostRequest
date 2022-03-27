package repository

import (
	"context"
	"handle-big-post-request/src/models"
)

type PostDataRepository interface {
	BatchInsert(ctx context.Context, postData *[]models.PostSubmit) error
}
