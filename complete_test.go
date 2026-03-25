package videoparser

import "testing"

func TestIsComplete(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"Cars.2.DVDR-TGP", true},
		{"Cars.2.2011.EN.SE.FI.PAL.DVDR-AMIRITE", true},
		{"The.Outsiders.1983.DUAL.COMPLETE.BLURAY-THEORY", true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := isComplete(c.input)
			if got != c.want {
				t.Errorf("isComplete(%q) = %v, want %v", c.input, got, c.want)
			}
		})
	}
}
