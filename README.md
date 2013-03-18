nginx2stathat
===============

Program to watch a nginx combined access.log and send HTTP stats to StatHat

Installation
============

 * Install Go: http://golang.org/doc/install#install
 * Setup your GO environment variables (GOPATH, GOBIN, PATH)

    $ sudo go install github.com/hgfischer/nginx2stathat 

Usage
=====

To see the command line options of nginx2stathat, just run it without parameters (or with -h):

    $ nginx2stathat -h

Currently these are the options:

    Usage: ./nginx2stathat [flags] [EZ Key] [access.log]
    Flags:
      -parsers=4: Number of parallel routines parsing log lines and queueing them to the posters
      -posters=4: Number of parallel routines sending results to StatHat
      -prefix="": Stat prefix. Ex.: "`hostname -s` live site"