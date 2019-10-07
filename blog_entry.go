package ephemeris

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/shurcooL/github_flavored_markdown"
)

// BlogComment is the comment associated with a blog post.
type BlogComment struct {

	// Author holds the name of the comment-submitter
	Author string

	// Body holds the body of the comment.
	Body string

	// Icon is generated from the email-address of the submitter.
	// This will point to a  gravitar link.
	Icon string

	// Link holds any user-submitted URL.
	Link string
}

// BlogEntry holds a single post.
//
// A post has a series of attributes associated with it, as you would
// expect.   There are no real surprises here, except for the `MonthName`,
// `Month` and `Year` attributes which are present solely to ease the
// generation of the site-archive.
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

	// These are used to simplify our archive-logic
	MonthName string
	Month     string
	Year      string
}

// NewBlogEntry creates a new blog object from the contents of the given
// file.
//
// If the file is formatted in Markdown it will be expanded to HTML as
// part of the creation-process.
func NewBlogEntry(path string, site *Ephemeris) (BlogEntry, error) {

	// The structure we'll return
	var result BlogEntry

	// Read the file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return result, err
	}

	// Split by newlines - because we want to process the header
	// specially, splitting it into fields.
	lines := strings.Split(string(bytes), "\n")

	// We'll process the header, and append the body here
	body := ""

	// We start off in the header
	header := true

	// Blog-posts might have `format: markdown`.  If not they'll
	// be treated as HTML.
	markdown := false

	// Compile a regular expression to save redoing this for
	// each header-line
	headerRegex, rErr := regexp.Compile("^([^:]+):(.*)$")
	if rErr != nil {
		return result, rErr
	}

	// Process each line
	for _, line := range lines {

		// If we're in the header ..
		if header {

			// Empty line?  Then header-time is over now.
			if len(line) < 1 {
				header = false
				continue
			}

			// Find the key + value which we expect to see
			// in the header.
			header := headerRegex.FindStringSubmatch(line)

			// If we did then we're good.
			if len(header) == 3 {

				// Save the key + value
				key := header[1]
				val := header[2]

				// Normalize keys & values.
				key = strings.ToLower(key)
				val = strings.TrimSpace(val)

				// Now process known-good keys
				switch key {
				case "date":
					t, err := time.Parse("02/01/2006 15:04", val)
					if err != nil {
						return result, err
					}
					result.Date = t

					//
					// These fields are calcuated
					// solely to simplify the construction
					// of the archive-pages.
					//
					_, m, _ := t.Date()
					result.MonthName = m.String()
					result.Month = fmt.Sprintf("%02d", int(m))
					result.Year = fmt.Sprintf("%v", t.Year())

				case "title", "subject":
					result.Title = val
				case "format":
					if val == "markdown" {
						markdown = true
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
			} else {
				return result, fmt.Errorf("malformed header '%s' in %s", line, path)
			}
		} else {
			body += line + "\n"
		}
	}

	//
	// Is this markdown?  If so convert it.
	//
	if markdown {
		body = string(github_flavored_markdown.Markdown([]byte(body)))
	}

	//
	// Otherwise update the object some more.
	//
	result.Path = path
	result.Content = body

	//
	// Make a Regex to say we only want letters and numbers
	//
	reg, err := regexp.Compile("[^a-zA-Z0-9]")
	if err != nil {
		return result, err
	}

	//
	// TODO: Move this elsewhere.
	//
	link := reg.ReplaceAllString(result.Title, "_") + ".html"
	link = strings.ToLower(link)
	result.Link = "https://blog.steve.fi/" + link

	//
	// Add any comments to the appropriate entry
	//
	for _, comment := range site.CommentFiles {

		//
		// We rely upon the fact that the naming
		// scheme of the comments matches the
		// title(s) of the post(s)
		//
		if strings.Contains(comment, link) {

			//
			// Read the blog-comment
			//
			x, err := readComment(comment)
			if err != nil {
				return result, err
			}
			result.CommentData = append(result.CommentData, x)
		}
	}

	//
	// And return
	//
	return result, nil
}

// readComment reads a comment from the named file, and returns that
// in a structured form.
func readComment(path string) (BlogComment, error) {

	// The comment we expect to read.
	var result BlogComment

	// Read the file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return result, err
	}

	// Split by newlines - because we want to process the header
	lines := strings.Split(string(bytes), "\n")

	// We'll process the header, and append the body here
	body := ""

	// We start off in the header
	header := true

	// Compile a regular expression to save redoing this for
	// each header-line
	headerRegex, rErr := regexp.Compile("^([^:]+):(.*)$")
	if rErr != nil {
		return result, rErr
	}

	// Process each line
	for _, line := range lines {

		// If we're in the header ..
		if header {

			// Empty line?  Then header-time is over
			if len(line) < 1 {
				header = false
				continue
			}

			// Find the key + value
			header := headerRegex.FindStringSubmatch(line)

			// If we did then we're good
			if len(header) == 3 {

				// Save the key + value
				key := header[1]
				val := header[2]

				// Cleanup
				key = strings.ToLower(key)
				val = strings.TrimSpace(val)

				// Now process known-good keys
				switch key {
				case "name":
					result.Author = val
				case "mail":
					m := strings.ToLower(val)
					h := fmt.Sprintf("%x", md5.Sum([]byte(m)))
					result.Icon = "//www.gravatar.com/avatar.php?gravatar_id=" + h + ";size=32"

				case "link":
					result.Link = val
				default:
					//			fmt.Printf("Unknown header-key %s in file %s", key, path)
				}
			} else {
				//		fmt.Printf("malformed header '%s' in %s", line, path)
			}
		} else {
			body += line + "\n"
		}
	}

	result.Body = body

	return result, nil
}
