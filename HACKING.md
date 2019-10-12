# Hacking

This document briefly discusses how you'd make changes to the output pages.  In short there are is no support for themes, and all generated pages are made via a series of embedded template-files.

If you wish to change the look/feel of the generated file you'll need to look in the `data/` directory beneath `cmd/ephemeris`.


## Making Changes

To make a change requires three steps:

* Make a change to the template/templates you wish.
* Regenerate the compiled version of the templates.
  * This uses the [implant](https://github.com/skx/implant) utility.
* Rebuild the application with the compiled version.

To get started visit the main application directory:

    cd cmd/ephemeris

Make your change(s) to the files beneath `data/`:

    vi data/index.tmpl

These files are not read directly, instead their contents are compiled into the file `static.go`.  You need to regenerate that file to see your changes reflected, once you rebuild the application.

So you should run:

    implant
    go build .
    [go install .]

That will ensure that the file `static.go` contains updated versions of the templates from beneath `data/` and the compiler has been rebuilt with an updated copy of them.
