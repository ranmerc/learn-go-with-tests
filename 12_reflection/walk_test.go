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
	t.Run("other types", func(t *testing.T) {
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
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow": "Moo",
			"Cat": "Meow",
		}

		var got []string

		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Meow")
	})

	t.Run("with channels", func(t *testing.T) {
		aChan := make(chan Profile)

		go func() {
			aChan <- Profile{Age: 23, City: "London"}
			aChan <- Profile{Age: 22, City: "Edinburgh"}

			close(aChan)
		}()

		var got []string
		want := []string{"London", "Edinburgh"}

		Walk(aChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
