package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"covid/internal/core/ports"
)

type repository struct {
	col *mongo.Collection
	db  string
}

func New(client mongo.Client, db string) ports.CovidRepository {
	// col := client.
	return &repository{
		db: db,
	}
}
