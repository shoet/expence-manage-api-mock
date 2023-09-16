package util

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/shoet/expense-manage-api-mock/entity"
)

// storeutil test ==================
func Test_storeutil_DeserializeFromJSON_Expense(t *testing.T) {
	testfile := "testdata/storeutil/expense.json"
	cwd, err := GetCurrentDirectory()
	if err != nil {
		t.Fatalf("failed GetCurrentDirectory: %v", err)
	}
	fp := filepath.Join(cwd, testfile)
	expenses := []*entity.Expense{}
	if err := DeserializeFromJSON(fp, &expenses); err != nil {
		t.Fatalf("failed LoadJSON: %v", err)
	}
	if len(expenses) == 0 {
		t.Fatalf("not found expense from JSON file %q", fp)
	}
	got := expenses[0]
	want := &entity.Expense{
		ID:            1,
		Title:         "test",
		Expense:       2000,
		UserID:        1,
		AccountTypeID: 1,
		ImageURL:      "test",
		Created:       FixedTime(),
		Modified:      FixedTime(),
	}
	AssertObject(t, got, want)
}

// fileutil test ===================
func Test_fileutil_GetCurrentDirectory(t *testing.T) {
	want := `expense-manage-api-mock/util`
	got, err := GetCurrentDirectory()
	if err != nil {
		t.Fatalf("failed GetCurrentDirectory: %v", err)
	}
	if !strings.HasSuffix(got, want) {
		t.Errorf("unmatch filepath endswith: got(%s), want(%s)", got, want)
	}
}
