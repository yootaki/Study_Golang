package my_even_test

import (
	"tester/my_even"
	"testing"
)

func TestIsEven(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "test1", input: 1, expected: false},
		{name: "test2", input: 2, expected: true},
		{name: "test3", input: 3, expected: false},
		{name: "test4", input: 4, expected: true},
		{name: "test5", input: 5, expected: false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := my_even.IsEven(c.input); c.expected != actual {
				t.Errorf("want IsEven(%d) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}
