package pkg

import "testing"

func assert(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
func assertLen(t testing.TB, got, want int) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
