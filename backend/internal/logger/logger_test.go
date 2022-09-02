package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var errSome = fmt.Errorf("some error")

func TestError(t *testing.T) {
	rescueStdout := os.Stdout
	reader, writer, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = writer
	logger := NewLogger()

	logger.Error(errSome)

	writer.Close()
	out, err := ioutil.ReadAll(reader)
	require.NoError(t, err)
	os.Stdout = rescueStdout
	require.Contains(t, string(out), errSome.Error())
}
