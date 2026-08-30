+++
title = "Building a Redis clone"
postname = "redis_clone"
description = "Learning about Redis by programming a clone of it in Go"
created_at = "30 Aug 2026"
last_updated = "30 Aug 2026"
author = "Poroca"
tags = ["daedalus"]
cover = "/images/redis_go.png"
+++

# Welcome back, me,

I’m getting back to posting on the blog, and I intend to keep creating and developing my skills, especially as a developer. I still work as a QA, but a future career change may be on the horizon.

## The project:

This project is a way to learn a little about Redis—how it works and how to use it—and to get back to building something personal simply for the sake of learning. At the moment, I’m implementing the basic features a database needs to have. You can follow the project on [GitHub](https://github.com/GuPoroca/Redis_clone).

## First steps:

I didn’t know much about Redis beyond the fact that it is used for caching and storing data. After some quick research, I learned about the RESP2 protocol, redis-cli, the basic Redis commands, and how it uses dynamic memory as a database to optimize performance.

## For the clone:

I’ll start with just the basics: SET, GET, DELETE, EXPIRE, and PERSIST commands. For now, it only supports dictionary-style storage with strings as both keys and values, but that will change in the near future.

## Using AI

In this project, I’m trying to partially integrate agents into my development flow, mainly to help me understand the code and guide my next steps while I explore unfamiliar territory. Since I develop on Linux, I’m currently trying out Neovim and [Pardes](https://git.sr.ht/~gbrls/pardes), an editor made by my friend gbrls. Making the workflow work with agents has been a little different from what I’m used to at work.

## That’s all for today, folks!

I plan to keep posting updates as the project takes shape, but that’s it for now. I’d also like to write a post about this site itself, its deployment pipeline, and how I accidentally recreated Hugo—but that is a story for another post.
