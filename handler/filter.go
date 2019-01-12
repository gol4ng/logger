package handler

import (
	"github.com/instabledesign/logger"
)

type Filter struct {
	handler  logger.HandlerInterface
	filterFn func(logger.Entry) bool
}

func (f *Filter) Handle(e logger.Entry) error {
	if f.filterFn(e) {
		return nil
	}

	return f.handler.Handle(e)
}

// exclude logs that have a higher level than a given level
func NewMinLevelFilter(h logger.HandlerInterface, lvl logger.Level) *Filter {
	return &Filter{h, func(e logger.Entry) bool {
		return e.Level < lvl
	}}
}

// exclude logs that have a level that are not between two given levels
func NewRangeLevelFilter(h logger.HandlerInterface, minLvl logger.Level, maxLvl logger.Level) *Filter {
	if minLvl >= maxLvl {
		panic("invalid logger range level : Min level must be lower than max level")
	}

	return &Filter{h, func(e logger.Entry) bool {
		return e.Level < minLvl || e.Level > maxLvl
	}}
}

func NewFilter(h logger.HandlerInterface, f func(e logger.Entry) bool) *Filter {
	return &Filter{handler: h, filterFn: f}
}
