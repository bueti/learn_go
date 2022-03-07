package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectString := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectString(t, got, want)
	})
	t.Run("repeat 10 times", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"
		assertCorrectString(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
