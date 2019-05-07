package logger

type Middleware func(handler HandlerInterface) HandlerInterface
