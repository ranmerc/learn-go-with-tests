package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with only string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string fields",
			Input: struct {
				Name string
				Age  int
			}{
				Name: "Chris",
				Age:  23,
			},
			ExpectedCalls: []string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			expected := test.ExpectedCalls

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("got %v want %v", got, expected)
			}
		})
	}
	// expected := "Chris"
	// var got []string

	// x := struct {
	// 	Name string
	// }{expected}

	// Walk(x, func(input string) {
	// 	got = append(got, input)
	// })

	// if got[0] != expected {
	// 	t.Errorf("got %q want %q", got[0], expected)
	// }
}
