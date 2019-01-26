package form

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	formCreationEndpoint = "https://api.typeform.com/forms"
)

type Form interface {
	Build(questions []string) error
	Push(tfapikey string) (string, error)
}

type Field struct {
	Title     string `json:"title"`
	Fieldtype string `json:"type"`
}

type FormDefinition struct {
	Title    string `json:"title"`
	Settings struct {
		Language string `json:"language"`
		Public   bool   `json:"is_public"`
	} `json:"settings"`
	Fields []Field `json:"fields"`
}

func New() *FormDefinition {
	return &FormDefinition{}
}

func (f *FormDefinition) Build(questions []string) error {
	f.Title = "title"
	f.Settings.Language = "en"
	f.Settings.Public = true
	fields := make([]Field, 0)
	for _, q := range questions {
		f := &Field{
			Title:     q,
			Fieldtype: "short_text",
		}
		fields = append(fields, *f)
	}
	f.Fields = fields
	return nil
}

func (f *FormDefinition) Push(tfapikey string) (string, error) {
	payload, err := json.Marshal(f)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal form definition")
	}
	req, err := http.NewRequest("POST", formCreationEndpoint, bytes.NewBuffer(payload))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tfapikey))
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to post form creation")
	}

	fmt.Println(resp)
	return "", nil
}
