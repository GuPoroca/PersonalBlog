+++
title = "Hextest REST API testing"
postname = "hextestv01"
description = "My main project currently, testing REST APIs in an easy and simple way"
created_at = "18 Sep 2025"
last_updated = "18 Sep 2025"
author = "Poroca"
tags = ["daedalus"]
cover = "/images/hextest.jpg"
+++

# About testing REST APIs

Right when I started at my current job, I felt like I had been hired to solve the project’s test automation problems. I didn’t have much experience with the testing framework they used (Mocha Chai), and it took me a while to get things working again little by little.

### I was the only one touching it

Apparently, before I joined, another QA Engineer had worked on the automated tests, but from the day she was fired until the day I left that project, I was the only one maintaining and interacting with that repository. This meant that when I was reassigned to another project, yet again someone else had to go through the same painful process of trying to understand what was going on there. And this wasn’t limited by the person’s seniority, but rather by the tests themselves.

## I wanted to get back to coding

In August of this year, I realized that I had become too comfortable, and I wanted to create a project to get back into practicing programming. My developer side had been dormant, and whenever I needed some code, I first tried to get it through AI.

<figure>
    <img src="/images/gopher.jpg" alt="photo">
    <figcaption>
        <p>Gopher</p>
    </figcaption>
</figure>

I chose Go as the language because of its charm. I’ve always liked more low-level languages, but I didn’t want to pick something like C or Zig because of the kind of application I wanted to build, and Rust honestly scares me. Go is fast and simple, and it makes me focus on solving problems instead of on the language itself. I’m somewhat resentful of today’s entire web development culture and its frameworks, and I’m not a fan of JavaScript either. So, Go it is.

## The tool

The plan was clear: a tool that would read a JSON containing the project data and its tests, build a struct inside the tool, make HTTP requests, and return the results. The result is pretty close to that, but how we got from (Project) -> (Suite) -> (Test) -> (Assert) -> (Check) is something only the late-night coding sessions can explain...

### Current state

I consider the tool to be in a pre-alpha -1.0 version (yes, negative versioning, where is your god now?). But we already have a front end to edit the tests without having to manually touch the JSON for those not interested, and most importantly, the tests actually run. Still, there’s a lot to improve: the front end is horrendous and could have infinitely better UX, the back end could generate better reports instead of just printing results, and many other things—but unfortunately, it’s just me.  
I’ll keep updating the development process of this tool here!

<figure>
    <img src="/images/hextest.jpg" alt="photo">
</figure>
