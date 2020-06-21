[![GoDoc](https://godoc.org/github.com/skx/ephemeris?status.svg)](http://godoc.org/github.com/skx/ephemeris)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/ephemeris)](https://goreportcard.com/report/github.com/skx/ephemeris)
[![license](https://img.shields.io/github/license/skx/ephemeris.svg)](https://github.com/skx/ephemeris/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/ephemeris.svg)](https://github.com/skx/ephemeris/releases/latest)


* [Ephemeris](#ephemeris)
  * [Chronicle History](#chronicle-history)
  * [Chronicle Migration](#chronicle-migration)
* [Installation](#installation)
* [Blog Generation](#blog-generation)
   * [Blog Format](#blog-format)
* [Demo Blog](#demo-blog)
* [Hacking](#hacking)
* [Feedback](#feedback)


# Ephemeris

Ephemeris is a golang application which will generate a blog from a collection of static text-files, complete with:

* Archive-view.
* Tag-cloud.
* RSS feed

The project was primarily written to generate [my own blog](https://blog.steve.fi/), which was previously generated with the perl-based [chronicle blog compiler](https://steve.fi/Software/chronicle/).


## Chronicle History

The chronicle blog compiler started life as a simple project, but grew in complexity over time.  Part of the reason the complexity grew was because the project was very flexible:

* It would read a whole bunch flat files, each of which contained a single blog-post.
* Each parsed post would then be inserted into an SQLite database.
* Using this intermediary database a series of plugins would each execute in turn:
  * The plugins would run SQL-queries to extract posts of interest.
    * For example building a tag-cloud.
    * For example building an archive-view.
    * For example outputting the front-page (10 most recent posts) & associated RSS-feed.
* Once complete the SQLite database would be destroyed.

My expectation was that the use of an intermediary SQLite database would allow content to be generated in a very flexible and extensible fashion, however over time it became apparant that I didn't need things to be too flexible!

In short this project was born to __replace__ chronicle, and perform the things I actually need, rather than what I _suspected_ I might want.



## Chronicle Migration

There are some [brief notes on migration](MIGRATION.md) available.




# Installation

You can install from source, by cloning the repository and running:

    cd ephemeris/cmd/ephemeris
    go build .
    go install .

Or you can visit the [release page](https://github.com/skx/evalfilter/releases) to download a binary.




# Blog Generation

The application has only a couple of configuration values, which must be setup
in the `ephemeris.json` file:

        {
          "Posts":    "posts/",
          "Comments": "/comments/",
          "Prefix":   "http://blog.steve.fi/"
        }

Create a suitable file, then run the application with no arguments:

    $ ephemeris

The generated output will be placed in the `output/` directory.  The configuration-keys in the JSON file are:

* `Posts`
  * This is the path to the directory containing your blog-posts.
  * The input directory will be searched recursively.
* `Comments`
  * This is the path to the directory containing your comments.
  * If this is empty then no comments will be read/inserted into your output
* `Prefix`
  * This is the URL-prefix used to generate all links.

There is a command-line flag which lets you specify an alternative configuration-file, if you do not wish to use the default.  Run `ephemeris -help` to see details.



## Blog Format

As with `chronicle` the input to this program is a directory tree containing a series of blog-posts.  Each post will be stored in a single file, with the entry being prefixed by a header containing meta-data.

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

As noted the input directory will be processed recursively, which allows you to group posts by topic, year, or in any other way you might prefer.  I personally file my entries by year:

```
data/
  ├── 2005
  ├── 2006
  ..
  ├── 2018
  ├── 2019
  └── 2020
```


# Demo Blog

There is a demo-blog contained within this repository; to use it:

```
cd _demo
ephemeris
```

This will generate `_demo/output/`, and you can then serve that via:

```
cd output/
python -m SimpleHTTPServer 8000
```

Finally open http://localhost:8000 in your browser




# Hacking

Some brief notes on the theme/output generation are available in [HACKING.md](HACKING.md).


# Feedback

This project is not documented to my usual and preferred standard, no doubt it will improve over time.

However there are many blog packages out there, so I expect this project will only be of interest to those wishing to switch from `chronicle`.

Steve
