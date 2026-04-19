package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// NodeTypeStyle represents the style for each node type
type NodeTypeStyle struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Type      string             `bson:"type" json:"type"`
	Color     string             `bson:"color" json:"color"`
	MaxRadius float64            `bson:"maxRadius" json:"maxRadius"`
}

// MapNode represents a map node
type MapNode struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Type        string             `bson:"type" json:"type"`
	Label       string             `bson:"label" json:"label"`
	Latitude    float64            `bson:"latitude" json:"latitude"`
	Longitude   float64            `bson:"longitude" json:"longitude"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}
