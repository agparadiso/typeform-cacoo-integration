package form

import (
	"encoding/json"
	"testing"
)

func TestBuild(t *testing.T) {
	expectedfd := `{"title":"Feedback on my Diagram","settings":{"language":"en","is_public":true},"fields":[{"title":"whats your name","type":"long_text","attachment":{"type":"image","href":"diagramURL"}},{"title":"whats your number","type":"long_text","attachment":{"type":"image","href":"diagramURL"}}]}`
	questions := []string{"whats your name", "whats your number"}
	form := New()
	form.Build(questions, "diagramURL", "tfapikey")
	actualfd, _ := json.Marshal(&form)
	if expectedfd != string(actualfd) {
		t.Fatalf("formdef expected: \n %s\ngot:\n %s", expectedfd, string(actualfd))
	}
}
