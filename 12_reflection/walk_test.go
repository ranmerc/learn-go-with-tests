package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

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
		{
			Name: "nested fields",
			Input: Person{
				Name:    "Chris",
				Profile: Profile{Age: 23, City: "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				Name:    "Chris",
				Profile: Profile{Age: 23, City: "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{
					Age:  22,
					City: "London",
				},
				{
					Age:  23,
					City: "Edinburgh",
				},
			},
			ExpectedCalls: []string{"London", "Edinburgh"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{
					Age:  22,
					City: "London",
				},
				{
					Age:  23,
					City: "Edinburgh",
				},
			},
			ExpectedCalls: []string{"London", "Edinburgh"},
		},
		{
			Name: "map",
			Input: map[string]string{
				"Cow": "Moo",
				"Cat": "Meow",
			},
			ExpectedCalls: []string{"Moo", "Meow"},
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
}
