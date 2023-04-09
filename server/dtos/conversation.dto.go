package dtos

import (
	"server/models"
	"time"
)

type ConversationBasicDTO struct {
	ID     string    `json:"id" bson:"_id"`
	Avatar string    `json:"avatar" bson:"avatar"`
	Me     UserDTO   `json:"me" bson:"me"`
	Users  []UserDTO `json:"users" bson:"users"`
}

type ConversationDTO struct {
	ID        string       `json:"id" bson:"_id"`
	Avatar    string       `json:"avatar" bson:"avatar"`
	Me        UserDTO      `json:"me" bson:"me"`
	Users     []UserDTO    `json:"users" bson:"users"`
	Messages  []MessageDTO `json:"messages" bson:"messages"`
	CreatedAt time.Time    `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt" bson:"updatedAt"`
}

func MapConversationBasicDTO(conversation *models.Conversation, myID string, users []*models.User) ConversationBasicDTO {
	var me *models.User
	var otherUsers []*models.User

	for _, user := range users {
		if user.ID != myID {
			otherUsers = append(otherUsers, user)
		} else {
			me = user
		}
	}

	return ConversationBasicDTO{
		ID:     conversation.ID,
		Avatar: conversation.Avatar,
		Me:     MapUserDTO(me),
		Users:  MapUsersDTO(otherUsers),
		// Users:  MapUsersDTO(users),
	}
}

func MapConversationDTO(conversation *models.Conversation, myID string, users []*models.User, messages []*models.Message) ConversationDTO {
	var me *models.User
	var otherUsers []*models.User

	for _, user := range users {
		if user.ID != myID {
			otherUsers = append(otherUsers, user)
		} else {
			me = user
		}
	}

	return ConversationDTO{
		ID:     conversation.ID,
		Avatar: conversation.Avatar,
		Me:     MapUserDTO(me),
		Users:  MapUsersDTO(otherUsers),
		// Users:     MapUsersDTO(users),
		Messages:  MapMessagesDTO(messages),
		CreatedAt: conversation.CreatedAt,
		UpdatedAt: conversation.UpdatedAt,
	}
}

func MapConversationsBasicDTO(conversations []*models.Conversation, myID string, users []*models.User) []ConversationBasicDTO {
	var conversationsDTO []ConversationBasicDTO

	for _, conversation := range conversations {
		var conversationUsers []*models.User

		for _, userID := range conversation.Users {
			for _, user := range users {
				if userID == user.ID {
					conversationUsers = append(conversationUsers, user)
				}
			}
		}

		conversationDTO := MapConversationBasicDTO(conversation, myID, conversationUsers)

		conversationsDTO = append(conversationsDTO, conversationDTO)
	}

	return conversationsDTO
}
