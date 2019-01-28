package typeform

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

type formRequest struct {
	questions   []string
	tfapikey    string
	cacooapikey string
	diagramID   string
}
type formResponse struct {
	Tflink string
	err    error
}

func makeBuildFormEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(formRequest)
		res := formResponse{}
		res.Tflink, res.err = svc.CreateForm(ctx, req.questions, req.tfapikey, req.cacooapikey, req.diagramID)
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

	cacooapikey := r.URL.Query().Get("cacooapikey")
	if cacooapikey == "" {
		return nil, fmt.Errorf("missing cacooapikey parameter")
	}

	diagramID := r.URL.Query().Get("diagramID")
	if cacooapikey == "" {
		return nil, fmt.Errorf("missing diagramID parameter")
	}

	return formRequest{
		questions:   strings.Split(questions, ","),
		tfapikey:    tfapikey,
		cacooapikey: cacooapikey,
		diagramID:   diagramID,
	}, nil
}

func makeEncodeBuildFormResponse() func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		res := response.(formResponse)
		if res.err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf(res.err.Error())
		}

		respayload, err := json.Marshal(res)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("failed to Marshal build form response")
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respayload)
		return nil
	}
}
