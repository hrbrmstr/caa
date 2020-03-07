
[![Project Status: Active – The project has reached a stable, usable
state and is being actively
developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![Signed
by](https://img.shields.io/badge/Keybase-Verified-brightgreen.svg)](https://keybase.io/hrbrmstr)
![Signed commit
%](https://img.shields.io/badge/Signed_Commits-100%25-lightgrey.svg)
[![Linux build
Status](https://travis-ci.org/hrbrmstr/caa.svg?branch=master)](https://travis-ci.org/hrbrmstr/caa)  
![Minimal R
Version](https://img.shields.io/badge/R%3E%3D-3.2.0-blue.svg)
![License](https://img.shields.io/badge/License-MIT-blue.svg)

# caa

R Wrapper for the Go dnscaa Library

## Description

Experimental R wrapper for the Go dnscall library. Builds off of work
done by Romain Francois (<https://github.com/rstats-go>).

## What’s Inside The Tin

The following functions are implemented:

  - `caa_dig`: Retrieve the CAA record values for a domain (if any)

## WAT?\!

Grabbing & parsing DNS CAA records isn’t the actual purpose of this
pkg/repo. Romain did a fantastic job cracking the Go nut for R, but not
many R packages seem to be based on using Go (C, Rcpp, Fortran, and even
Rust seems to be more en vogue) and the existing examples work with
basic types.

This package uses Go libraries and makes a data frame from the results
of Go-side network queries for DNS CAA records. It relies 100% on Go for
networking and DNS record parsing, which only leaves getting the results
back into R.

`main.go` contains some useful idioms, such as the comment block at the
top right before the first `import` which makes “unsafe” C things there
available to Go (so we get `SEXP`\!) Midway down Go String slices are
turned into C-compatible character arrays, along with the code to free
up those bits of heap memory at the end (`defer`s). We use the external
reference to `MakeDF()` (`C.MakeDF()`) in `main.c` to make the `SEXP`
data frame we’re returning. We could have tried to keep that in Go but
it’s cleaner (IMO) to just let C do the work. NOTE that it’s important
to use `free()` on each element of the 2D (`char **`) character array on
the C side since they were allocated on the heap on the Go side and Go’s
memory manager won’t auto-free those for you.

`main.go` also has:

``` go
func init() {
  log.SetOutput(ioutil.Discard)
}
```

in it to stop modules that use the Go logger from outputting things (Go
programmers tend to be “systems programmers” and, as a result, log the
heck out of everything).

Also on the C side is the `R_caa_dig()` function which we `.Call()` from
the R side (along with the package registration code for it). Given that
I had to stick a `.c` file in `src/` I’ll eventually (likely) modify the
`Makefile` and put the required R registration code there.

The `.Rbuildignore` removes some Go module repository cruft.

Ultimately, it still does not pass CRAN checks. The lib dir is YUGE,
`cgo`’s preferred/required compilation settings are against the CRAN
rules, and the requirement for GNU Make is also somewhat verbotten.
However, a package check with CRAN checks enabled only results in 1
warning and 2 notes (on macOS).

For this particular package, the `dnscaa` Go library had a call to
`fmt.Println()` in `dnscaa.go` when there was an NXDOMAIN result which I
had to turn into `log.Println()`. Unfortunately, the way `go get` works
with git repos, they’re submodules, so those changes will get
overwritten and you’ll have to add them manuall (keep them as git
submodules so you can more easily update the underlying libraries if
need be).

As a result of ^^ you’ll need to do this to clone this repo effectively:

    $ git clone --recurse-submodules -j8 ssh://git@git.rud.is:7777/hrbrmstr/caa.git

Replace `git clone --recurse-submodules -j8
ssh://git@git.rud.is:7777/hrbrmstr/caa.git` with the `https` or `ssh`
URL of where you’re viewing this and have access.

There are comments peppered throughout `main.go` and especially
`main.c`, but if anything needs clarification, drop an issue.

This is my first time through the `cgo` interface (Go itself is
straightforward, `cgo` has lots of gotchas), so if I missed anything
there def lemme know.

## Installation

``` r
remotes::install_git("https://git.rud.is/hrbrmstr/caa.git")
# or
remotes::install_gitlab("hrbrmstr/caa")
# or
remotes::install_github("hrbrmstr/caa")
```

NOTE: To use the ‘remotes’ install options you will need to have the
[{remotes} package](https://github.com/r-lib/remotes) installed.

## Usage

``` r
library(caa)

# current version
packageVersion("caa")
## [1] '0.1.0'
```

``` r
# one record
caa_dig("google.com")
## # A tibble: 1 x 2
##   tag   value   
##   <chr> <chr>   
## 1 issue pki.goog

# multiple
caa_dig("www.comodo.com")
## # A tibble: 3 x 2
##   tag   value                       
##   <chr> <chr>                       
## 1 iodef mailto:sslabuse@comodoca.com
## 2 issue comodoca.com                
## 3 issue digicert.com

# none (lookup error)
caa_dig("www.comodo.comm")
## # A tibble: 0 x 2
## # … with 2 variables: tag <chr>, value <chr>
```

## caa Metrics

| Lang               | \# Files |  (%) |    LoC |  (%) | Blank lines |  (%) | \# Lines |  (%) |
| :----------------- | -------: | ---: | -----: | ---: | ----------: | ---: | -------: | ---: |
| Go                 |     1293 | 0.92 | 376551 | 0.95 |       40523 | 0.95 |    38063 | 0.95 |
| Assembly           |       69 | 0.05 |  10426 | 0.03 |        1258 | 0.03 |     1424 | 0.04 |
| XML                |        1 | 0.00 |   4780 | 0.01 |         228 | 0.01 |       13 | 0.00 |
| HTML               |       10 | 0.01 |   2214 | 0.01 |         441 | 0.01 |       16 | 0.00 |
| Bourne Shell       |        5 | 0.00 |    798 | 0.00 |         113 | 0.00 |      429 | 0.01 |
| C                  |        6 | 0.00 |    359 | 0.00 |         117 | 0.00 |       83 | 0.00 |
| Dockerfile         |        3 | 0.00 |     91 | 0.00 |          27 | 0.00 |       23 | 0.00 |
| Bourne Again Shell |        2 | 0.00 |     72 | 0.00 |          12 | 0.00 |        6 | 0.00 |
| C/C++ Header       |        1 | 0.00 |     48 | 0.00 |          28 | 0.00 |       10 | 0.00 |
| YAML               |        2 | 0.00 |     45 | 0.00 |           0 | 0.00 |        0 | 0.00 |
| make               |        2 | 0.00 |     25 | 0.00 |           7 | 0.00 |        4 | 0.00 |
| R                  |        3 | 0.00 |     15 | 0.00 |           8 | 0.00 |       25 | 0.00 |
| Rmd                |        1 | 0.00 |     14 | 0.00 |          34 | 0.00 |       52 | 0.00 |

## Code of Conduct

Please note that this project is released with a Contributor Code of
Conduct. By participating in this project you agree to abide by its
terms.
