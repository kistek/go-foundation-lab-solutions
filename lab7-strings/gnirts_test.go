package gnirts

import (
	"testing"
)

func TestReverseText(t *testing.T) {

	const s string = "Python has superior string manipulation."
	const expected string = ".noitalupinam gnirts roirepus sah nohtyP"
	result := ReverseText(s)

	if result != expected {
		t.Errorf("result %q did not match expected %q", result, expected)
	}
}

func TestReverseWords(t *testing.T) {

	const s1 string = "Python has superior string manipulation"
	const d1 string = " "
	const expected1 string = "manipulation string superior has Python"

	result1 := ReverseWords(s1, d1)
	if result1 != expected1 {
		t.Errorf("result1 %q did not match expected1 %q", result1, expected1)
	}

	const s2 string = "This.is.cool"
	const d2 string = "."
	const expected2 string = "cool.is.This"

	result2 := ReverseWords(s2, d2)
	if result2 != expected2 {
		t.Errorf("result2 %q did not match expected2 %q", result2, expected2)
	}

	const s3 string = ".This.is.cool"
	const d3 string = "."
	const expected3 string = "cool.is.This."

	result3 := ReverseWords(s3, d3)
	if result3 != expected3 {
		t.Errorf("result3 %q did not match expected3 %q", result3, expected3)
	}

}
