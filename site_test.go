package ephemeris

import (
	"strings"
	"testing"
)

// Test we can load/parse our demo-site
func TestDemoSite(t *testing.T) {

	x := New("_demo/data", "_demo/comments")

	//
	// Blog entries
	//
	url := "https://steve.kemp.example.com/"

	entries, err := x.Entries(url)
	if err != nil {
		t.Errorf("unexpected error getting entries %s", err.Error())
	}

	//
	// ensure each entry has the prefix
	//
	for _, ent := range entries {
		if !strings.HasPrefix(ent.Link, url) {
			t.Errorf("blog entry doesn't have our url-prefix")
		}
	}
}
