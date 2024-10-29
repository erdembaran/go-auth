package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`        
	Username  string             `bson:"username" json:"username"`        
	Email     string             `bson:"email" json:"email"`             
	Password  string             `bson:"password" json:"password"`       
	CreatedAt primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt"` 
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt"` 
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
