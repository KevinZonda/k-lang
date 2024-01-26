package eval_test

import (
	"testing"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
)

func TestMatchStmt(t *testing.T) {
	code := `
x := 11
match (x) {
    case 7 {
        println("seven")
    }
	case 11 {
	    println("eleven")
	}
	default {
	    println("default")
	}
}
`
	expected := "eleven\n"
	tester.GeneralTest(false, t, code, expected)
}

func TestMatchStmtWithoutDefault(t *testing.T) {
	code := `
x := 18
match (x) {
    case 7 {
        println("seven")
    }
	case 11 {
        println("eleven")
	}
}
`
	expected := ""
	tester.GeneralTest(false, t, code, expected)
}

func TestMatchStmtDefault(t *testing.T) {
	code := `
x := 19
match (x) {
    case 7 {
        println("seven")
    }
	case 11 {
        println("eleven")
	}
	default {
        println("default")
	}
}
`
	expected := "default\n"
	tester.GeneralTest(false, t, code, expected)
}
