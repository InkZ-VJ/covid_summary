package ports

import (
	"context"

	"covid/internal/dtos"
)

type CovidAdapter interface {
	GetCovidStat(ctx context.Context) (*dtos.CovidResponse, error)
}
