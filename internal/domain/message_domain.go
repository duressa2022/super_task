package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	MessageCollection      = "messages"
	ConversationCollection = "conversation"
)

type Message struct {
	MessageId      primitive.ObjectID `json:"messageID" bson:"_messageID"`
	ConversationId primitive.ObjectID `json:"conversationID" bson:"_conversationID"`
	SenderID       primitive.ObjectID `json:"senderID" bson:"_senderID"`
	ReceipentID    primitive.ObjectID `json:"receipentID" bson:"_receipentID"`
	MessageType    int                `json:"messageType" bson:"messageType"`
	MessageContent string             `json:"messageContent" bson:"messageContent"`
	MediaUrl       string             `json:"mediaUrl" bson:"medaiurl"`
	Status         string             `json:"status" bson:"status"`
	TimeStamp      primitive.DateTime `json:"timeStamp" bson:"timeStamp"`
}

type OneToOneMessage struct {
	SenderID       string `json:"senderID" bson:"_senderID"`
	ReceipentID    string `json:"receipentID" bson:"_receipentID"`
	MessageContent string `json:"messageContent" bson:"messageContent"`
	MediaUrl       string `json:"mediaUrl" bson:"medaiurl"`
}

type GroupMessage struct {
	MessageId      primitive.ObjectID `json:"messageID" bson:"_messageID"`
	ConversationId primitive.ObjectID `json:"conversationID" bson:"_conversationID"`
	GroupID        primitive.ObjectID `json:"groupID" bson:"_groupID"`
	SenderID       primitive.ObjectID `json:"senderID" bson:"_senderID"`
	MessageType    int                `json:"messageType" bson:"messageType"`
	MessageContent string             `json:"messageContent" bson:"messageContent"`
	MediaUrl       string             `json:"mediaUrl" bson:"medaiurl"`
	Status         string             `json:"status" bson:"status"`
	TimeStamp      primitive.DateTime `json:"timeStamp" bson:"timeStamp"`
}
type GroupMessageRequest struct {
	SenderID       string `json:"senderID" bson:"_senderID"`
	GroupID        string `json:"groupID" bson:"_groupID"`
	MessageContent string `json:"messageContent" bson:"messageContent"`
	MediaUrl       string `json:"mediaUrl" bson:"medaiurl"`
}

type MessageRequest struct {
	ConversationId primitive.ObjectID `json:"conversationID" bson:"_conversationID"`
	MessageType    string             `json:"messageType" bson:"messageType"`
	MessageContent string             `json:"messageContent" bson:"messageContent"`
	MediaUrl       string             `json:"mediaUrl" bson:"medaiurl"`
	Status         string             `json:"status" bson:"status"`
}

type EditMessage struct {
	MessageContent string `json:"messageContent" bson:"messageContent"`
	MediaUrl       string `json:"mediaUrl" bson:"medaiurl"`
	Status         string `json:"status" bson:"status"`
}

type Conversation struct {
	ConversationID primitive.ObjectID `json:"conversationID" bson:"_conversationID"`
	IsGroup        bool               `json:"isGroup" bson:"isGroup"`
	Participants   []Participant      `json:"participants" bson:"participants"`
	LastMessage    Message            `json:"lastMessage" bson:"lastMessage"`
	CreatedAt      primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt      primitive.DateTime `json:"updatedAt" bson:"UpdatedAt"`
}

type GroupConversation struct {
	ConversationID primitive.ObjectID `json:"conversationID" bson:"_conversationID"`
	GroupID        primitive.ObjectID `json:"groupID" bson:"_groupID"`
	LastMessage    GroupMessage       `json:"lastMessage" bson:"lastMessage"`
	CreatedAt      primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt      primitive.DateTime `json:"updatedAt" bson:"UpdatedAt"`
}

type Participant struct {
	ParticipantID primitive.ObjectID  `json:"ParticipantID" bson:"_ParticipantID"`
	JoinedAt      primitive.Timestamp `json:"joinedAt" bson:"joinedAt"`
}
type ParticipantRequest struct {
	FirstParticipantID  primitive.ObjectID `json:"firstParticipantID" bson:"_firstParticipantID"`
	SecondParticipantID primitive.ObjectID `json:"secondParticipantID" bson:"_secondParticipantID"`
}
