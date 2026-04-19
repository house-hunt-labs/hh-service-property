package repositories

import (
	"context"

	"github.com/house-hunt-labs/hh-service-property/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NodeTypeStyleRepository struct {
	collection *mongo.Collection
}

func NewNodeTypeStyleRepository(db *mongo.Database) *NodeTypeStyleRepository {
	return &NodeTypeStyleRepository{
		collection: db.Collection("node_type_styles"),
	}
}

func (r *NodeTypeStyleRepository) Create(ctx context.Context, style *models.NodeTypeStyle) error {
	_, err := r.collection.InsertOne(ctx, style)
	return err
}

func (r *NodeTypeStyleRepository) GetAll(ctx context.Context) ([]models.NodeTypeStyle, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var styles []models.NodeTypeStyle
	err = cursor.All(ctx, &styles)
	return styles, err
}

func (r *NodeTypeStyleRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.NodeTypeStyle, error) {
	var style models.NodeTypeStyle
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&style)
	if err != nil {
		return nil, err
	}
	return &style, nil
}

func (r *NodeTypeStyleRepository) Update(ctx context.Context, id primitive.ObjectID, style *models.NodeTypeStyle) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": style})
	return err
}

func (r *NodeTypeStyleRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *NodeTypeStyleRepository) GetByType(ctx context.Context, nodeType string) (*models.NodeTypeStyle, error) {
	var style models.NodeTypeStyle
	err := r.collection.FindOne(ctx, bson.M{"type": nodeType}).Decode(&style)
	if err != nil {
		return nil, err
	}
	return &style, nil
}
