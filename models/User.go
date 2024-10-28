package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`         // MongoDB ObjectID
	Username  string             `bson:"username" json:"username"`        // Username field
	Email     string             `bson:"email" json:"email"`              // Email field
	Password  string             `bson:"password" json:"password"`        // Password field (usually hashed)
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt"` // Creation timestamp
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt"` // Last updated timestamp
}
