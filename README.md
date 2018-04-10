# Contentful-Hugo Connector

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

### Prerequisites

* Hugo up and running
* access to a Contentful Space

_Note: The Connector and Hugo are built with Go, but for the normal usage of both of these you don't need to install Go._

### Step 1: Download

Two things are needed:
  * executable
  * config file

Download them both here: [releases](https://github.com/thedodobird2/contentful-hugo/releases).

#### executable
Chose the the executable suiting your environment. If you can't find an executable that works for you, you can let Go build
the executable from the source locally (see the _Develop & Contribute_ step).

#### config file
The config file needs to have the name `config.json` and should look something like this:

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

Both the config file and executable need to be placed alongside your Hugo website, at the same level as your Hugo root directory.

```
.
├── hugo-site
├── config.json
└── contentful-hugo
```

### Step 2: Configure

#### contentful
To get the `spaceId` and Contentful Delivery API `accessToken` for your config, an API Key for your Contentful Space needs to be created.
This can be found in your [app](https://app.contentful.com/) in `Space settings -> API Keys`.

#### hugo
The `root` property is the name of your Hugo root directory.

Add all of these properties to your `config.json`.

## Use

At the moment you can get _all_ the Entries from _one_ existing Content Type at a time.

Specify with the flag `-c` or `-contentType` the ID of the Content Type you would like to add to Hugo.

```$bash
contentful-hugo -c exampleContentType
```

#### what happens
The command will get all the _published_ Entries from Contentful with the Content Type ID `exampleContentType`. A Markdown
file is created for each Entry. In each file all the Contentful data of the given Entry is written into the Front Matter
(in JSON). All of the generated Markdowns are put in [sections](https://gohugo.io/content-management/sections/) according
to the Entry's Content Type.

#### now what?
For the [content](https://gohugo.io/content-management/) that you now have in Hugo you can create [templates](https://gohugo.io/templates/).
Accessing the properties out of Front Matter for your template can be done like so:

```gotemplate
{{ .Param "contentful.fields.title"}}
```

For further help with Hugo, refer to their [documentation](https://gohugo.io/documentation/).

## Develop & Contribute

If you want to contribute in development, simply:
```$bash
go get github.com/thedodobird2/contentful-hugo
```

#### otherwise
If you have problems or suggestions for improvement, just open an issue or a pull request.