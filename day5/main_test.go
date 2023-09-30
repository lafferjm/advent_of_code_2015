package main

import "testing"

type containsTest struct {
	arg      string
	expected bool
}

func TestContainsPairOfLetters(t *testing.T) {
	input := "xxyxx"
	containsPair := containsPairOfLetters(input)
	if !containsPair {
		t.Fatalf("Input does not contain pair, want true got false")
	}
}

func TestContainsRepeatOfLetter(t *testing.T) {
	input := "xxyxx"
	containsRepeat := containsRepeatOfLetter(input)
	if !containsRepeat {
		t.Fatalf("Input does not contain repeat, want true got false")
	}
}

var containsTests = []containsTest{
	{"qjhvhtzxzqqjkmpb", true},
	{"xxyxx", true},
	{"aaaa", true},
	{"aaa", false},
	{"uurcxstgmygtbstg", false},
	{"ieodomkazucvgmuy", false},
}

func TestContainsRepeatAndPair(t *testing.T) {
	for _, test := range containsTests {
		isValid := containsRepeatOfLetter(test.arg) && containsPairOfLetters(test.arg)
		if isValid != test.expected {
			t.Errorf("Output %t not equal to expected %t for input %s", isValid, test.expected, test.arg)
		}
	}
}
