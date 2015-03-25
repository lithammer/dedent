package dedent

import (
	"fmt"
	"testing"
)

func TestFindMargin(t *testing.T) {
	m := findMargin("Foo\n\t\t\tBar\n\t\tBaz")
	if m != "\t\t" {
		t.Errorf("Expected margin to be '\t\t'; got '%s'", m)
	}
}

func TestFindMarginMultiline(t *testing.T) {
	m := findMargin(`Foo
			Bar
		Baz
`)
	if m != "\t\t" {
		t.Errorf("Expected margin to be '\t\t'; got '%s'", m)
	}
}

func TestDedentMultilineString(t *testing.T) {
	s1 := `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
		Curabitur justo tellus, facilisis nec efficitur dictum,
		fermentum vitae ligula. Sed eu convallis sapien.
	`
	s2 := `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Curabitur justo tellus, facilisis nec efficitur dictum,
fermentum vitae ligula. Sed eu convallis sapien.
`

	if Dedent(s1) != s2 {
		t.Errorf("expected string '%s'; got '%s'", s2, Dedent(s1))
	}
}

func ExampleDedent() {
	fmt.Println(Dedent(`Lorem ipsum dolor sit amet,
		consectetur adipiscing elit.
		Curabitur justo tellus, facilisis nec efficitur dictum,
		fermentum vitae ligula. Sed eu convallis sapien.`))
	// Output:
	// Lorem ipsum dolor sit amet,
	// consectetur adipiscing elit.
	// Curabitur justo tellus, facilisis nec efficitur dictum,
	// fermentum vitae ligula. Sed eu convallis sapien.
}

func BenchmarkDedent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dedent(`Lorem ipsum dolor sit amet, consectetur adipiscing elit.
		Curabitur justo tellus, facilisis nec efficitur dictum,
		fermentum vitae ligula. Sed eu convallis sapien.`)
	}
}

func BenchmarkFindMargin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMargin(`Lorem ipsum dolor sit amet, consectetur adipiscing elit.
		Curabitur justo tellus, facilisis nec efficitur dictum,
		fermentum vitae ligula. Sed eu convallis sapien.`)
	}
}
