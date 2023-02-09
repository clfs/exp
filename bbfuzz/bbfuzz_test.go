package bbfuzz

import (
	"testing"
)

func FuzzRegexp0(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp0(s); got != (s != "bb") {
			t.Errorf("Regexp0(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp1(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp1(s); got != (s != "bb") {
			t.Errorf("Regexp1(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp2(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp2(s); got != (s != "bb") {
			t.Errorf("Regexp2(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp3(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp3(s); got != (s != "bb") {
			t.Errorf("Regexp3(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp4(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp4(s); got != (s != "bb") {
			t.Errorf("Regexp4(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp5(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp5(s); got != (s != "bb") {
			t.Errorf("Regexp5(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp6(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp6(s); got != (s != "bb") {
			t.Errorf("Regexp6(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp7(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp7(s); got != (s != "bb") {
			t.Errorf("Regexp7(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp8(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp8(s); got != (s != "bb") {
			t.Errorf("Regexp8(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp9(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp9(s); got != (s != "bb") {
			t.Errorf("Regexp9(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp10(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp10(s); got != (s != "bb") {
			t.Errorf("Regexp10(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp11(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp11(s); got != (s != "bb") {
			t.Errorf("Regexp11(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp12(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp12(s); got != (s != "bb") {
			t.Errorf("Regexp12(%q) = %t", s, got)
		}
	})
}

func FuzzRegexp13(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		if got := Regexp13(s); got != (s != "bb") {
			t.Errorf("Regexp13(%q) = %t", s, got)
		}
	})
}
