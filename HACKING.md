# Hacking

This document briefly discusses how you'd make changes to the output pages.  In short there are is no support for themes, and all generated pages are made via a series of embedded template-files.

If you wish to change the look/feel of the generated file you'll need to look in the `data/` directory beneath `cmd/ephemeris`.


## Making Changes

Making a change to the templates requires two steps:

* Make a change to the template/templates you wish to modify.
* Rebuild the application to make those changes available at runtime.

To get started visit the main application directory:

    cd cmd/ephemeris

Make your change(s) to the files beneath `data/`:

    vi data/index.tmpl

Now rebuild the application:

    go build .

Finally you can regenerate your blog with the updated templates:

    ephemeris [args]
