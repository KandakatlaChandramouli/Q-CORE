package version

import "testing"

func TestGet(t *testing.T) {
	info := Get()

	if info.GoVersion == "" {
		t.Fatal("GoVersion should not be empty")
	}

	if info.GOOS == "" {
		t.Fatal("GOOS should not be empty")
	}

	if info.GOARCH == "" {
		t.Fatal("GOARCH should not be empty")
	}
}

func TestString(t *testing.T) {
	s := Get().String()

	if len(s) == 0 {
		t.Fatal("String() returned empty output")
	}
}
