package covidsvc

import (
	"context"

	"covid/internal/core/domains"
	"covid/internal/core/ports"
	"covid/internal/dtos"
)

type service struct {
	ca ports.CovidAdapter
	cr ports.CovidRepository
}

func New(ca ports.CovidAdapter, cr ports.CovidRepository) ports.CovidService {
	return &service{
		ca: ca,
		cr: cr,
	}
}

func (s *service) GetSummary(ctx context.Context) (*domains.CovidSummary, error) {
	data, err := s.ca.GetCovidRecords(ctx)
	if err != nil {
		return nil, err
	}
	sum := s.Summary(data)
	out, err := s.cr.Create(ctx, *sum)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *service) Summary(in *dtos.CovidResponse) *domains.CovidSummary {
	summary := &domains.CovidSummary{
		Province: make(map[string]int),
		AgeGroup: make(map[string]int),
	}
	for _, record := range in.Data {
		if record.Province != "" {
			summary.Province[record.Province]++
		} else {
			summary.Province["N/A"]++
		}

		switch {
		case record.Age >= 0 && record.Age <= 30:
			summary.AgeGroup["0-30"]++
		case record.Age >= 31 && record.Age <= 60:
			summary.AgeGroup["31-60"]++
		case record.Age > 60:
			summary.AgeGroup["61+"]++
		default:
			summary.AgeGroup["N/A"]++
		}
	}
	return summary
}
