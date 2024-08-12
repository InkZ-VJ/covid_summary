package ports

import (
	"context"

	"covid/internal/core/domains"
	"covid/internal/dtos"
)

type CovidService interface {
	GetSummary(ctx context.Context) (*domains.CovidSummary, error)
	Summary(in *dtos.CovidResponse) *domains.CovidSummary
}
