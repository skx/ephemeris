package ephemeris

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/skx/headerfile"
)

// BlogComment is the structure to describe a comment associated with a
// blog post.
type BlogComment struct {

	// Author holds the name of the comment-submitter
	Author string

	// Body holds the body of the comment.
	Body string

	// Icon is generated from the email-address of the submitter.
	// This will be a gravitar link.
	Icon string

	// Link holds any user-submitted URL.
	Link string
}

// NewBlogComment reads a comment from the named file, and returns that
// in a structured form.
func NewBlogComment(path string) (BlogComment, error) {

	// The result
	var result BlogComment

	// Create a helper to read the comment-file.
	reader := headerfile.New(path)

	// Read the headers from the comment-file.
	headers, err := reader.Headers()
	if err != nil {
		return result, err
	}

	// Read the body from the comment-file.
	body := ""
	body, err = reader.Body()
	if err != nil {
		return result, err
	}

	//
	// Now process known-good keys
	//
	// The comment-files I have are a little random/ad-hoc,
	// so we ignore the keys from previous-systems.
	//
	for key, val := range headers {

		switch key {
		case "name":
			result.Author = val
		case "mail":
			m := strings.ToLower(val)
			h := fmt.Sprintf("%x", md5.Sum([]byte(m)))
			result.Icon = "//www.gravatar.com/avatar.php?gravatar_id=" + h + ";size=32"

		case "link":
			result.Link = val

		}
	}

	//
	// Save the body and return the object.
	//
	result.Body = body

	return result, nil
}
