package ephemeris

import (
	"strings"
	"testing"
)

// Test reading a file that doesn't exist.
func TestMissingBlog(t *testing.T) {

	// fake-site
	site := New("", "")

	_, err := NewBlogEntry("_test/file/is/not/present", site)
	if err == nil {
		t.Errorf("expected error reading a missing file, got none")
	}
}

// Test reading a valid blog-entry
func TestValidBlog(t *testing.T) {

	// fake-site
	site := New("", "")

	b, err := NewBlogEntry("_test/blog_entry/1.txt", site)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	// Ensure the entry has some expected tags
	if len(b.Tags) != 3 {
		t.Errorf("Blog entry has wrong tags: %v", b.Tags)
	}

}

// Test reading a blog-entry with a bogus date
func TestBogusBlogDate(t *testing.T) {

	// fake-site
	site := New("", "")

	_, err := NewBlogEntry("_test/blog_entry/bogus-date.txt", site)
	if err == nil {
		t.Errorf("we expected an error, but found none")
	}

	// Test it looks sane
	if !strings.Contains(err.Error(), "cannot parse") {
		t.Errorf("the error didn't look like a date-parse failure: %s", err.Error())
	}
}

// Test reading a blog-entry with a bogus header
func TestBogusBlogHeader(t *testing.T) {

	// fake-site
	site := New("", "")

	_, err := NewBlogEntry("_test/blog_entry/unknown_header.txt", site)
	if err == nil {
		t.Errorf("we expected an error, but found none")
	}

	// Test it looks sane
	if !strings.Contains(err.Error(), "unknown header-key") {
		t.Errorf("the error didn't look like a header-key failure: %s", err.Error())
	}
}

// Test the `format` header.
func TestBlogFormat(t *testing.T) {

	// fake-site
	site := New("", "")

	_, err := NewBlogEntry("_test/blog_entry/bogus-format.txt", site)
	if err == nil {
		t.Errorf("we expected an error, but found none")
	}

	// Test the error looks sane
	if !strings.Contains(err.Error(), "unknown entry-format") {
		t.Errorf("the error didn't look like a date-parse failure: %s", err.Error())
	}

	//
	// Now we're looking for Markdown
	//
	var b BlogEntry
	b, err = NewBlogEntry("_test/blog_entry/markdown.txt", site)
	if err != nil {
		t.Errorf("unexpected error parsing markdown file %s", err.Error())
	}

	// Test the entry looks sane
	if !strings.Contains(b.Content, "href=\"https://steve.fi/\"") {
		t.Errorf("the blog body doesn't have our link in it: %s", b.Content)
	}
}
