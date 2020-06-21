# Migration from chronicle

As with `chronicle` the input to this program is a directory containing a series of blog-posts, along with a path from which to load static-comments.

There are two changes in this project which you will have to adjust to:

* The format of blog-posts became more strict.
* The naming of the comment-files bacme more strict.


## Blog Posts

Each post will be stored in a single file, with the entry being prefixed by a header containing meta-data. A sample post would look like this:

```
Tags: markdownshare, puppet, marionette, github, oodi, university
Date: 07/04/2020 09:00
Subject: A busy few days
Format: markdown

Over the past few weeks things have been pretty hectic.
```

There are a few things to note here:

* The date **MUST** be in the specified format.
* If there is no `format: markdown` header then the body will be assumed to be HTML.
  * All my early posts were written in HTML.
  * Later I switched to markdown.

A related change is that it is now a __fatal-error__ for a blog-post to have a header-key which is unknown.  To provide a concrete example it was previously possible to write:

```
Subject: I won't write another email client
Tags: golang, email, maildir, maildir-utils
Date: 08/01/2020 19:19
Blah: foo
Publish: later
Format: markdown

Once upon a time I wrote an email client, in a combination of C++ and Lua.
```

Now `Blah`, and `Publish` are explicitly prohibited.



## Comment Files

In the past comments would be written to files with names such as:

* `you_ve_had_this_coming_since_the_day_you_arrived.html.23-November-2008-13:18:09`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.23-November-2008-13:20:39`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.23-November-2008-14:20:40`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.23-November-2008-14:44:15`

We now expect these to be named:

* `${link}.${ctime}`

So these examples would become:

* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227432366`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227439089`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227439239`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227442840`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227444255`

See [COMMENTS.md](COMMENTS.md) for a discussion on comment-setup.




## Migration Summary

To migration first of all try pointing the configuration file at your existing entries, ignoring comments.  You'll need to fix the `Date:` header of your posts  until the errors go away.

Once you've done that you can now rename the comments, if you use them.
