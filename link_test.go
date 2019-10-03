package link

import (
	"strings"
	"testing"
)

const testHTML = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
  <a href="/second-link">This is a second link</a>
  <div>
	<a href="/third-link">This is a third link</a>
  </div>
</body>
</html>
`

func TestParseLink(t *testing.T) {
	r := strings.NewReader(testHTML)

	links, err := ParseLinks(r)
	if err != nil {
		t.Error(err)
	}

	t.Run("Parse successful", func(t *testing.T) {
		expected := 3
		got := len(links)

		if got != expected {
			t.Errorf("Expected: %d, Got: %d\n", expected, len(links))
		}
	})

	t.Run("Text is parsed", func(t *testing.T) {
		expected := "This is a second link"
		got := links[1].Text

		if got != expected {
			t.Errorf("Expected: %s, Got: %s\n", expected, got)
		}
	})

	t.Run("Comments are ignored", func(t *testing.T) {
		expected := "dog cat"
		got := links[0].Text

		if got != expected {
			t.Errorf("Expected: %s, Got: %s\n", expected, got)
		}
	})
}
