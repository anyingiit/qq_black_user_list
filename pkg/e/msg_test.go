package e

import "testing"

// TestGetMsg is a smoke test that the package compiles and that its lookup
// table resolves the codes it actually defines, plus the fallback for a code
// nobody registered. It stays inside code.go's own numbering: GetMsg/MsgFlags
// is a separate, legacy table whose codes happen to overlap numerically with
// the newer Error codes in code_def.go (both define 20001), so looking up an
// Error's Code here would silently read the wrong table's entry.
func TestGetMsg(t *testing.T) {
	if got := GetMsg(SUCCESS); got != "ok" {
		t.Fatalf("GetMsg(SUCCESS) = %q, want %q", got, "ok")
	}
	if got := GetMsg(ERROR_AUTH_CHECK_TOKEN_FAIL); got != MsgFlags[ERROR_AUTH_CHECK_TOKEN_FAIL] {
		t.Fatalf("GetMsg(ERROR_AUTH_CHECK_TOKEN_FAIL) = %q, want %q", got, MsgFlags[ERROR_AUTH_CHECK_TOKEN_FAIL])
	}
	if got := GetMsg(-99999); got != MsgFlags[ERROR] {
		t.Fatalf("GetMsg(unknown) = %q, want fallback %q", got, MsgFlags[ERROR])
	}
}

// TestErrNoQqInListMessage checks the error text the QQ-lookup handler
// (router/api/v1/blackListEven.go) actually sends to the client. That
// handler calls ResponseForJson(c, nil, e.ErrNoQqInList), which reads the
// Error struct's own Message field directly — it never goes through
// GetMsg/MsgFlags — so this is what a real smoke test of that path needs
// to assert.
func TestErrNoQqInListMessage(t *testing.T) {
	if ErrNoQqInList.Message != "没有找到该QQ用户" {
		t.Fatalf("ErrNoQqInList.Message = %q, want %q", ErrNoQqInList.Message, "没有找到该QQ用户")
	}
}
