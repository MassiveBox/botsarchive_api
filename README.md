# BotsArchive API

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=MassiveBox_botsarchive_api&metric=alert_status)](https://sonarcloud.io/dashboard?id=MassiveBox_botsarchive_api) [![TravisCI](https://travis-ci.com/MassiveBox/botsarchive_api.svg?branch=master)](https://travis-ci.com/github/MassiveBox/botsarchive_api/builds/) [<img class="badge" tag="github.com/MassiveBox/botsarchive_api" src="https://goreportcard.com/badge/github.com/MassiveBox/botsarchive_api">](https://goreportcard.com/report/github.com/MassiveBox/botsarchive_api) [![DM me badge](https://img.shields.io/badge/contact-@MassiveBox-blue?logo=telegram)](https://t.me/MassiveBox) [![Donate](https://img.shields.io/badge/support-the%20project-yellow?logo=symantec)](https://massivebox.eu.org/?page=4)

BotsArchive API aims to provide a simple interface to access botsarchive.com's API protocol with Go in a easy and fast manner.  

### Why?

BotsArchive's API is known to be a pain to work with in Go: numbers as strings, values that sometimes are there and sometimes aren't, bools as integers...  

With this package, my goal is to simplify everything so you don't have to deal with all this bullcrap.  

To achieve my goal, I've decided to completely change the API return values instead of keeping the original ones. Read the docs for this package to know where to find everything.

### How to use

Install the package:

```
go get github.com/MassiveBox/botsarchive_api
```

and refer to the docs for a detailed explanation. You can find some examples in the wiki as well.

### Technical details

To improve performances and memory usage, instead of net/http, I'm using [FastHTTP](https://github.com/valyala/fasthttp) and instead of encoding/json, I'm using [JSON Iterator.](https://github.com/json-iterator/go)

This package will send data to the host `api.botsarchive.com` - Please make sure you aren't blocking it with a firewall.

### Coverage

This package is supposed to be 100% compatible with BotsArchive's *public* API. If you find something isn't working as intended, open an issue.

I also have plans for the future to support BotsArchive's *private* API, the one that is used in the website to gather information about the bots, top bots, categories, etc.

### Legal

All the data you get with this package comes from BotsArchive's API. It's exclusively BotsArchive's responsibility if you find something that violates the law.
