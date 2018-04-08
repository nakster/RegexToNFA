## A Program in GO 
![](https://blog.newrelic.com/wp-content/uploads/golang-gopher.jpg)

## Intro
A program in the Go programming language that can build a NFA from a regular expression
This Project is for 3rd Year of software.

## Context
The following project concerns a real-world problem – one that you could
be asked to solve in industry. You are not expected to know how to do it
off the top of your head. Rather, it is expected that you will research and
investigate possible ways to tackle the problem, and then come up with your
own solution based on those. A quick search for solutions online will convince
you that many people have written solutions to the problem already, in many
different programming languages, and many of those are not experienced
software developers. Note that a series of videos will be provided to students
on the course page to help them with the project.

## Problem statement
You must write a program in the Go programming language [2] that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “∗
” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission.

## Setup Instructions

 - First you need to have Go Downloaded on your machine  
 - Download or clone
 - First step is to clone my repository
 - Then go to cmd and change to the folder you cloned
 - run the following commands
 - Go Build "main.go"
 - Then run the Excutable file
 - ./"main.go".exe
 - or
 - Just run as "Go run Main.go"

## Technologies I used
[GoLang](https://golang.org/)
## Links 
- [links](./Research.md)