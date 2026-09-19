package e

import "testing"

// TestGetMsg is the smoke test this repository shipped with none of: the
// package compiles, and the lookup table it defines actually resolves the
// code the QQ lookup handler returns when a number is not on the list
// (router/api/v1/blackListEven.go uses e.ErrNoQqInList), plus the fallback
// for a code nobody registered.
func TestGetMsg(t *testing.T) {
	if got := GetMsg(SUCCESS); got != "ok" {
		t.Fatalf("GetMsg(SUCCESS) = %q, want %q", got, "ok")
	}
	if got := GetMsg(ErrNoQqInList.Code); got != ErrNoQqInList.Message {
		t.Fatalf("GetMsg(ErrNoQqInList.Code) = %q, want %q", got, ErrNoQqInList.Message)
	}
	if got := GetMsg(-99999); got != MsgFlags[ERROR] {
		t.Fatalf("GetMsg(unknown) = %q, want fallback %q", got, MsgFlags[ERROR])
	}
}
