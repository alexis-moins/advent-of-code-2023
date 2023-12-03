package dayone_test

import (
	"testing"

	dayone "github.com/alexis-moins/advent-of-code-2023/internal/day_one"
)

func TestGetNumberInLine(t *testing.T) {
	got := dayone.GetNumberInLine("9sixsevenz3")
	want := 93

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = dayone.GetNumberInLine("seven1cvdvnhpgthfhfljmnq")
	want = 71

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = dayone.GetNumberInLine("6tvxlgrsevenjvbxbfqrsk4seven")
	want = 67

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = dayone.GetNumberInLine("9zml")
	want = 99

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = dayone.GetNumberInLine("52sevenone")
	want = 51

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
