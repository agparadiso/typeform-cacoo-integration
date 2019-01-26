package typeform

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

type formRequest struct {
	questions []string
	tfapikey  string
}
type formResponse struct {
	tflink string
	err    error
}

func makeBuildFormEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(formRequest)
		res := formResponse{}
		res.tflink, res.err = svc.CreateForm(ctx, req.questions, req.tfapikey)
		return res, nil
	}
}

func decodeBuildFormRequest(_ context.Context, r *http.Request) (interface{}, error) {
	questions := r.URL.Query().Get("questions")
	if questions == "" {
		return nil, fmt.Errorf("missing questions parameter")
	}

	tfapikey := r.URL.Query().Get("tfapikey")
	if tfapikey == "" {
		return nil, fmt.Errorf("missing tfapikey parameter")
	}

	return formRequest{
		questions: strings.Split(questions, ","),
		tfapikey:  tfapikey,
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
