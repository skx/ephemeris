// Package ephemeris holds some minimal code to create a blog.
//
// A blog is largely made up of blog-posts, which are parsed from
// a series of text-files.
//
// Each post will have a small header to include tags, date, title,
// and will be transformed into a simple site.
//
// Compared to `chronicle`, which this package replaces, the biggest
// difference is a lack of support for comments at this time.
package ephemeris

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Ephemeris holds our site structure.
//
// Currently we only have a site-root, which lets us know
// where to get our content from, and a path to the comments.
//
type Ephemeris struct {
	// Root is the source of our posts.
	Root string

	// CommentFiles holds the filenames of comments we've found.
	CommentFiles []string

	// Prefix is the absolute URL prefix for the blog
	Prefix string
}

// New creates a new site object.
func New(directory string, commentPath string) *Ephemeris {

	// Create object
	x := &Ephemeris{Root: directory}

	// Now we can find comments - by reading the given
	// directory and adding each of them.
	comments, err := ioutil.ReadDir(commentPath)
	if err != nil {
		return x
	}

	// For each comment file - save the name.
	for _, f := range comments {

		// By appending
		x.CommentFiles = append(x.CommentFiles, filepath.Join(commentPath, f.Name()))
	}

	return x
}

// Entries returns the blog-entries within a site
//
// The entries are returned in a random-order, and contain a complete
// copy of all the text in the entries.  This means that there is a reasonable
// amount of memory overhead here.
func (e *Ephemeris) Entries(prefix string) ([]BlogEntry, error) {

	// Results
	var results []BlogEntry

	//
	// Save the prefix
	//
	e.Prefix = prefix

	//
	// Find the files
	//
	files, err := ioutil.ReadDir(e.Root)
	if err != nil {
		return results, err
	}

	//
	// For each file, return a structure
	//
	for _, f := range files {

		// Build up the complete path
		path := filepath.Join(e.Root, f.Name())

		// Parse the blog-post
		out, err := NewBlogEntry(path, e)
		if err != nil {
			return results, fmt.Errorf("failed to parse %s - %s", path, err.Error())
		}

		// Store
		results = append(results, out)
	}

	return results, nil
}
