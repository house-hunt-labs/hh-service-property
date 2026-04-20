package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// NodeTypeStyle represents the style for each node type
type NodeTypeStyle struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Type      string             `bson:"type" json:"type"`
	Color     string             `bson:"color" json:"color"`
	MaxRadius float64            `bson:"maxRadius" json:"maxRadius"`
}

// MapNodeMetadata contains additional metadata from Google Places
type MapNodeMetadata struct {
	GooglePlaceID    string  `bson:"google_place_id" json:"google_place_id"`
	Rating           float64 `bson:"rating" json:"rating"`
	UserRatingsTotal int32   `bson:"user_ratings_total" json:"user_ratings_total"`
}

// MapNode represents a map node
type MapNode struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Type        string             `bson:"type" json:"type"`
	Label       string             `bson:"label" json:"label"`
	Latitude    float64            `bson:"latitude" json:"latitude"`
	Longitude   float64            `bson:"longitude" json:"longitude"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Metadata    *MapNodeMetadata   `bson:"metadata,omitempty" json:"metadata,omitempty"`
}
