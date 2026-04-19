package services

import (
	"context"

	"github.com/house-hunt-labs/hh-service-property/internal/models"
	"github.com/house-hunt-labs/hh-service-property/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NodeTypeStyleService struct {
	repo *repositories.NodeTypeStyleRepository
}

func NewNodeTypeStyleService(repo *repositories.NodeTypeStyleRepository) *NodeTypeStyleService {
	return &NodeTypeStyleService{repo: repo}
}

func (s *NodeTypeStyleService) Create(ctx context.Context, style *models.NodeTypeStyle) error {
	return s.repo.Create(ctx, style)
}

func (s *NodeTypeStyleService) GetAll(ctx context.Context) ([]models.NodeTypeStyle, error) {
	return s.repo.GetAll(ctx)
}

func (s *NodeTypeStyleService) GetByID(ctx context.Context, id primitive.ObjectID) (*models.NodeTypeStyle, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *NodeTypeStyleService) Update(ctx context.Context, id primitive.ObjectID, style *models.NodeTypeStyle) error {
	return s.repo.Update(ctx, id, style)
}

func (s *NodeTypeStyleService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}

func (s *NodeTypeStyleService) GetByType(ctx context.Context, nodeType string) (*models.NodeTypeStyle, error) {
	return s.repo.GetByType(ctx, nodeType)
}
