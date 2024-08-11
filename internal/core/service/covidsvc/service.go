package covidsvc

import (
	"context"

	"covid/internal/core/domains"
	"covid/internal/core/ports"
	"covid/internal/dtos"
)

type service struct {
	ca ports.CovidAdapter
}

func New(ca ports.CovidAdapter) ports.CovidService {
	return &service{
		ca: ca,
	}
}

func (s *service) GetSummary(ctx context.Context) (*domains.CovidSummary, error) {
	data, err := s.ca.GetCovidRecords(ctx)
	if err != nil {
		return nil, err
	}
	sum := s.summary(data)
	return sum, nil
}

func (s *service) summary(in *dtos.CovidResponse) *domains.CovidSummary {
	var summary domains.CovidSummary
	for _, record := range in.Data {
		// Count patients by province
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
	return &summary
}
