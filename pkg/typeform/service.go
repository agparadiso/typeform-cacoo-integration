package typeform

import (
	"context"
	"fmt"
)

type Service interface {
	BuildForm(ctx context.Context, question string) error
}

type typeform struct{}

func NewService() Service {
	return &typeform{}
}

func (t *typeform) BuildForm(ctx context.Context, question string) error {
	fmt.Println("buildingForm")
	return nil
}
