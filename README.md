
# go-server

A toy project for a go server: used to experiment with some construct to prepare for a more complex application.

# Installation

- Clone the repo: `git clone git@github.com:Knoblauchpilze/go-server.git`.
- Install Go from [here](https://go.dev/doc/install). **NOTE**: this project expects Go 1.20 to be available on the system.
- Go to the project's directory `cd ~/path/to/the/repo`.
- Compile and install: `make`.
- Execute any application with `make run app_name`.

# Learnings

We used the [chi](https://github.com/go-chi/chi) router to handle the routing. This replaces [gorilla mux](https://github.com/gorilla/mux) as this one is now archived.

We got a nice code coverage badge:

[![codecov](https://codecov.io/gh/codecov/go-server/branch/master/graph/badge.svg?token=tNKcOjlxLo)](https://codecov.io/gh/codecov/go-server)