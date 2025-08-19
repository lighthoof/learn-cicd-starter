package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	got, err := GetAPIKey(http.Header{"Authorization": []string{"asdafdsagdfgsdh"}})
	want := "asdafdsagdfgsdh"
	if err != nil {
		t.Fatalf("Error executing GetAPIKey function : %v", err)
	}
	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Fatal(diff)
	}
}
