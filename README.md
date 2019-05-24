# Logger

[![GoDoc](https://godoc.org/github.com/gol4ng/logger?status.svg)](https://godoc.org/github.com/gol4ng/logger)

Gol4ng/Logger is another GO logger. The main line is to provide a friendly and fast API to send your log whenever you want. 

## Installation

`go get -u github.com/gol4ng/logger`

## Quick Start

```go
//TODO write example
```

## Logger interface

This library expose some quite simple interfaces.

Simplest one
```go
type LogInterface interface {
	Log(message string, level Level, context *Context) error
}
```

The friendly one
```go
type LoggerInterface interface {
	LogInterface
	Debug(message string, context *Context) error
	Info(message string, context *Context) error
	Notice(message string, context *Context) error
	Warning(message string, context *Context) error
	Error(message string, context *Context) error
	Critical(message string, context *Context) error
	Alert(message string, context *Context) error
	Emergency(message string, context *Context) error
}
```

## Handlers

The handler is the main part of the logger

## Formatters

The formatter is ... 

## Writers

The writer is ... 

### Todo
- benchmark
- improve err handling
- Implement all the handler
    - SSE http endpoint
    - websocket server 
    - socket server
    - https://craig.is/writing/chrome-logger
    - rotating file
    - fingercross
    - syslog
    - syslog udp
    - grpc / protobuff
    - curl
    - socket
    - Gelf
    - Mail
    - Slack
    - hipchat
    - amqp
    - redis
    - elasticsearch
    - newrelic
    - browser console
    - couchdb
    - cube
    - ifttt
    - InsightOps
    - mandrill
    - pushover
    - raven
    - rollbar
    - sampling
    - LogEntries
    - ???
    
- Implement all the formatter
    - html
    - gelf
    - proto
    - slack
    - flowdoc
    - fluentd
    - logstash
    - mongodb
    - wildfire
 
### Idea

- add shortcut to log time.now
- log server with log routing
