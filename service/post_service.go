package service

import (
	"context"
	"handle-big-post-request/models"
	"handle-big-post-request/queue"
	"handle-big-post-request/queue/payload"
	"handle-big-post-request/repository"
)

type postService struct {
	postDataRepo repository.PostDataRepository
}

func NewPostService(dataRepository repository.PostDataRepository) PostDataService {
	return &postService{dataRepository}
}

func (s *postService) BatchInsert(ctx context.Context, postData *[]models.PostData) error {
	return s.postDataRepo.BatchInsert(ctx, postData)
}

func (s *postService) PushDataToQueue(poolJob *queue.PoolJob, payload payload.Payload) {
	poolJob.PushDataToQueue(payload)
}
