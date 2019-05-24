package middleware

import (
	"github.com/gol4ng/logger"
)

func Filter(filterFn func(logger.Entry) bool) logger.Middleware {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			if filterFn(entry) {
				return nil
			}
			return handler(entry)
		}
	}
}

// exclude logs that have a minor level than a given level
func MinLevelFilter(level logger.Level) logger.Middleware {
	return Filter(func(entry logger.Entry) bool {
		return entry.Level > level
	})
}

// exclude logs that have a major level than a given level
func MaxLevelFilter(level logger.Level) logger.Middleware {
	return Filter(func(entry logger.Entry) bool {
		return entry.Level < level
	})
}

// exclude logs that have a level that are not between two given levels
func RangeLevelFilter(level1 logger.Level, level2 logger.Level) logger.Middleware {
	if level1 > level2 {
		l := level1
		level1 = level2
		level2 = l
	}
	return Filter(func(entry logger.Entry) bool {
		return entry.Level < level1 || entry.Level > level2
	})
}
