package logger

// logger.MiddleWareInterface allows you to implement middlewares wrapping a logger handler.
// You can find basic middleware implementations in the middleware package.
type MiddlewareInterface func(handler HandlerInterface) HandlerInterface
