package handler

import (
	"github.com/gol4ng/logger"
)

type Filter struct {
	handler  logger.HandlerInterface
	filterFn func(logger.Entry) bool
}

func (f *Filter) Handle(entry logger.Entry) error {
	if f.filterFn(entry) {
		return nil
	}

	return f.handler.Handle(entry)
}

// exclude logs that have a higher level than a given level
func NewMinLevelFilter(handler logger.HandlerInterface, level logger.Level) *Filter {
	return NewFilter(handler, func(entry logger.Entry) bool {
		return entry.Level > level
	})
}

func NewMinLevelWrapper(level logger.Level) logger.WrapperHandler {
	return func(h logger.HandlerInterface) logger.HandlerInterface {
		return NewMinLevelFilter(h, level)
	}
}

// exclude logs that have a level that are not between two given levels
func NewRangeLevelFilter(handler logger.HandlerInterface, minLevel logger.Level, maxLevel logger.Level) *Filter {
	if minLevel <= maxLevel {
		panic("invalid logger range level : Min level must be lower than max level")
	}

	return NewFilter(handler, func(entry logger.Entry) bool {
		return entry.Level > minLevel || entry.Level < maxLevel
	})
}

func NewRangeLevelWrapper(minLevel logger.Level, maxLevel logger.Level) logger.WrapperHandler {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return NewRangeLevelFilter(handler, minLevel, maxLevel)
	}
}

func NewFilter(handler logger.HandlerInterface, filterFn func(logger.Entry) bool) *Filter {
	return &Filter{handler: handler, filterFn: filterFn}
}

func NewFilterWrapper(filterFn func(logger.Entry) bool) logger.WrapperHandler {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return NewFilter(handler, filterFn)
	}
}
