package domains

import "go.mongodb.org/mongo-driver/bson/primitive"

type CovidSummary struct {
	Province map[string]int     `bson:"province" json:"Province"`
	AgeGroup map[string]int     `bson:"ageGroup" json:"AgeGroup"`
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
}
