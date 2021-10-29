package formatter_test

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"testing"
)

type formatterToTest struct {
	name      string
	formatter logger.FormatterInterface
}

func formattersToBench() []formatterToTest {
	return []formatterToTest{
		{
			"default",
			formatter.NewDefaultFormatter(),
		},
		{
			"default colored",
			formatter.NewDefaultFormatter(formatter.WithColor(true)),
		},
		{
			"default context",
			formatter.NewDefaultFormatter(formatter.WithContext(true)),
		},
		{
			"default full",
			formatter.NewDefaultFormatter(formatter.WithContext(true), formatter.WithColor(true)),
		},
		{"gelf",
			formatter.NewGelf(),
		},
		{
			"json",
			formatter.NewJSONEncoder(),
		},
		{
			"line",
			formatter.NewLine(formatter.LineFormatLevelFirst),
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

func BenchmarkFormatter_Format(b *testing.B) {
	for _, entry := range getEntries() {
		for _, f := range formattersToBench() {
			b.Run(f.name+"_"+entry.Message, func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					f.formatter.Format(entry)
				}
			})
		}
	}
}
