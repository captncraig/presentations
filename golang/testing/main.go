package main

// START OMIT
import (
	"testing"
)

func TestFoo(t *testing.T) {
	x := 7 + 8
	if x != 14 {
		t.Errorf("Wrong sum: %v != 14\n", x)
	}
}

// END OMIT
func main() {
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, []testing.InternalTest{
		testing.InternalTest{F: TestFoo, Name: "TestFoo"},
	}, nil, nil)
}
