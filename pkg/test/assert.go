package test

import "testing"

func Assert(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
func AssertLen(t testing.TB, got, want int) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
