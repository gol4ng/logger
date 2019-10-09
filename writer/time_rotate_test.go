package writer_test

import (
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger/mocks"
	"github.com/gol4ng/logger/writer"
)

func TestTimeRotateWriter_StartWithError(t *testing.T) {
	tickerChan := make(chan time.Time, 1)
	monkey.Patch(time.NewTicker, func(d time.Duration) *time.Ticker {
		assert.Equal(t, 50*time.Millisecond, d)

		return &time.Ticker{
			C: tickerChan,
		}
	})
	defer monkey.UnpatchAll()

	mockRotateWriter := mocks.RotateWriter{}
	mockRotateWriter.On("Rotate").Return(func() error { return errors.New("fake_rotate_error") })

	tr := writer.TimeRotateWriter{
		RotateWriter: &mockRotateWriter,
		Interval:     50 * time.Millisecond,
		PanicHandler: func(err error) {
			assert.EqualError(t, err, "fake_rotate_error")
		},
	}
	tr.Start()
	mockRotateWriter.AssertNotCalled(t, "Rotate")
	tickerChan <- time.Now()
	time.Sleep(1 * time.Millisecond) //we have to wait a bit to be shure Rotate call has been made
	mockRotateWriter.AssertCalled(t, "Rotate")
}

func TestNewTimeRotateWriter(t *testing.T) {
	mockRotateWriter := mocks.RotateWriter{}
	mockRotateWriter.On("Rotate").Return(func() error { return nil })

	writer.NewTimeRotateWriter(&mockRotateWriter, 50*time.Millisecond)
	mockRotateWriter.AssertNotCalled(t, "Rotate")
}

func TestNewTimeRotateFileWriter_WithError(t *testing.T) {
	_, err := writer.NewTimeRotateFileWriter(func(io.Writer) (io.Writer, error) {
		return &os.File{}, errors.New("fake_file_provider_error")
	}, 1*time.Second)
	assert.EqualError(t, err, "fake_file_provider_error")
}

func TestNewTimeRotateFileWriter(t *testing.T) {
	_, err := writer.NewTimeRotateFileWriter(func(io.Writer) (io.Writer, error) {
		return &os.File{}, nil
	}, 1*time.Second)
	assert.Nil(t, err)
}
