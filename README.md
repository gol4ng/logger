# Logger

<img src="logo.png" alt="gol4ng/logger: Golang logger" title="A new golang logger" align="right" width="200px">

[![Go Report Card](https://goreportcard.com/badge/github.com/gol4ng/logger)](https://goreportcard.com/report/github.com/gol4ng/logger)
[![Maintainability](https://api.codeclimate.com/v1/badges/a234f5fd2bcae54ed85e/maintainability)](https://codeclimate.com/github/gol4ng/logger/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a234f5fd2bcae54ed85e/test_coverage)](https://codeclimate.com/github/gol4ng/logger/test_coverage)
[![Build Status](https://travis-ci.org/gol4ng/logger.svg?branch=master)](https://travis-ci.org/gol4ng/logger)
[![GoDoc](https://godoc.org/github.com/gol4ng/logger?status.svg)](https://godoc.org/github.com/gol4ng/logger)

Gol4ng/Logger is another GO logger. The main line is to provide a friendly and fast API to send your log whenever you want. 

## Why another one?

When i start GO i searched a logger that can be **simple to use**, **efficient**, **multi output**, **multi formats** and quite easy to **extend**. 
That's why i created this logger with built-in [handlers](#Handlers)(process a log), [formatters](#formatters)(format log in another representation), [middlewares](#middlewares)(log modification before handler)

## Installation

`go get -u github.com/gol4ng/logger`

## Quick Start

```go
package main

import (
	"os"
	"runtime"
	
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

func main(){
	// logger will print on STDOUT with default line format
	l := logger.NewLogger(handler.Stream(os.Stdout, formatter.NewDefaultFormatter(formatter.WithContext(true))))
	
	l.Debug("Go debug information", logger.String("go_os", runtime.GOOS), logger.String("go_arch", runtime.GOARCH))
	// <debug> MyExample message {"go_arch":"amd64","go_os":"darwin"}
	
	l.Info("Another")
    //<info> Another
}
```

## Logger interface

This library expose some quite simple interfaces.

Simplest one
```go
type LogInterface interface {
	Log(message string, level Level, field ...Field)
}
```

The friendly one
```go
type LoggerInterface interface {
	LogInterface
	Debug(message string, field ...Field)
	Info(message string, field ...Field)
	Notice(message string, field ...Field)
	Warning(message string, field ...Field)
	Error(message string, field ...Field)
	Critical(message string, field ...Field)
	Alert(message string, field ...Field)
	Emergency(message string, field ...Field)
}
```

## Handlers

Handlers are log entry processor. It received a log entry and process it in order to send log to it's destination 

Available handlers:
- **stream** _it will write formatted log in io.Writer_
- **socket** _it will write formatted log in net.Conn_
- **chan** _send all entry in provided go channel_
- **gelf** _format to gelf and send to gelf server (TCP or UDP gzipped chunk)_
- **group** _it will send log to all provided child handlers_
- **rotate** _it will write formatted log in file and rotate file (mode: interval/archive)_
- **syslog** _send formatted log in syslog server_

## Formatters

The formatter convert log entry to a string representation (text, json, gelf...)
They are often inject in handler to do the conversion process

Available formatters:
- **default** `<info> My log message {"my_key":"my_value"}`
- **line** _it's just a `fmt.Sprintf` facade_ format:`%s %s %s` will produce `My log message info <my_key:my_value>`
- **gelf** _format log entry to gelf_ `{"version":"1.1","host":"my_fake_hostname","level":6,"timestamp":513216000.000,"short_message":"My log message","full_message":"<info> My log message [ <my key:my_value> ]","_my_key":"my_value"}`
- **json** _format log entry to json_ `{"Message":"My log message","Level":6,"Context":{"my_key":"my_value"}}`

## Middlewares

The middleware are handler decorator/wrapper. It will allow you to do some process around child handler 

Available middleware:
- **caller** _it will add caller file/line to context_ `<file:/my/example/path/file> <line:31>`
- **context** _it permit to have a default context value_ useful when you want to set global context value
- **error** _it will print provided handler error_ (can be configure to silent it)
- **filter** _it will permit to filter log entry_ level filter are available or you can use your own callback filter
- **placeholder** _it will replace log message placeholder with contextual value_
- **recover** _it will convert handler panic to error_
- **timestamp** _it will add timestamp to log context_

## Writers

Writers are use by handler to write/send log to appropriate destination

Available writer:
- **compress** _it will compress log to gzip/zlib_
- **gelf_chunked** _it will chunked log entry into gelf chunk_
- **rotate** _it will write in io.Writer and rotate writer on demand_
- **time_rotate** _it's a rotate writer that rotate with `time.Ticker`_

### Todo
- benchmark
- Implement all the handler
    - SSE http endpoint
    - websocket server 
    - socket server
    - https://craig.is/writing/chrome-logger
    - fingercross
    - grpc / protobuf
    - curl
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
    - proto
    - slack
    - flowdoc
    - fluentd
    - logstash
    - mongodb
    - wildfire
 
### Idea

- log server with log routing
