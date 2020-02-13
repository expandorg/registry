package registrationfetcher

import (
	"context"
	"encoding/json"

	"github.com/expandorg/registry/pkg/apierror"
	service "github.com/expandorg/registry/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeRegistrationFetcherEndpoint(svc service.RegistryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(JobRegistrationRequest)

		registration, err := svc.GetJobRegistration(req.JobID)

		if err != nil {
			return nil, errorResponse(err)
		}
		if registration.ID == 0 {
			return json.RawMessage("{}"), nil
		}
		return registration, nil
	}
}

type JobRegistrationRequest struct {
	JobID uint64
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
