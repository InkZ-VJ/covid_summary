package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"covid/internal/core/domains"
	"covid/internal/core/ports"
)

type repository struct {
	mc  *mongo.Client
	col *mongo.Collection
	db  string
}

func NewCovidRepo(mc *mongo.Client, db string) ports.CovidRepository {
	return &repository{
		mc:  mc,
		col: mc.Database(db).Collection("summary"),
		db:  db,
	}
}

func (r *repository) Create(ctx context.Context, in domains.CovidSummary) (*domains.CovidSummary, error) {
	return r.insertOne(ctx, in)
}

func (r *repository) insertOne(ctx context.Context, in domains.CovidSummary) (*domains.CovidSummary, error) {
	in.CreateAt = time.Now().Local().UTC()
	result, err := r.col.InsertOne(ctx, in)
	if err != nil {
		return nil, err
	}
	oid, _ := result.InsertedID.(primitive.ObjectID)
	in.ID = oid
	return &in, err
}
