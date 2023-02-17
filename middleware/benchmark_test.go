package middleware_test

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
	"testing"
)

type middlewareToTest struct {
	name       string
	middleware logger.MiddlewareInterface
}

func middlewaresToBench() []middlewareToTest {
	return []middlewareToTest{
		{
			"caller 0",
			middleware.Caller(0),
		},
		{
			"context",
			middleware.Context(logger.NewContext(
				logger.String("my_base_key_1", "my_base_value_1"),
				logger.String("my_base_key_2", "my_base_value_2"),
			)),
		},
		{
			"min level filter",
			middleware.MinLevelFilter(logger.WarningLevel),
		},
		{
			"max level filter",
			middleware.MaxLevelFilter(logger.WarningLevel),
		},
		{
			"range level filter",
			middleware.RangeLevelFilter(logger.WarningLevel, logger.InfoLevel),
		},
		{
			"placeholder",
			middleware.Placeholder(),
		},
	}
}

func getEntries() []logger.Entry {
	return []logger.Entry{
		{
			Message: "test %my_key% message",
			Level:   logger.WarningLevel,
			Context: logger.NewContext().
				Add("my_key", "my_value").
				Add("my_key2", "my_value2").
				Add("my_key3", "my_value3").
				Add("my_key4", "my_value4").
				Add("my_key5", "my_value5"),
		},
		{
			Message: "test message",
			Level:   logger.WarningLevel,
			Context: logger.NewContext().
				Add("my_key", "my_value").
				Add("my_key2", "my_value2").
				Add("my_key3", "my_value3").
				Add("my_key4", "my_value4").
				Add("my_key5", "my_value5"),
		},
	}
}

func BenchmarkMiddleware(b *testing.B) {
	for _, entry := range getEntries() {
		for _, m := range middlewaresToBench() {
			b.Run(m.name+"_"+entry.Message, func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					m.middleware(logger.NopHandler)(entry)
				}
			})
		}
	}
}
