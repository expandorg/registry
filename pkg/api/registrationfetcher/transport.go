package registrationfetcher

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gemsorg/dispute/pkg/apierror"
	service "github.com/gemsorg/registry/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s service.RegistryService) http.Handler {
	return kithttp.NewServer(
		makeRegistrationFetcherEndpoint(s),
		decodeRegistrationFetcherRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeRegistrationFetcherRequest(_ context.Context, r *http.Request) (interface{}, error) {
	dr := JobRegistrationRequest{}
	vars := mux.Vars(r)
	jobID, ok := vars["job_id"]
	if !ok {
		return nil, errorResponse(&apierror.ErrBadRouting{Param: "job_id"})
	}
	id, err := strconv.ParseUint(jobID, 10, 64)
	if err != nil {
		return nil, errorResponse(&apierror.ErrBadRouting{Param: "dispute_id"})
	}
	dr.JobID = id
	return dr, nil
}
