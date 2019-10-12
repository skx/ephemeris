package ephemeris

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/shurcooL/github_flavored_markdown"
	"github.com/skx/headerfile"
)

// BlogEntry holds a single post.
//
// A post has a series of attributes associated with it, as you would
// expect, such as a title a set of tags, and an associated set of comments.
type BlogEntry struct {
	// Title holds the blog-title
	Title string

	// Path holds the path to the source-file, on-disk
	Path string

	// Tags contains a list of tags for the given post.
	Tags []string

	// Content contains the post-body
	Content string

	// The link to the post
	Link string

	// Date is when the post was created.
	Date time.Time

	// CommentData contains any blog-comments upon this entry
	CommentData []BlogComment
}

// NewBlogEntry creates a new blog object from the contents of the given
// file.
//
// If the file is formatted in Markdown it will be expanded to HTML as
// part of the creation-process.
func NewBlogEntry(path string, site *Ephemeris) (BlogEntry, error) {

	// The structure we'll return
	var result BlogEntry

	// Create a helper to read the entry.
	reader := headerfile.New(path)

	// Read the headers from the post
	headers, err := reader.Headers()
	if err != nil {
		return result, err
	}

	// Read the body from the post
	//
	// errors can't happen here, because if they were present
	// they would have happened in the header-read.
	body, _ := reader.Body()

	// Sanity-check the headers
	for key, val := range headers {

		// Now process known-good keys
		switch key {
		case "date":
			t, err := time.Parse("02/01/2006 15:04", val)
			if err != nil {
				return result, err
			}
			result.Date = t

		case "title", "subject":
			result.Title = val
		case "format":
			if val == "markdown" {
				body = string(github_flavored_markdown.Markdown([]byte(body)))
			} else {
				return result, fmt.Errorf("unknown entry-format %s", val)
			}
		case "tags":
			tags := strings.Split(val, ",")
			for _, t := range tags {
				t = strings.TrimSpace(t)
				t = strings.ToLower(t)
				if len(t) >= 1 {
					result.Tags = append(result.Tags, t)
				}
			}
			sort.Strings(result.Tags)
		default:
			return result, fmt.Errorf("unknown header-key %s in file %s", key, path)
		}
	}

	//
	// Otherwise update the object some more.
	//
	result.Path = path
	result.Content = body

	//
	// Make a Regex to say we only want letters and numbers
	//
	reg := regexp.MustCompile("[^a-zA-Z0-9]")

	//
	// Normalise the output
	//
	link := reg.ReplaceAllString(result.Title, "_") + ".html"

	//
	// Make our link absolute.
	//
	result.Link = site.Prefix + link

	//
	// Add any comments to the appropriate entry
	//
	for _, comment := range site.CommentFiles {

		//
		// We rely upon the fact that the naming
		// scheme of the comments matches the
		// title(s) of the post(s)
		//
		if strings.Contains(comment, strings.ToLower(link)) {

			//
			// Read the blog-comment
			//
			x, err := NewBlogComment(comment)
			if err != nil {
				return result, err
			}

			//
			// Append it to our list.
			//
			result.CommentData = append(result.CommentData, x)
		}
	}

	//
	// And return
	//
	return result, nil
}
