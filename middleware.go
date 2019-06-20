package logger

// MiddlewareInterface is a function to decorate handler
// See middleware implementations in the middleware package.
type MiddlewareInterface func(handler HandlerInterface) HandlerInterface
