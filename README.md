# BotsArchive API Wrapper for Go

<a href="https://sonarcloud.io/dashboard?id=MassiveBox_botsarchive_api"><img src="https://sonarcloud.io/api/project_badges/measure?project=MassiveBox_botsarchive_api&metric=alert_status"></a> <a href="https://travis-ci.com/github/MassiveBox/botsarchive_api/builds/"><img src="https://travis-ci.com/MassiveBox/botsarchive_api.svg?branch=main"></a> <a href="https://goreportcard.com/report/github.com/MassiveBox/botsarchive_api"><img class="badge" tag="github.com/MassiveBox/botsarchive_api" src="https://goreportcard.com/badge/github.com/MassiveBox/botsarchive_api"></a> <a href="https://t.me/MassiveBox"><img src="https://img.shields.io/badge/contact-@MassiveBox-blue?logo=telegram"></a> <a href="https://massivebox.eu.org/?page=4"><img src="https://img.shields.io/badge/support-the%20project-yellow?logo=symantec"></a> 

**Reference:** <a href="https://pkg.go.dev/github.com/MassiveBox/botsarchive_api"><img src="https://img.shields.io/static/v1?label=read%20the&message=documentation&color=blue&logo=go"></a> <a href="https://github.com/MassiveBox/botsarchive_api/wiki"><img src="https://img.shields.io/static/v1?label=check%20out&message=the%20wiki&color=white&logo=Read%20the%20Docs"></a>

BotsArchive API Wrapper for Go aims to provide a simple interface to access botsarchive.com's API protocol with Go in a easy and fast manner.  

### Why?

Using the original BotsArchive's API directly with Go is a pain in the butt: numbers as strings, values that sometimes are there and sometimes aren't, bools as integers, parameters that disappear...

With this package, my goal is to simplify everything so you don't have to deal with all this bullcrap.  

To achieve my goal, I've decided to completely change the API return values instead of keeping the original ones. Read the [docs](https://pkg.go.dev/github.com/MassiveBox/botsarchive_api) for this package to know where to find everything.

### How to use

Install the package:

```
go get github.com/MassiveBox/botsarchive_api
```

and refer to the [docs](https://pkg.go.dev/github.com/MassiveBox/botsarchive_api) or the [wiki](https://github.com/MassiveBox/botsarchive_api/wiki)  for a detailed explanation. You can find some examples in the wiki as well.

### Technical details

To improve performances and memory usage, instead of net/http, I'm using [FastHTTP](https://github.com/valyala/fasthttp) and instead of encoding/json, I'm using [JSON Iterator.](https://github.com/json-iterator/go)

This package will send data to the host `api.botsarchive.com` - Please make sure you aren't blocking it with a firewall.

### Coverage

This package is supposed to be 100% compatible with BotsArchive's [*public* API](https://botsarchive.com/docs.html). If you find something isn't working as intended, open an issue.

I also have plans for the future to support BotsArchive's *private* API, the one that is used in the website to gather information about the bots, top bots, categories, etc.

### Legal

All the data you get with this package comes from BotsArchive's API. It's exclusively BotsArchive's responsibility if you find something that violates the law.