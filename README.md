# Ephemeris

This is a simple blogging engine, for https://blog.steve.fi/


## Blog Format

Blog format is the obvious one, a header prior to the content with date/subject/tags.

A sample post would look like this:

```
Subject: This is my post
Date: DD/MM/YYYY HH:MM
Format: markdown
Tags: foo, bar baz

This is my post
```

There are only two things to note:

* Format is assumed to be HTML.  Unless you specify `format: markdown`.
* The date **MUST** be in the specified format.


## Implementation

Most of the code, in the core package, is associated with parsing a single
blog-post from a text-file.

Once we have a single blog post then a site is just an array of them!

The only real complication is walking that array of posts to output
each page.  We do that in the `cmd` subdirectory.


## TODO

* Style pages
