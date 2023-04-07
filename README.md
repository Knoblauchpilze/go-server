
# go-server

A toy project for a go server: used to experiment with some construct to prepare for a more complex application.

# Installation

- Clone the repo: `git clone git@github.com:Knoblauchpilze/go-server.git`.
- Install Go from [here](https://go.dev/doc/install). **NOTE**: this project expects Go 1.20 to be available on the system.
- Go to the project's directory `cd ~/path/to/the/repo`.
- Compile and install: `make`.
- Execute any application with `make run app_name`.

# Learnings

## General idea

The goal was to build a small client-server application where we could practice REST design, logging, authentication and more. The result is a very simple server where a user can sign-up, then log in and finally access to the list of users registered in the server or details about one user. These features are only available if the user is logged in.

*Warning:* We didn't implement any persistent mechanism (DB or otherwise) so for now every information is only persisted in memory.

## Router

Sadly, [gorilla mux](https://github.com/gorilla/mux) was archived due to a lack of maintainers. The Internet says that it does not mean that it is now completely useless, but out of fear we decided to go for another router: [chi](https://github.com/go-chi/chi).

This router seems similar to the gorilla one. We played around with it and created a few middlewares and custom endpoints.

### Context

Generally it looks like this (see [request](pkg/middlewares/request.go)):
```go
func RequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd := NewRequestData()

		ctx := context.WithValue(r.Context(), requestDataKey, rd)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
```

We create some value, and pass it on to the context before calling the next handler. In the next handler we can then use the value registered in the context with the key. In order to be self-contained, the package also defines the following method:
```go
func GetRequestDataFromContextOrFail(w http.ResponseWriter, r *http.Request) (RequestData, bool) {
	ctx := r.Context()
	reqData, ok := ctx.Value(requestDataKey).(RequestData)
	if !ok {
		http.Error(w, "huho", http.StatusInternalServerError)
	}

	return reqData, ok
}
```

This makes very easy in following handlers to try to access to the request like so:
```go
func customHandler(w http.ResponseWriter, r *http.Request) {
	reqData, ok := middlewares.GetRequestDataFromContextOrFail(w, r)
	if !ok {
		return
	}

	// Do some processing
}
```

### How to use a context

After a context is defined, one can use it in routers like so:
```go
func SomeRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.RequestCtx)
		r.Post("/", customHandler)
	})

	return r
}
```

An example can be found in the login handler [here](internal/routes/login.go).

### Handlers with params

Another interesting point concerns creating handlers which need some arguments: the signature of the http handler is fixed, but we can use a construct like so:
```go
func generateHandlerWithParams(foo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We can use foo here.
		w.Write(foo)
	}
}
```

Putting it together we end up with a code like this:
```go
func SomeRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.RequestCtx)
		r.Post("/", generateHandlerWithParams("haha"))
	})

	return r
}
```

This is very similar to what is described in the context [section](#how-to-use-a-context), except that there's an indirection to generate the handler.

### Context with params

The same is true for contexts: it is a bit trickier as contexts for `chi` are supposed to look like this:
```go
func(http.Handler) http.Handler
```

So a func which takes a handler and generates a new handler. In order to have a parameterized context, we need to use a syntax similar to the following:
```go
func generateCustomContext(foo string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "key", foo)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
```

We can use it as middleware like so:
```go
r := chi.NewRouter()

r.Route("/", func(r chi.Router) {
	r.Use(generateCustomContext("haha"))
		r.Get("/", customHandler)
	})
```

## Logger

We decided to go for [logrus](https://github.com/sirupsen/logrus). Not much was experimented with it, we just know that it has built-in support for log levels, formatting and seems to in general be highly configurable. For example it seems possible to prefix some elements (useful for request tracing) and also to output data with json format, to a file, with structured fields...

This seems much more interesting than building our own logger system.

## Errors

Errors are usually a tricky topic. In the past we used extensively `fmt.Errorf` and provided basic wrapping with:
```go
if err := opProducingErrors(); err != nil {
	return fmt.Errorf("failed (cause: %v)", err)
}
```

This was nice but made chaining of errors pretty hard as we can't easily access to the cause of an error.

For this project we used a different approach summarized by the [errorImpl](pkg/errors/error.go) struct:
```go
type ErrorWithCode interface {
	Code() ErrorCode
}

type errorImpl struct {
	hasCode bool
	Value   ErrorCode `json:"Code"`
	Message string
	Cause   error `json:",omitempty"`
}
```

So basically we define an interface for an error with code which some of the errors we return in the apps. If not then we have a `MarshalJSON` method for it which marshal the error in a structured way.

There are also convenience defines which allow to build an error from a known code like so:
```go
if cause != nil {
	return errors.WrapCode(cause, errors.ErrNotLoggedIn)
} else {
	return errors.NewCode(errors.ErrNotLoggedIn)
}
```

This will automatically create an error with the specified code and the case as provided in input, allowing chaining.

## API calls

In order to have a structured way to interact with the server, we also spent a bit of effort making sure that we know what the server returns.

By default the answer now looks like this in case of success:
```json
{
	"RequestId": "115ceabe-d522-11ed-984c-18c04d0e6a41",
	"Status": "ERROR",
	"Details": {
		"Code": 1,
		"Message": "user name is invalid"
	}
}
```

Or in case of failure:
```json
{
	"RequestId": "16426e09-d522-11ed-984c-18c04d0e6a41",
	"Status": "SUCCESS",
	"Details": {
		"Id": "16426ecf-d522-11ed-984c-18c04d0e6a41"
	}
}
```

Building such structures is made easy with the [RequestData](pkg/middlewares/request.go) struct which comes with the request middleware. For example we could have a construct like this:
```go
func someHandler(w http.ResponseWriter, r *http.Request) {
	reqData, ok := middlewares.GetRequestDataFromContextOrFail(w, r)
	if !ok {
		return
	}

	out, err := doSomeProcessing()
	if err != nil {
		reqData.FailWithErrorAndCode(err, http.StatusInternalServerError, w)
		return
	}

	reqData.WriteDetails(ud, w)
}
```

It seems like an important feature to be able to know that we always have the same structure for an interaction with the server as it simplifies analyzing the failures for any endpoint.

## Command line

When trying to test that the endpoints of the server where correctly responding, we ended up (as usual) typing in a lot of curl commands. At some point it was nice to have a program to automate this.

We usually rely on a bash script but this time we tried to be a bit more sophisticated and use a cli tool. The idea is to have a syntax like:
```bash
cli sign-up toto 123456
```

Which will then automatically contact the server with an appropriate request and do this. Obviously we also want to add more features, like listing users, or logging in, etc. It quickly becomes a mess to only handle this through command line arguments. So we switched to use [cobra](https://github.com/spf13/cobra).

This automates the creation of command in a nice way and allows nesting very easily. The code can be found in the [client](cmd/client) app.

It also seems like a very interesting tool especially for the [sogserver](https://github.com/Knoblauchpilze/sogserver) project where we could created universes, ship description, etc realtively easily using such a tool rather than handcrafting some requests with curl.

## Authentication

We wanted to implement an authentication mechanism for a bit of time already. The authentication for now is provided in the [auth](pkg/auth) package. It consists of the following interface:
```go
type Token struct {
	User       uuid.UUID
	Value      string
	Expiration time.Time
}

type Authenticater interface {
	GenerateToken(user uuid.UUID, password string) (Token, error)
	GetToken(user uuid.UUID) (Token, error)
}
```

The idea is that for each user we maintain a set of tokens which can be used to authenticate them. For now the `Value` of the token is just a plain string.

Token can expire and need to be passed to the requests otherwise they fail. This mechanism is provided by the authentication middleware (see [here](pkg/middlewares/auth_context.go)).

For now we don't have any mechanism to refresh a token: it might be a good idea to automatically bump the validity of a token when it's being used or we could also leave the responsibility to the client to regularly perform a login.

The duration could also be configurable, but here as well there are downsides (what if we voluntarily ask for a token valid for 100 years?).

All in all it was interesting to implement and it seems like this mechanism could be reused as is in future projects.

## Testing

The first thing to mention is that testing in Go seems relatively easy. We stumbled upon an interesting resource on [mocking](https://www.myhatchpad.com/insight/mocking-techniques-for-go/).

We also started to use interfaces in general to mask certain behaviors. This leads to a better testability.

We also set up some CI (see [PR](https://github.com/Knoblauchpilze/go-server/pull/1)), and code coverage (see [PR](https://github.com/Knoblauchpilze/go-server/pull/2)). This gives a tangible visualization of the efforts we make on testing and not breaking stuff.

This was also further materialized by the badges:

[![codecov](https://codecov.io/gh/Knoblauchpilze/go-server/branch/master/graph/badge.svg?token=T0AX4BIS85)](https://codecov.io/gh/Knoblauchpilze/go-server)

Out of convenience, we also used the [testify](https://github.com/stretchr/testify) package to have some `assert.Condition` like syntax. This is not strictly necessary (some even advise against this) but as it was a first attempt, it made sense to try to use some tooling to make things easier.
