
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
![License](https://img.shields.io/badge/License-AGPL-blue.svg)

# caa

R Wrapper for the Go dnscaa Library

## Description

Experimental R wrapper for the Go dnscall library.

## What’s Inside The Tin

The following functions are implemented:

  - `caa_dig`: Retrieve the first CAA issuer value for a domain (if any)

## Installation

``` r
remotes::install_gitlab("hrbrmstr/caa")
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
caa_dig("google.com")
## [1] "pki.goog"
```

## caa Metrics

| Lang               | \# Files |  (%) |    LoC |  (%) | Blank lines |  (%) | \# Lines |  (%) |
| :----------------- | -------: | ---: | -----: | ---: | ----------: | ---: | -------: | ---: |
| Go                 |     1293 | 0.92 | 374699 | 0.95 |       40503 | 0.95 |    39876 | 0.95 |
| Assembly           |       69 | 0.05 |  10426 | 0.03 |        1258 | 0.03 |     1424 | 0.03 |
| XML                |        1 | 0.00 |   4780 | 0.01 |         228 | 0.01 |       13 | 0.00 |
| HTML               |       10 | 0.01 |   2214 | 0.01 |         441 | 0.01 |       16 | 0.00 |
| Bourne Shell       |        5 | 0.00 |    798 | 0.00 |         113 | 0.00 |      429 | 0.01 |
| C                  |        6 | 0.00 |    332 | 0.00 |         100 | 0.00 |       77 | 0.00 |
| Dockerfile         |        3 | 0.00 |     91 | 0.00 |          27 | 0.00 |       23 | 0.00 |
| Bourne Again Shell |        2 | 0.00 |     72 | 0.00 |          12 | 0.00 |        6 | 0.00 |
| YAML               |        2 | 0.00 |     45 | 0.00 |           0 | 0.00 |        0 | 0.00 |
| C/C++ Header       |        1 | 0.00 |     42 | 0.00 |          24 | 0.00 |       10 | 0.00 |
| make               |        2 | 0.00 |     25 | 0.00 |           7 | 0.00 |        4 | 0.00 |
| Rmd                |        1 | 0.00 |      9 | 0.00 |          16 | 0.00 |       30 | 0.00 |
| R                  |        3 | 0.00 |      7 | 0.00 |           2 | 0.00 |       15 | 0.00 |

## Code of Conduct

Please note that this project is released with a Contributor Code of
Conduct. By participating in this project you agree to abide by its
terms.
