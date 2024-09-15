package leetcode

import "testing"

func TestDemo(t *testing.T) {

	checks := []struct {
		a        string
		expacted string
	}{

		{"abba", ""},
		{"abbc", "ac"},
		{"addyl", "ayl"},
	}

	for _, v := range checks {

		get := Demo(v.a)

		if get != v.expacted {

			t.Errorf("We have an error")
		}

	}

}
