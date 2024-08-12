package covidsvc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"covid/internal/core/ports"
	"covid/internal/core/ports/mocks"
	"covid/internal/core/service/covidsvc"
	"covid/internal/dtos"
)

var (
	ctx       = context.Background()
	errorMock = errors.New("error")
)

type test struct {
	mockFn   func(*testModule)
	assestFn func()
	name     string
	args     []interface{}
}

type testModule struct {
	ca      *mocks.CovidAdapter
	cr      *mocks.CovidRepository
	service ports.CovidService
}

func newService(t *testing.T) testModule {
	ca := mocks.NewCovidAdapter(t)
	cr := mocks.NewCovidRepository(t)
	service := covidsvc.New(ca, cr)
	return testModule{
		ca:      ca,
		cr:      cr,
		service: service,
	}
}

func TestGetSummary(t *testing.T) {
	var resp error
	in := &dtos.CovidResponse{
		Data: RandomRecords(int(RandomInt(1, 2000))),
	}

	tests := []test{
		{
			name: "Request Records Fail",
			args: []interface{}{ctx},
			mockFn: func(m *testModule) {
				m.ca.On("GetCovidRecords", ctx).Return(nil, errorMock)
			},
			assestFn: func() {
				assert.Error(t, resp)
			},
		},
		{
			name: "Insert to Database Failed",
			args: []interface{}{ctx},
			mockFn: func(m *testModule) {
				m.ca.On("GetCovidRecords", ctx).Return(in, nil)
				out := m.service.Summary(in)
				m.cr.On("Create", ctx, *out).Return(nil, errorMock)
			},
			assestFn: func() {
				assert.Error(t, resp)
			},
		},
		{
			name: "Happy",
			args: []interface{}{ctx},
			mockFn: func(m *testModule) {
				m.ca.On("GetCovidRecords", ctx).Return(in, nil)
				out := m.service.Summary(in)
				m.cr.On("Create", ctx, *out).Return(out, nil)
			},
			assestFn: func() {
				assert.NoError(t, resp)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mod := newService(t)
			if tt.mockFn != nil {
				tt.mockFn(&mod)
			}
			_, resp = mod.service.GetSummary(tt.args[0].(context.Context))
			tt.assestFn()
		})
	}
}

func TestSummary(t *testing.T) {
	t.Parallel()
	svc := newService(t)
	for range 10 {
		num := int(RandomInt(1, 2000))
		in := dtos.CovidResponse{
			Data: RandomRecords(num),
		}
		out := svc.service.Summary(&in)

		cp, ca := 0, 0
		for _, value := range out.Province {
			cp += value
		}

		for _, value := range out.AgeGroup {
			ca += value
		}

		assert.Equal(t, num, cp)
		assert.Equal(t, num, ca)
	}
}
