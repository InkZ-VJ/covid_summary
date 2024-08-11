package ports

import (
	"context"

	"covid/internal/core/domains"
)

type CovidService interface {
	GetSummary(ctx context.Context) (*domains.CovidSummary, error)
}
