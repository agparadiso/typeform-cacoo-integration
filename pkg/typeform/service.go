package typeform

import (
	"context"

	"github.com/agparadiso/tfcacoo/pkg/form"
	"github.com/pkg/errors"
)

type Service interface {
	CreateForm(ctx context.Context, questions []string, tfapikey, cacooapikey, diagramID string) (string, error)
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

func (t *typeform) CreateForm(ctx context.Context, questions []string, tfapikey, cacooapikey, diagramID string) (string, error) {
	diagramURL, err := t.form.UploadImage(diagramID, cacooapikey, tfapikey)
	if err != nil {
		return "", errors.Wrap(err, "failed to upload image")
	}

	err = t.form.Build(questions, diagramURL, tfapikey)
	if err != nil {
		return "", errors.Wrap(err, "Failed to build the form")
	}

	return t.form.Push(tfapikey)
}
