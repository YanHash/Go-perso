package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tool struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name"`
	Usage             string             `bson:"usage"`
	AssociatedProject string             `bson:"associated_project"`
	MasteryLevel      int                `bson:"mastery_level"`
	IconURL           string             `bson:"icon_url"`
}

type Experience struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Position  string             `bson:"position"`
	Company   string             `bson:"company"`
	Missions  string             `bson:"missions"`
	StartDate string             `bson:"start_date"`
	EndDate   string             `bson:"end_date"`
}

type Education struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Institution string             `bson:"institution"`
	DegreeName  string             `bson:"degree_name"`
	Description string             `bson:"description"`
	StartDate   string             `bson:"start_date"`
	EndDate     string             `bson:"end_date"`
	Result      string             `bson:"result"`
}

type Contact struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ContactType string             `bson:"contact_type"`
	ContactLink string             `bson:"contact_link"`
}
