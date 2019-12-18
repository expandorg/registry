package registrationfetcher

import (
	"context"

	"github.com/gemsorg/registry/pkg/apierror"
	service "github.com/gemsorg/registry/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeRegistrationFetcherEndpoint(svc service.RegistryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(JobRegistrationRequest)

		registrations, err := svc.GetJobRegistrations(req.JobID)
		if err != nil {
			return nil, errorResponse(err)
		}
		return registrations, nil
	}
}

type JobRegistrationRequest struct {
	JobID uint64
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
