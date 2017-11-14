package main

import (
	"strings"
	"testing"
)

func TestSign(t *testing.T) {
	//TODO: write unit test cases for sign()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader

	cases := []struct {
		input          string
		signingKey     string
		expectedOutput string
	}{
		{
			input:          "test case 1",
			signingKey:     "password",
			expectedOutput: "RBjN26OdbSi7QlGAhJRNPLtWz14bTXxbnwqdh4uYt9A=",
		},
	}

	for _, c := range cases {
		if output, _ := sign(c.signingKey, strings.NewReader(c.input)); output != c.expectedOutput {
			t.Errorf("%s: got %s but expected %s", c.input, output, c.expectedOutput)
		}
	}
}

func TestVerify(t *testing.T) {
	//TODO: write until test cases for verify()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
}
