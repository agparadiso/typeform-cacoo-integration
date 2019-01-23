package typeform

import (
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler() http.Handler {
	svc := NewService()
	handler := kithttp.NewServer(
		makeBuildFormEndpoint(svc),
		decodeBuildFormRequest,
		makeEncodeBuildFormResponse(),
	)
	return handler
}
