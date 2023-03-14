package models

import "time"

type User struct {
	ID           string    `json:"id" bson:"_id"`
	Username     string    `json:"username" bson:"username"`
	Email        string    `json:"email" bson:"email"`
	Avatar       string    `json:"avatar" bson:"avatar"`
	Password     string    `json:"password" bson:"password"`
	Status       uint8     `json:"status" bson:"status"`
	CustomStatus string    `json:"customStatus" bson:"customStatus"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}

type UserResponse struct {
	ID           string `json:"id" bson:"_id"`
	Username     string `json:"username" bson:"username"`
	Email        string `json:"email" bson:"email"`
	Avatar       string `json:"avatar" bson:"avatar"`
	Status       string `json:"status" bson:"status"`
	CustomStatus string `json:"customStatus" bson:"customStatus"`
}

func TransformStatus(status uint8) string {
	enum := map[uint8]string{
		0: "Offline",
		1: "Do Not Disturb",
		2: "Idle",
		3: "Online",
	}

	return enum[status]
}

func UserFilteredResponse(user *User) UserResponse {
	return UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Avatar:       user.Avatar,
		Status:       TransformStatus(user.Status),
		CustomStatus: user.CustomStatus,
	}
}
