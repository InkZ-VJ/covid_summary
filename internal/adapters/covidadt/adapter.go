package covidadt

import (
	"context"

	"covid/config"
	"covid/internal/core/ports"
	"covid/internal/dtos"

	"github.com/go-resty/resty/v2"
)

type adapter struct {
	rc *resty.Client
}

func New() ports.CovidAdapter {
	return &adapter{
		rc: resty.New(),
	}
}

func (a *adapter) GetCovidRecords(ctx context.Context) (*dtos.CovidResponse, error) {
	var out dtos.CovidResponse
	_, err := a.rc.R().
		SetResult(&out).
		Get(config.Get().External.Covid)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
