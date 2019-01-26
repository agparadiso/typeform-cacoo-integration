package form

import (
	"encoding/json"
	"testing"
)

func TestBuild(t *testing.T) {
	expectedfd := `{"title":"title","settings":{"language":"en","is_public":true},"fields":[{"title":"whats your name","type":"short_text"},{"title":"whats your number","type":"short_text"}]}`
	questions := []string{"whats your name", "whats your number"}
	form := New()
	form.Build(questions)
	actualfd, _ := json.Marshal(&form)
	if expectedfd != string(actualfd) {
		t.Fatalf("formdef expected: \n %s\ngot:\n %s", expectedfd, string(actualfd))
	}
}
