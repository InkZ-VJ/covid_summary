package covidsvc

import "covid/internal/core/ports"

type service struct {
	ca ports.CovidAdapter
}

func New(ca ports.CovidAdapter) ports.CovidService {
	return &service{
		ca: ca,
	}
}
