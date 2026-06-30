package numtowords_test

import (
	"testing"

	"github.com/NayanaMitti22/numtowords"
)

func TestInvalidNumber(t *testing.T) {
	_, err := numtowords.Convert(numtowords.MaxNum + 1)
	if err == nil {
		t.Log("expected error")
		t.Fail()
	}

	_, err = numtowords.Convert(numtowords.MinNum - 1)
	if err == nil {
		t.Log("expected error")
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	result, err := numtowords.Convert(0)
	if err != nil {
		t.Log("Expected zero")
		t.FailNow()
	}
	if result != "zero" {
		t.Logf("expected zero, recieved %v", result)
		t.FailNow()
	}
}

func TestUnits(t *testing.T) {
	units := [20]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	for i, v := range units {
		result, err := numtowords.Convert(i)
		if err != nil {
			t.Logf("Error while converting %v: %v", i, err)
			t.Fail()
		}
		if result != units[i] {
			t.Logf("while converting %v, expected %v, got %v", i, v, result)
			t.Fail()
		}
	}
}

func TestTens(t *testing.T) {
	testcses := map[int]string{
		34: "thirty four",
		48: "forty eight",
		60: "sixty",
		18: "eighteen",
	}
	for i, v := range testcses {
		result, err := numtowords.Convert(i)
		if err != nil {
			t.Logf("Error while converting")
			t.Fail()
		}
		if result != v {
			t.Logf("error while testing")
			t.Fail()
		}
	}
}

func TestHundereds(t *testing.T) {
	testcses := map[int]string{
		108: "one hundred and eight",
		333: "three hundred and thirty three",
		700: "seven hundred",
		230: "two hundred and thirty",
	}
	for i, v := range testcses {
		result, err := numtowords.Convert(i)
		if err != nil {
			t.Logf("Error while converting")
			t.Fail()
		}
		if result != v {
			t.Logf("error while testing expected %v got %v", v, result)
			t.Fail()
		}
	}
}

func TestNegative(t *testing.T) {
	testcases := map[int]string{
		-100: "minus one hundred",
		-205: "minus two hundred and five",
		-346: "minus three hundred and forty six",
		-640: "minus six hundred and forty",
	}
	for i, v := range testcases {
		result, err := numtowords.Convert(i)
		if err != nil {
			t.Logf("Error while converting")
			t.Fail()
		}
		if result != v {
			t.Logf("error while testing expected %v got %v", v, result)
			t.Fail()
		}
	}
}
