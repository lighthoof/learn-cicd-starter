package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"simple": {input: http.Header{"Authorization": []string{"ApiKey asdafdsagdfgsdh"}}, want: "asdafdsagdfgsdh"},
	}

	//got, err := GetAPIKey(http.Header{"Authorization": []string{"asdafdsagdfgsdh"}})
	//want := "ApiKey asdafdsagdfgsdh"

	for name, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil {
			t.Fatalf("%v: Error executing GetAPIKey function : %v", name, err)
		}
		diff := cmp.Diff(tc.want, got)
		if diff != "" {
			t.Fatalf("%v: %v", name, diff)
		}
	}

}
