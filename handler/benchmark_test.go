package handler_test

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"testing"
)

type NopWriter struct{}

func (w *NopWriter) Write(_ []byte) (n int, err error) {
	return 0, nil
}

type handlerToTest struct {
	name      string
	handler logger.HandlerInterface
}

func handlersToBench() []handlerToTest {
	return []handlerToTest{
		{
			"group nop",
			handler.Group(
				logger.NopHandler,
				logger.NopHandler,
			),
		},
		{
			"stream nop formatter nop writer",
			handler.Stream(&NopWriter{}, logger.NewNopFormatter()),
		},
	}
}

func getEntries() []logger.Entry {
	return []logger.Entry{
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

func BenchmarkHandler_Handle(b *testing.B) {
	for _, entry := range getEntries() {
		for _, h := range handlersToBench() {
			b.Run(h.name+"_"+entry.Message, func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = h.handler(entry)
				}
			})
		}
	}
}
