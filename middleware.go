package logger

type MiddlewareInterface func(handler HandlerInterface) HandlerInterface
