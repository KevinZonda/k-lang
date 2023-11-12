package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/tx"
	"testing"
)

func TestNestedLoop(ts *testing.T) {
	code := `
for (i := 0; i <= 5; i := i + 1) {
    for (j := i; j <= 5; j := j + 1) {
        print(i + ", " + (j + 1) + " | ")
    }
    println()
}`
	expected := `0, 1 | 0, 2 | 0, 3 | 0, 4 | 0, 5 | 0, 6 | 
1, 2 | 1, 3 | 1, 4 | 1, 5 | 1, 6 | 
2, 3 | 2, 4 | 2, 5 | 2, 6 | 
3, 4 | 3, 5 | 3, 6 | 
4, 5 | 4, 6 | 
5, 6 |
`
	if err := tx.IsStdoutAsExpected(func() {
		e := New(parserHelper.Ast(code))
		e.Do()
	}, expected); err != nil {
		ts.Fatal(err)
	}
}
