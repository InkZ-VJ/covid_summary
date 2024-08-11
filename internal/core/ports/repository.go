package ports

import (
	"context"

	"covid/internal/core/domains"
)

type CovidRepository interface {
	Create(ctx context.Context, in domains.CovidSummary) (*domains.CovidSummary, error)
}
