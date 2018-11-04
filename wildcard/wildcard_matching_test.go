package problems

import (
	"testing"
)

func TestMatch(t *testing.T) {
	cases := map[string]map[string]bool{
		"aa": {
			"a":  false,
			"*":  true,
			"a?": true,
			"?a": true,
			"??": true,
			"*?": true,
			"?*": true,
		},

		"cb": {
			"?a":  false,
			"c?b": false,
		},
		"adceb": {
			"*a*b": true,
		},
		"acdcb": {
			"a*c?b": false,
		},
		"aaaa": {
			"***a": true,
		},
		"ho": {
			"ho**": true,
		},
		"aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba": {
			"a*******b": false,
		},
		"baababaababaaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaabaaaaaa": {
			"*aa*ba*a*aa*b*a*aaaaaa*aaaaa*": false,
		},
	}

	for s, ps := range cases {
		for p, expect := range ps {
			got := Match(s, p)
			if got != expect {
				t.Errorf("expected Match(\"%s\", \"%s\") == %t; got %t", s, p, expect, got)
			}
		}
	}
}
