package logger

// MiddlewareInterface is a function to decorate handler
// See middleware implementations in the middleware package.
type MiddlewareInterface func(HandlerInterface) HandlerInterface

// Middlewares was a collection of middleware
// the slice order matter when middleware are compose with DecorateHandler
type Middlewares []MiddlewareInterface

// Decorate will return the handler after decorate with middlewares
func (m Middlewares) Decorate(handler HandlerInterface) HandlerInterface {
	return DecorateHandler(handler, m...)
}

// DecorateHandler will return the handler after decorate with middlewares
func DecorateHandler(handler HandlerInterface, middlewares ...MiddlewareInterface) HandlerInterface {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

// MiddlewareStack helper return Middlewares slice from given middlewares
func MiddlewareStack(middlewares ...MiddlewareInterface) Middlewares {
	return middlewares
}
