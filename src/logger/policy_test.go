package logger

import (
	"fmt"
	"os"
	"testing"
	"github.com/go-testutil/testutil"
)

func TestLogFile_LogPolicy(t *testing.T) {
	testName := "TestLogPolicy"

	// Test normal static log creation
	t.Run(testName+"=1", func(t *testing.T) {
		l, err := File(testName)
		testutil.AssertNil(t, err, fmt.Sprintf("%s; \"%s\"\n", err, testName))
		testutil.AssertNotNil(t, l, fmt.Sprintf("*LogFile is nil: \"%s\"\n", testName))

		p := l.LogPolicy()
		testutil.AssertFalse(t, p.IsSizeLimited(), "Expected static file policy, got "+p.String())
		testutil.AssertTrue(t, p.isNone(), "Expected Static file policy, got "+p.String())
		testutil.AssertFalse(t, p.IsDaily(), "Expected Static file policy, got "+p.String())
		testutil.AssertFalse(t, p.IsTimed(), "Expected Static file policy, got "+p.String())

		name := l.LogFilename()
		defer os.Remove(name)
		testutil.AssertNotEmptyString(t, name, fmt.Sprintf("Value: \"%s\"\n", testName))
		l.Close()

		_, ok := os.Stat(name)
		testutil.AssertNil(t, ok, fmt.Sprintf("%s; File: \"%s\"\n", err, testName))
	})
}
