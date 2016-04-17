### Write an Image Cropping program with Go

Here you will find the main code, ``crop.go`` for the article in Linux Voice, Issue 26. 

#### Go version

At the time of writing the article, the latest version of Golang was 1.5. However, by the time of publication, Go 1.6 was out. The programs should all work without any issues with this version. 

As recommended in the article, please refer to the [Go install guide](https://golang.org/doc/install) for installing the latest Go tools for your operating system. Your distribution's repository may not have the latest version of the tools.

#### Playing with Go

In the article, we setup a proper development environment. However,when you are experimenting or want to share a Go code snippet with someone, the [Go playground](https://play.golang.org) may prove to be useful.

#### Go editor

If you are looking for the perfect editor for Go programs, please refer to the guide [here](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) for all your options.

#### Formatting your Go code

The [gofmt](https://blog.golang.org/go-fmt-your-code) tool which should be installed for you takes care of formatting your Go code appropriately.

### Documentation for Go packages

[GoDoc](https://godoc.org/) hosts the documentation for Go packages. The interesting thing about GoDoc is that you don't have to do anything to add your package for it. Simply, go to the URL https://godoc.org/<path/to/your/package> and you will able to see the docs (One caveat is that your package must be hosted on GitHub, BitBucket, Launchpad or Google code). 

For example, documentation of the cutter package can be viewed at [https://godoc.org/github.com/oliamb/cutter](https://godoc.org/github.com/oliamb/cutter). You can find the documentation for all the standard library packages at [https://golang.org/pkg/](https://golang.org/pkg/).

It's important though that we follow the [guidelines](http://blog.golang.org/godoc-documenting-go-code) while writing our code so that GoDoc renders it properly and the important parts of our package is rendered correctly.


### Resources

While setting up the development environment, I referred to the term "workspace". The documentation on [Go workspaces](https://golang.org/doc/code.html#Workspaces) should be a good next document to learn more about it.

The following resource may help you continue your learning:

- [Understanding packages in Go](http://thenewstack.io/understanding-golang-packages)
- [An introduction to programming in Go](https://www.golang-book.com/books/intro)
- [Go by example](https://www.gobyexample.com)
- [An incomplete list of Go tools](http://dominik.honnef.co/posts/2014/12/an_incomplete_list_of_go_tools/)

### Miscellaneous

- [Why does the golang flag package return pointers to primitives like bool instead of values?](https://plus.google.com/u/1/104364549279117274863/posts/EerjHxeeiDE)
