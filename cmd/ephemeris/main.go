// This is a proof of concept project to read/parse all the blog-files which
// I have written, and generate a static-blog.
//
// We use the `ephemeris` package to find the text-files beneath a given
// root-directory, and then interate over them in various ways to build up:
//
// * The blog-entries themselves
//
// * The tag-cloud
//
// * An archive.
//
// * The index & RSS feed.
//
// The way that the system is setup the most recent post will allow
// comments to be submitted upon it - all others will be read-only.
//
package main

import (
	"flag"
	"fmt"
	"html"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/skx/ephemeris"
)

//
// Variables which are reused live here.
//

// We generate a static-version (read:embedded copy) of all the
// template-files beneath `data/`.
// We load each of these as golang templates here, so that they
// are globally available and we don't have to re-read/re-parse them
// more than once.
var tmpl *template.Template

// We load a JSON configuration file when we launch, which contains
// the mandatory settings.  We make this configuration object global
// to access those variables even though that is a bad design.
var config Config

// mkdirIfMissing makes a directory, if it is missing.
//
// The overhead of calling `stat` probably makes it cheaper to just
// always call `mkdir` and ignore the error, but this is cleaner.
func mkdirIfMissing(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}

// getRecentPosts fetches the most recent N posts from our collection
// of entries.  It is used to render the partial/sidebar
//
// We moved this routine into its own function for cleanliness, although
// it is only called once.
func getRecentPosts(posts []ephemeris.BlogEntry, count int) []ephemeris.BlogEntry {

	// The return-value
	var recent []ephemeris.BlogEntry

	// If there aren't too many posts then return them as-is.
	if len(posts) <= count {
		return posts
	}

	// Sort the list of posts by date.
	sort.Slice(posts, func(i, j int) bool {
		a := posts[i].Date
		b := posts[j].Date
		return a.Before(b)
	})

	// Build up the given number of posts.
	c := 0
	for c < count {
		ent := posts[len(posts)-1-c]
		recent = append(recent, ent)
		c++
	}

	// All done
	return recent
}

// loadTemplates returns a collection of all the templates we have
// stored within our `static.go` file.
//
// In addition to loading the templates we also populate a function-map,
// to allow various functions to be made available to all templates.
//
// The functions defined are:
//
// ISO8601          - Needed for RSS generation.
// LOWER            - Lower-case a string.  Used for link-generation.
// ESCAPE           - Escape HTML-text for RSS_generation too.
// RECENT_POST_DATE - The date format used for the "most recent entries" list in the sidebar.
// BLOG_POST_DATE   - The format used in the index/archive/tag-view.
//
func loadTemplates() *template.Template {

	// Create a helper-template, with no name.
	t := template.New("").Funcs(template.FuncMap{

		// Date-format for RSS feed
		"ISO8601": func(d time.Time) string {
			return (fmt.Sprintf("%v", d.Format(time.RFC3339)))
		},

		// Escape HTML in RSS feed
		"ESCAPE": func(in string) string {
			return (html.EscapeString(in))
		},

		// Convert a string to lower-case.
		//
		// This is used to make sure all links point to their
		// lower-cased version of the URL.
		//
		"LOWER": func(in string) string {
			return (strings.ToLower(in))
		},

		// Date used on "recent posts"
		"RECENT_POST_DATE": func(d time.Time) string {
			year, month, day := d.Date()
			return (fmt.Sprintf("%d %s %d", day, month.String(), year))
		},

		// Date used on all blog posts.
		"BLOG_POST_DATE": func(d time.Time) string {
			year, month, day := d.Date()
			return (fmt.Sprintf("%d %s %d %02d:%02d", day, month.String(), year, d.Hour(), d.Minute()))
		},

		// Date used on comments.
		"COMMENT_POST_DATE": func(d time.Time) string {
			year, month, day := d.Date()
			return (fmt.Sprintf("at %02d:%02d on %d %s %d", d.Hour(), d.Minute(), day, month.String(), year))
		},
	})

	// Add in each static-file
	for _, entry := range getResources() {

		// Get the data.
		data, err := getResource(entry.Filename)
		if err != nil {
			fmt.Printf("Error getting resource %s - %s\n", entry.Filename, err.Error())
		}

		// Add the data
		t = t.New(entry.Filename)
		t, err = t.Parse(string(data))
		if err != nil {
			fmt.Printf("Error parsing template resource %s - %s\n", entry.Filename, err.Error())
		}
	}

	return t
}

// outputTags writes out the tag-specific pages.
//
// First of all build up a list of tags, then render
// a template for each one.
func outputTags(posts []ephemeris.BlogEntry, recentPosts []ephemeris.BlogEntry) error {

	//
	// OK we'll now try to build up a list of tags.
	//
	// The tag list will consist of:
	//
	//   tagName -> [array-index-1, array-index-2, ...]
	//
	// Where the array indexes are the indexes of the posts
	// array which contain the tag.
	//
	// We do this because it is lightweight.
	//
	tagMap := make(map[string][]int)

	//
	// For each post ..
	//
	for i, e := range posts {

		// If it has tags ..
		if len(e.Tags) > 0 {

			// For each tag
			for _, tag := range e.Tags {

				// Add the entry-number to the tag-list
				existing := tagMap[tag]
				existing = append(existing, i)
				tagMap[tag] = existing
			}
		}
	}

	//
	//  Page-Structure for a tag-page view.
	//
	//  i.e. /tags/z80
	//
	type TagPage struct {
		Entries     []ephemeris.BlogEntry
		Tag         string
		Prefix      string
		RecentPosts []ephemeris.BlogEntry
	}
	var pageData TagPage
	pageData.RecentPosts = recentPosts
	pageData.Prefix = config.Prefix

	//
	// Create a per-page tag-template
	//
	for key, uses := range tagMap {

		mkdirIfMissing(filepath.Join(config.OutputPath, "tags", key))

		// Empty the tags from the previous run
		pageData.Entries = nil
		pageData.Tag = key

		// Add the entries
		for _, e := range uses {
			pageData.Entries = append(pageData.Entries, posts[e])
		}

		// Sort by date - tags will be viewed in creation-order
		sort.Slice(pageData.Entries, func(i, j int) bool {
			a := pageData.Entries[i].Date
			b := pageData.Entries[j].Date
			return a.Before(b)
		})

		output, err := os.Create(filepath.Join(config.OutputPath, "tags", key, "index.html"))
		if err != nil {
			return err
		}

		// Render the template
		err = tmpl.ExecuteTemplate(output, "data/tag_page.tmpl", pageData)
		if err != nil {
			return err
		}
		output.Close()
	}

	//
	// /tags/index.html
	//

	//
	// We want a sorted list of tags
	//
	var tagNames []string
	for key := range tagMap {
		tagNames = append(tagNames, key)
	}
	sort.Strings(tagNames)

	//
	// We want to have a tag-cloud for the /tags/index.html page.
	//
	type TagMap struct {
		// Tag contains the name of the string.
		Tag string

		// TSize contains the text-size of the entry in the cloud.
		TSize int

		// Count shows how many times the tag was used.
		Count int
	}

	//
	// Page-data: The tag-data, and the recent-posts list.
	//
	type TagCloudPage struct {
		Tags        []TagMap
		Prefix      string
		RecentPosts []ephemeris.BlogEntry
	}
	var tagCloud TagCloudPage
	tagCloud.RecentPosts = recentPosts
	tagCloud.Prefix = config.Prefix

	//
	// Now we have a sorted list of unique tag-names we can build up
	// that array for the template-page
	//
	for _, tag := range tagNames {
		count := len(tagMap[tag])
		size := (count * 5) + 5
		if size > 60 {
			size = 60
		}
		tagCloud.Tags = append(tagCloud.Tags, TagMap{Tag: tag, Count: count, TSize: size})
	}

	//
	// Populate the page
	//
	ti, err := os.Create(filepath.Join(config.OutputPath, "tags", "index.html"))
	if err != nil {
		return err
	}
	defer ti.Close()

	// Render the template
	err = tmpl.ExecuteTemplate(ti, "data/tags.tmpl", tagCloud)
	if err != nil {
		return err
	}

	return nil
}

// output a year/month page for each distinct period
// in which we have posts.
func outputArchive(posts []ephemeris.BlogEntry, recentPosts []ephemeris.BlogEntry) error {

	//
	// We'll build up a list of year/mon pages.
	//
	// The map will consist of:
	//
	//   year/mon -> [array-index-1, array-index-2, ...]
	//
	// Where the array indexes are the indexes of the posts
	// array which were posted in the given month.
	//
	archiveMap := make(map[string][]int)

	for i, e := range posts {

		// The key is "YYYY/NN"
		key := e.Year() + "/" + e.MonthNumber()

		existing := archiveMap[key]
		existing = append(existing, i)
		archiveMap[key] = existing
	}

	//
	// Archive page contains data to be shown in
	// the archive page of:
	//
	//  /archive/2019/03/
	//
	type PageData struct {
		Entries []ephemeris.BlogEntry
		Year    string
		Month   string

		Prefix      string
		RecentPosts []ephemeris.BlogEntry
	}
	var pageData PageData
	pageData.RecentPosts = recentPosts
	pageData.Prefix = config.Prefix

	//
	// Create a per-page output
	//
	for key, uses := range archiveMap {

		mkdirIfMissing(filepath.Join(config.OutputPath, "archive", key))

		// Empty the tags from the previous run
		pageData.Entries = nil

		// Add the entries
		for _, e := range uses {
			pageData.Entries = append(pageData.Entries, posts[e])

			//
			// Our archiveMap contains keys of the form:
			//
			//    year/mon
			//
			// But when we present this to the viewers we want
			// to show the month-name, and year name.
			//
			// We can calculate that from the `Date` field
			// in the post itself :)
			//
			pageData.Year = posts[e].Year()
			pageData.Month = posts[e].MonthName()
		}

		// Sort by date - posts will be in order they've been written
		sort.Slice(pageData.Entries, func(i, j int) bool {
			a := pageData.Entries[i].Date
			b := pageData.Entries[j].Date
			return a.Before(b)
		})

		output, err := os.Create(filepath.Join(config.OutputPath, "archive", key, "index.html"))
		if err != nil {
			return err
		}

		// Render the template
		err = tmpl.ExecuteTemplate(output, "data/archive_page.tmpl", pageData)
		if err != nil {
			return err
		}
		output.Close()
	}

	//
	// Page data for the archive-index.
	//
	//  i.e. /archive/index.html
	//
	type ArchiveCount struct {
		Year      string
		Month     string
		MonthName string
		Count     string
	}

	mappy := make(map[string][]ArchiveCount)

	type ArchiveIndex struct {
		Year        string
		Prefix      string
		Data        []ArchiveCount
		RecentPosts []ephemeris.BlogEntry
	}
	var ai []ArchiveIndex

	//
	// Build up the count of posts in the given month/year.
	//
	for _, e := range archiveMap {

		//
		// Since all the posts are have the same year + month
		// pair we're able to just use the first entry in
		// each returned set.
		//
		y := posts[e[0]].Year()
		m := posts[e[0]].MonthNumber()
		n := posts[e[0]].MonthName()
		c := fmt.Sprintf("%d", len(e))

		// Append
		existing := mappy[y]
		existing = append(existing, ArchiveCount{
			Count:     c,
			Year:      y,
			Month:     m,
			MonthName: n,
		})

		mappy[y] = existing
	}

	// Sort the entries now we've generated them.
	for k, entries := range mappy {

		vals := mappy[k]
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Month < entries[j].Month
		})
		mappy[k] = vals

	}

	//
	// Now a data-structure
	//
	// sorted keys
	//
	var years []string
	for year := range mappy {
		years = append(years, year)
	}
	sort.Strings(years)

	//
	// For each year we add the data
	//
	for _, year := range years {
		ai = append(ai, ArchiveIndex{Year: year, Data: mappy[year], RecentPosts: recentPosts, Prefix: config.Prefix})
	}

	// Create the output file.
	ar, err := os.Create(filepath.Join(config.OutputPath, "archive", "index.html"))
	if err != nil {
		return err
	}

	// Render the template
	err = tmpl.ExecuteTemplate(ar, "data/archive.tmpl", ai)
	if err != nil {
		return err
	}
	ar.Close()
	return nil
}

// outputIndex outputs the /index.html file.
//
// We don't need to sort, or limit ourselves here, because we only use
// the "most recent posts" we've already discovered.
//
func outputIndex(posts []ephemeris.BlogEntry, recentPosts []ephemeris.BlogEntry) error {

	mkdirIfMissing(config.OutputPath)

	// Page-structure for the site.
	type Recent struct {

		// Entries has the most recent entries.
		Entries []ephemeris.BlogEntry

		// RecentPosts has the same data, but for
		// the side-bar.  It is redundant.
		RecentPosts []ephemeris.BlogEntry

		// Prefix to the site
		Prefix string
	}

	//
	// The data we'll store for the page.
	//
	// Our front-page shows the same number of posts as
	// the recent-list in the sidebar, so we don't need
	// to do anything special here, we show the same
	// list for both of them.
	//
	var pageData Recent
	pageData.Entries = recentPosts
	pageData.RecentPosts = recentPosts
	pageData.Prefix = config.Prefix

	// Create the output file.
	output, err := os.Create(filepath.Join(config.OutputPath, "index.html"))
	if err != nil {
		return err
	}

	// Render the template
	err = tmpl.ExecuteTemplate(output, "data/index.tmpl", pageData)
	if err != nil {
		return err
	}
	output.Close()

	//
	// Create the output file.
	//
	rss, err := os.Create(filepath.Join(config.OutputPath, "index.rss"))
	if err != nil {
		return err
	}

	//
	// Render the RSS template too, with the same data
	//
	err = tmpl.ExecuteTemplate(rss, "data/index.rss", pageData)
	if err != nil {
		return err
	}
	rss.Close()

	return nil

}

// outputRSS outputs the /index.rss file.
//
// We don't need to sort, or limit ourselves here, because we only use
// the "most recent posts" we've already discovered.
//
func outputRSS(posts []ephemeris.BlogEntry, recentPosts []ephemeris.BlogEntry) error {

	mkdirIfMissing(config.OutputPath)

	// Page-structure for the site.
	type Recent struct {

		// Entries has the most recent entries.
		Entries []ephemeris.BlogEntry

		// RecentPosts has the same data, but for
		// the side-bar.  It is redundant.
		RecentPosts []ephemeris.BlogEntry

		// Prefix to the site
		Prefix string
	}

	//
	// The data we'll store for the page.
	//
	// Our front-page shows the same number of posts as
	// the recent-list in the sidebar, so we don't need
	// to do anything special here, we show the same
	// list for both of them.
	//
	var pageData Recent
	pageData.Entries = recentPosts
	pageData.RecentPosts = recentPosts
	pageData.Prefix = config.Prefix

	//
	// Create the output file.
	//
	rss, err := os.Create(filepath.Join(config.OutputPath, "index.rss"))
	if err != nil {
		return err
	}

	//
	// Render the RSS template too, with the same data
	//
	err = tmpl.ExecuteTemplate(rss, "data/index.rss", pageData)
	if err != nil {
		return err
	}
	rss.Close()

	return nil

}

// Output one page for each entry.
func outputEntries(posts []ephemeris.BlogEntry, recentPosts []ephemeris.BlogEntry) error {

	mkdirIfMissing(config.OutputPath)

	// Page-structure for the site.
	type Recent struct {

		// The blog-entry
		Entry ephemeris.BlogEntry

		// Should we display the add-comment form for this post?
		AddComment bool

		// Prefix for the site
		Prefix string

		// The recent posts for the sidebar.
		RecentPosts []ephemeris.BlogEntry
	}

	//
	// The data we use for output.
	//
	var pageData Recent
	pageData.RecentPosts = recentPosts
	pageData.AddComment = false
	pageData.Prefix = config.Prefix

	//
	// Create a per-page output
	//
	for _, entry := range posts {

		//
		// Populate the page-data with this entry.
		//
		pageData.Entry = entry

		//
		// The most recent post has comments enabled,
		// all others do not.
		//
		pageData.AddComment = config.AddComments && (entry.Path == recentPosts[0].Path)

		//
		// We have a link and that points to a filename.
		//
		// We get the latter from the former by removing the
		// prefix.
		//
		u, err := url.Parse(entry.Link)
		if err != nil {
			return err
		}

		// Get the path.
		path := u.RequestURI()

		// Remove the leading slash.
		path = strings.TrimPrefix(path, "/")

		//
		// Lower-case PATH and write to that too
		//
		dest := strings.ToLower(path)

		//
		// Create the output file.
		//
		output, err := os.Create(filepath.Join(config.OutputPath, dest))
		if err != nil {
			return err
		}

		//
		// Render the template into it.
		//
		err = tmpl.ExecuteTemplate(output, "data/entry.tmpl", pageData)
		if err != nil {
			return err
		}
		output.Close()

		//
		// Create symlink
		//
		os.Symlink(dest, filepath.Join(config.OutputPath, path))

	}

	return nil

}

// main is our entry-point.
func main() {

	//
	// Command-line arguments which are accepted.
	//
	confFile := flag.String("config", "ephemeris.json", "The path to our configuration file.")
	allowComments := flag.Bool("allow-comments", true, "Enable comments to be added to the most recent entry.")

	//
	// Parse the flags.
	//
	flag.Parse()

	//
	// Record our start-time
	//
	start := time.Now()

	//
	// Load our configuration file (JSON)
	//
	var err error
	config, err = loadConfig(*confFile)
	if err != nil {
		fmt.Printf("Failed to load configuration file %s %s\n", err.Error())
		return
	}

	//
	// Setup defaults if missing
	//
	if config.OutputPath == "" {
		config.OutputPath = "output"
	}
	if config.Posts == "" {
		config.Posts = "data/"
	}
	if config.Comments == "" {
		config.Comments = "comments/"
	}

	//
	// Preserve comment setting
	//
	config.AddComments = *allowComments

	//
	// Create an object to generate our blog from
	//
	site := ephemeris.New(config.Posts, config.Comments)

	//
	// Parse all the blog-posts in the site.
	//
	entries, err := site.Entries(config.Prefix)
	if err != nil {
		fmt.Printf("Failed to read blog-entries: %s\n", err.Error())
		return
	}

	//
	// We can now load the collection of templates which we've stored
	// in `static.go`.
	//
	// Our templates are loaded en masse, and each one of them
	// has some (custom/bonus/extra) functions available to them.
	//
	tmpl = loadTemplates()

	//
	// We have a list of "recent" posts, which will be embedded in
	// many of our output-pages, as a partial.  (i.e. sidebar)
	//
	// To avoid processing the entries too many times we'll get that
	// list once here.
	//
	recent := getRecentPosts(entries, 10)

	//
	// We're going to run the page-generation in a series of threads
	// now.  So we'll add a synchronizer here.
	//
	var wg sync.WaitGroup

	//
	// Ensure we use all the CPU we have available.
	//
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Four routines
	wg.Add(5)

	//
	// Output tag-cloud, and per-tag pages.
	//
	go func() {
		err := outputTags(entries, recent)
		if err != nil {
			fmt.Printf("Error rendering tag-pages:%s\n", err.Error())
			os.Exit(1)
		}
		wg.Done()
	}()

	//
	// Output the per year/month archive, and the archive-index.
	//
	go func() {
		err := outputArchive(entries, recent)
		if err != nil {
			fmt.Printf("Error rendering archive-pages:%s\n", err.Error())
			os.Exit(1)
		}
		wg.Done()
	}()

	//
	// Output index page.
	//
	go func() {
		err := outputIndex(entries, recent)
		if err != nil {
			fmt.Printf("Error rendering index.html: %s\n", err.Error())
			os.Exit(1)
		}
		wg.Done()
	}()

	//
	// Output RSS feed which has the same information as the index-page.
	//
	go func() {
		err := outputRSS(entries, recent)
		if err != nil {
			fmt.Printf("Error rendering /index.rss: %s\n", err.Error())
			os.Exit(1)
		}
		wg.Done()
	}()

	//
	// Output each entry.
	//
	go func() {
		err := outputEntries(entries, recent)

		if err != nil {
			fmt.Printf("Error rendering blog-posts: %s\n", err.Error())
			os.Exit(1)
		}
		wg.Done()
	}()

	wg.Wait()

	//
	// Report on our runtime
	//
	elapsed := time.Since(start)
	fmt.Printf("Compilation took %s\n", elapsed)

}
