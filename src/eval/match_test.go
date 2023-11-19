package eval_test

import (
	"testing"
)

func TestMatchStmt(t *testing.T) {
	code := `
x := 11
match (x) {
    case 7: {
        println("seven")
    }
	case 11: {
	    println("eleven")
	}
	default: {
	    println("default")
	}
}
`
	expected := "eleven\n"
	GeneralTest(false, t, code, expected)
}
