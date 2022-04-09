package main

import "testing"

func TestValidAnagram(t *testing.T) {

	tc := []struct {
		name    string
		s       string
		t       string
		isValid bool
	}{
		{
			name:    "!equal length",
			s:       "lily",
			t:       "lem",
			isValid: false,
		},
		{
			name:    "!equal letters",
			s:       "lily",
			t:       "lele",
			isValid: false,
		},
		{
			name:    "anagram!",
			s:       "lily",
			t:       "yill",
			isValid: true,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			if c.isValid != IsValidAnagram(c.s, c.t) {
				t.Error()
			}
		})
	}
}
