package main

import "testing"

/*
Things you can do:
1) Check your coverage with this command:
    go test -cover

2) Get your coverage in the browser with this command:
    go test -coverprofile=coverage.out && go tool cover -html=coverage.out
*/

// table test

type Test struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}

var testValid = Test{
	name:     "valid-data",
	dividend: 100,
	divisor:  10.0,
	expected: 10.0,
	isErr:    false,
}

var testInvalid = Test{
	name:     "invalid-data",
	dividend: 100,
	divisor:  0.0,
	expected: 0.0,
	isErr:    true,
}

var testFive = Test{
	name:     "expect-5",
	dividend: 50,
	divisor:  10.0,
	expected: 5.0,
	isErr:    false,
}

var testFraction = Test{
	name:     "expect-fraction",
	dividend: -1.0,
	divisor:  -777.0,
	expected: 0.0012870013,
	isErr:    false,
}

func TestDivision(t *testing.T) {
	var tests []Test
	tests = append(tests, testValid)
	tests = append(tests, testInvalid)
	tests = append(tests, testFive)
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error, but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}
