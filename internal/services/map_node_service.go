package services

import (
	"context"

	"github.com/house-hunt-labs/hh-service-property/internal/models"
	"github.com/house-hunt-labs/hh-service-property/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MapNodeService struct {
	repo *repositories.MapNodeRepository
}

func NewMapNodeService(repo *repositories.MapNodeRepository) *MapNodeService {
	return &MapNodeService{repo: repo}
}

func (s *MapNodeService) Create(ctx context.Context, node *models.MapNode) error {
	return s.repo.Create(ctx, node)
}

func (s *MapNodeService) GetAll(ctx context.Context) ([]models.MapNode, error) {
	return s.repo.GetAll(ctx)
}

func (s *MapNodeService) GetByID(ctx context.Context, id primitive.ObjectID) (*models.MapNode, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MapNodeService) Update(ctx context.Context, id primitive.ObjectID, node *models.MapNode) error {
	return s.repo.Update(ctx, id, node)
}

func (s *MapNodeService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}

func (s *MapNodeService) GetByType(ctx context.Context, nodeType string) ([]models.MapNode, error) {
	return s.repo.GetByType(ctx, nodeType)
}
