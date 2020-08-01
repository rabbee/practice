package three_steps_problem_lcci

import "testing"

func TestWaysToStep(t *testing.T) {
	if waysToStep(3) != 4 {
		t.Fail()
	}
	if waysToStep(5) != 13 {
		t.Fail()
	}
}
