package form

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	formCreationEndpoint      = "https://api.typeform.com/forms"
	imageUploadEndpoint       = "https://api.typeform.com/images"
	cacooDiagramImageEndpoint = "https://cacoo.com/api/v1/diagrams"
)

type Form interface {
	UploadImage(diagramID, cacooapikey, tfapikey string) (string, error)
	Build(questions []string, diagramURL, tfapikey string) error
	Push(tfapikey string) (string, error)
}

type Attachment struct {
	Type string `json:"type"`
	Href string `json:"href"`
}

type Field struct {
	Title      string     `json:"title"`
	Fieldtype  string     `json:"type"`
	Attachment Attachment `json:"attachment"`
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

func (f *FormDefinition) Build(questions []string, diagramURL, tfapikey string) error {
	f.Title = "Feedback on my Diagram"
	f.Settings.Language = "en"
	f.Settings.Public = true
	fields := make([]Field, 0)

	for _, q := range questions {
		f := &Field{
			Title:     q,
			Fieldtype: "long_text",
			Attachment: Attachment{
				Type: "image",
				Href: diagramURL,
			},
		}
		fields = append(fields, *f)
	}
	f.Fields = fields
	return nil
}

func (f *FormDefinition) Push(tfapikey string) (string, error) {
	type formcreationresp struct {
		Link struct {
			Display string `json:"display"`
		} `json:"_links"`
	}

	payload, err := json.Marshal(f)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal form definition")
	}

	req, err := http.NewRequest("POST", formCreationEndpoint, bytes.NewBuffer(payload))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tfapikey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to post form creation")
	}
	defer res.Body.Close()

	formcreated := formcreationresp{}
	err = json.NewDecoder(res.Body).Decode(&formcreated)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode form creation response")
	}
	return formcreated.Link.Display, nil
}

func (f *FormDefinition) UploadImage(diagramID, cacooapikey, tfapikey string) (string, error) {
	type uploadReq struct {
		URL      string `json:"url"`
		Filename string `json:"file_name"`
	}

	type uploadres struct {
		Src string `json:"src"`
	}

	diagramURL := fmt.Sprintf("%s/%s.png?apiKey=%s", cacooDiagramImageEndpoint, diagramID, cacooapikey)
	ur := &uploadReq{URL: diagramURL, Filename: fmt.Sprintf("%s.png", diagramID)}
	payload, err := json.Marshal(&ur)
	if err != nil {
		return "", fmt.Errorf("failed to marshal upload request")
	}

	req, err := http.NewRequest("POST", imageUploadEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create upload request")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tfapikey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to tf")
	}
	defer res.Body.Close()

	imgURL := uploadres{}
	err = json.NewDecoder(res.Body).Decode(&imgURL)
	if err != nil {
		return "", fmt.Errorf("failed to decode fileupload response")
	}

	return imgURL.Src, nil
}
