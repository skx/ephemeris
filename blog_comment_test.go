package ephemeris

import (
	"strings"
	"testing"
)

// Test reading a file that doesn't exist.
func TestMissing(t *testing.T) {

	_, err := NewBlogComment("_test/file/is/not/present")
	if err == nil {
		t.Errorf("expected error reading a missing file, got none")
	}
}

// Test reading a valid blog-comment
func TestValid(t *testing.T) {

	b, err := NewBlogComment("_test/blog_comment/valid.html.12345")
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if !strings.Contains(b.Body, "example.com") {
		t.Errorf("comment-body didn't contain our link")
	}

	if b.Link != "http://example.net" {
		t.Errorf("link doesn't contain a http-prefix: %s", b.Link)
	}
}

// Test reading a badly-named comment-file
//
// Blog-posts should be "${name}.html.${ctime}".
func TestBadNaming(t *testing.T) {

	_, err := NewBlogComment("_test/blog_comment/invalid.html.html")
	if err == nil {
		t.Errorf("received no error, but wanted one")
	}

	if !strings.Contains(err.Error(), "parsing") {
		t.Errorf("the error we got didn't seem to refer to parsing 'html' as an integer")
	}
}

// Test reading a badly-formed file.
//
// The header should be "key:val" or "key = val".  Anything else is bogus.
func TestBadHeader(t *testing.T) {

	_, err := NewBlogComment("_test/blog_comment/invalid.html.12345")
	if err == nil {
		t.Errorf("received no error, but wanted one")
	}

	if !strings.Contains(err.Error(), "header") {
		t.Errorf("the error we got didn't seem to refer to parsing 'html' as an integer")
	}
}
