package eval_test

import "testing"

func TestArrayInitializer(t *testing.T) {
	code := `x := [1, 2, 3, 4, 5]
print(x)`
	expected := `[1 2 3 4 5]`
	GeneralTest(false, t, code, expected)
}

func TestMapInitializer(t *testing.T) {
	code := `x := map{
    1 : "one",
    13: "Thirteen"
}
print(x)`
	expected := `map[1:one 13:Thirteen]`
	GeneralTest(true, t, code, expected)
}
