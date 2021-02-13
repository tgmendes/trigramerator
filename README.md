# Go Service Template

A simple project template to start a Go web service. Highly inspired by [ArdanLabs Service](https://github.com/ardanlabs/service).


## What's included?

This template has the basics to start running a Go server:

* A main file with a server bootstrap (handling graceful shutdowns via channels)

* A Config package to handle configurations from environment variables

* An HTTP package, wrapping [httprouter](https://github.com/julienschmidt/httprouter), and including
utilities for HTTP responses

* A simple HTTP client to make external web requests

* A basic example on how to run and test the service

## What's missing?

* Middleware examples

* Tracing and advanced logging 
  
* Other advanced features


