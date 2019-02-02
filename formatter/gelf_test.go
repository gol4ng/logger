package formatter_test

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"

	"github.com/gol4ng/logger/formatter"

	"github.com/gol4ng/logger"
)

func TestGelfFormatter(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	monkey.Patch(os.Hostname, func() (string, error) { return "my_fake_hostname", nil })
	defer monkey.UnpatchAll()

	gelfFormatter, _ := formatter.NewGelf()

	assert.Equal(
		t,
		`{"version":"1.1","host":"my_fake_hostname","short_message":"My log message","full_message":"TODO","timestamp":513216000,"level":6}`,
		gelfFormatter.Format(logger.Entry{Message: "My log message", Level: logger.InfoLevel, Context: &map[string]interface{}{"my_key": "my_value"}}),
	)
}
