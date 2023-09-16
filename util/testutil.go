package util

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func AssertObject(t *testing.T, got, want any) {
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("got differs: (-got +want)\n%s", diff)
	}
}

func AssertJSON(t *testing.T, got, want []byte) {
	t.Helper()

	var jw, jg any
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("cannot unmarshal want %q: %v", want, err)
	}

	if err := json.Unmarshal(got, &jg); err != nil {
		t.Fatalf("cannot unmarshal got %q: %v", got, err)
	}

	if diff := cmp.Diff(jg, jw); diff != "" {
		t.Errorf("got differs: (-got +want)\n%s", diff)
	}
}

func FixedTime() time.Time {
	return time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
}
