# contentful-hugo

[![GoDoc](https://godoc.org/github.com/thedodobird2/contentful-hugo?status.svg)](https://godoc.org/github.com/thedodobird2/contentful-hugo)
[![Build Status](https://travis-ci.org/thedodobird2/contentful-hugo.svg?branch=master)](https://travis-ci.org/thedodobird2/contentful-hugo)
[![Go Report Card](https://goreportcard.com/badge/github.com/thedodobird2/contentful-hugo)](https://goreportcard.com/report/github.com/thedodobird2/contentful-hugo)

A tool built with [Go](https://github.com/golang/go) to generate files for [Hugo](https://github.com/gohugoio/hugo)
using content from [Contentful](https://www.contentful.com/)

## About

This application for [Hugo](https://github.com/gohugoio/hugo) allows you to build a static site using the data stored at
[Contentful](https://www.contentful.com/). It is built with [Go](https://github.com/golang/go) and uses the
[Contentful Delivery API](https://www.contentful.com/developers/docs/references/content-delivery-api/).

## Getting started

### Set up

In order to be able to use the contentful-hugo connector you need three things:
* config file
* executable

Both of these should be placed alongside your Hugo website, at the same level as your Hugo root directory for the Connector to work.

.

├── hugo-site

├── config.json

└── contentful-hugo

### Executable

The executable is available in the release package for each [release](https://github.com/thedodobird2/contentful-hugo/releases).
Chose the one suiting your environment.

At the moment these Systems are supported:
* Linux
* OSX
* Windows

If you can't find one that works for you, refer to the Develop step and let Go create the executable for you.

### Config File

Create your config file and name it `config.json`. It should be structure as the following:

```json
{
  "contentful": {
    "spaceId": "",
    "accessToken": ""
  },
  "hugo": {
    "root": ""
  }
}
```

To get the Space ID and Access Token, an API Key for your Contentful Space needs to be created which will give you those
two properties.

The Hugo root should be the name of your Hugo root directory.

## Use
The Connector is a command line tool. It has one flag (`-c` or `-contentType`) with which you can tell the Connector the
Content Type in Contentful you want all the Entries from to be added to your Hugo site.

```$bash
contentful-connector -c contentType
```

This command will get all the Entries from Contentful with the Content Type ID _contentType_ and put them as separate
markdown files in the section with the name of their Content Type. All markdown file names will be Entry IDs.

## Develop & Contribute

If you want to contribute in development, simply:
```$bash
go get github.com/thedodobird2/contentful-connector
```

Otherwise if you have problems or suggestions for improvement, just open an issue or a pull request.