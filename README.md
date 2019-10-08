# Ephemeris

For many years I've generated [my blog](https://blog.steve.fi/) with the [chronicle blog compiler](https://steve.fi/Software/chronicle/), but now that project is showing its age.

The chronicle blog compiler started life as a simple project, but grew in complexity over time.  Part of the reason the complexity grew was because the project was very flexible:

* It would read a whole bunch flat files, each of which contained a single blog-post.
* Each parsed post would then be inserted into an SQLite database.
* Using this intermediary database a series of plugins would each execute in turn:
  * The plugins would run SQL-queries to extract posts of interest.
    * For example building a tag-cloud.
    * For example building an archive-view.
    * For example outputting the front-page (10 most recent posts) & associated RSS-feed.
* Once complete the SQLite database would be destroyed.

The intention was that the use of the intermediary SQLite database would allow everything to be generated in a flexible fashion, with very loose coupling.  However over time most of the complexity I had previous required fell away, as I completed the process of tidying up entries - which had been imported from previous blogging-solutions.

In short this project was started to do only what I actually require:

* Read the blog-posts which chronicle previously read.
* Generate all the pages.

There is no plugin-mechanism, there is no theming-support, instead there is a single theme which matches the previous default theme I used/wrote.


## Blog Format

As with `chronicle` the input to this program is a directory containing a series of blog-posts.  Each post will be stored in a single file, with the entry being prefixed by a header containing meta-data.

A sample post would look like this:

```
Subject: This is my post
Date: DD/MM/YYYY HH:MM
Format: markdown
Tags: foo, bar baz

This is my post
```

There are a few things to note here:

* The date **MUST** be in the specified format.
* If there is no `format: markdown` header then the body will be assumed to be HTML.
  * All my early posts were written in HTML.
  * Later I switched to markdown.


## Implementation

Most of the code, in the core package, is associated with parsing a single
blog-post from a text-file.

Once we have a single blog post then a site is just an array of them!

The only real complication is walking that array of posts to output
each page.  We do that in the `cmd` subdirectory.


## TODO

* Style pages
