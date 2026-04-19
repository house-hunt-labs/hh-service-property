package repositories

import (
	"context"

	"github.com/house-hunt-labs/hh-service-property/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MapNodeRepository struct {
	collection *mongo.Collection
}

func NewMapNodeRepository(db *mongo.Database) *MapNodeRepository {
	return &MapNodeRepository{
		collection: db.Collection("map_nodes"),
	}
}

func (r *MapNodeRepository) Create(ctx context.Context, node *models.MapNode) error {
	_, err := r.collection.InsertOne(ctx, node)
	return err
}

func (r *MapNodeRepository) GetAll(ctx context.Context) ([]models.MapNode, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var nodes []models.MapNode
	err = cursor.All(ctx, &nodes)
	return nodes, err
}

func (r *MapNodeRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.MapNode, error) {
	var node models.MapNode
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func (r *MapNodeRepository) Update(ctx context.Context, id primitive.ObjectID, node *models.MapNode) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": node})
	return err
}

func (r *MapNodeRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MapNodeRepository) GetByType(ctx context.Context, nodeType string) ([]models.MapNode, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"type": nodeType})
	if err != nil {
		return nil, err
	}
	var nodes []models.MapNode
	err = cursor.All(ctx, &nodes)
	return nodes, err
}
