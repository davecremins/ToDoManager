package content

import (
	"strings"
	"testing"
)

func TestContentSearchFindsTerm(t *testing.T) {
	content := `TODOs 13/01/2020
================
Do Something in the evening
Write some code

Completed 13/01/2020
====================
Finished a book`

	reader := strings.NewReader(content)
	result := FindSearchStr(reader, int64(len(content)), "Completed")
	t.Logf("Search result: %s", result)
}
