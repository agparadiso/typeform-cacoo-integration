package typeform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type formRequest struct {
	question string
}
type formResponse struct {
	err error
}

func makeBuildFormEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(formRequest)
		res := formResponse{}
		res.err = svc.BuildForm(ctx, req.question)
		return res, nil
	}
}

func decodeBuildFormRequest(_ context.Context, r *http.Request) (interface{}, error) {
	question := r.URL.Query().Get("question")
	if question == "" {
		return nil, fmt.Errorf("missing question parameter")
	}
	return formRequest{
		question: question,
	}, nil
}

func makeEncodeBuildFormResponse() func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		res := response.(formResponse)
		if res.err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//write some error
		}
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
