package videoparser

import "testing"

func TestRemoveFileExtension(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			"Whats.Eating.Gilbert.Grape.1993.720p.BluRay.x264-SiNNERS.mkv",
			"Whats.Eating.Gilbert.Grape.1993.720p.BluRay.x264-SiNNERS",
		},
		{
			"melite-spr-720p-rpk.mkv",
			"melite-spr-720p-rpk",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.input, func(t *testing.T) {
			got := RemoveFileExtension(c.input)
			if got != c.want {
				t.Errorf("RemoveFileExtension(%q) = %q, want %q", c.input, got, c.want)
			}
		})
	}
}
