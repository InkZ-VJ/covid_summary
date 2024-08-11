package ports

import (
	"context"

	"covid/internal/dtos"
)

type CovidAdapter interface {
	GetCovidRecords(ctx context.Context) (*dtos.CovidResponse, error)
}
