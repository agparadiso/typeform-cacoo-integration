package typeform

import (
	"context"

	"github.com/agparadiso/tfcacoo/pkg/form"
	"github.com/pkg/errors"
)

type Service interface {
	CreateForm(ctx context.Context, questions []string, tfapikey string) (string, error)
}

type typeform struct {
	form form.Form
}

func NewService() Service {
	form := form.New()
	return &typeform{
		form: form,
	}
}

func (t *typeform) CreateForm(ctx context.Context, questions []string, tfapikey string) (string, error) {
	err := t.form.Build(questions)
	if err != nil {
		return "", errors.Wrap(err, "Failed to build the form")
	}

	return t.form.Push(tfapikey)
}
