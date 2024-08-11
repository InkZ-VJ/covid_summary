package covidrepo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"covid/internal/core/ports"
)

type repo struct {
	col *mongo.Collection
	db  string
}

func New() ports.CovidRepository {
	return &repo{}
}
