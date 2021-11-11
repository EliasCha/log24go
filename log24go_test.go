package log24go

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLog24goInfo(t *testing.T) {
	var buf bytes.Buffer
	var testStr string = "TestLog24goInfo"
	t.Log(fmt.Sprintf("Testing with testStr: %s", testStr))

	logger := NewLog24go(&buf)
	logger.Info(testStr)
	bufStr := buf.String()
	t.Log(bufStr)

	tagVal := strings.Contains(bufStr, INFO)
	if !tagVal {
		t.Log(fmt.Sprintf("Buffer string does not contain string: %s", INFO))
		t.FailNow()
	}

	testVal := strings.Contains(bufStr, testStr)
	if !testVal {
		t.Log(fmt.Sprintf("Buffer string does not contain string: %s", testStr))
		t.FailNow()
	}
}

func TestLog24goWarn(t *testing.T) {
	var buf bytes.Buffer
	var testStr string = "TestLog24goWarn"
	t.Log(fmt.Sprintf("Testing with testStr: %s", testStr))

	logger := NewLog24go(&buf)
	logger.Warn(testStr)
	bufStr := buf.String()
	t.Log(bufStr)

	tagVal := strings.Contains(bufStr, WARN)
	if !tagVal {
		t.Log(fmt.Sprintf("Buffer string does not contain string: %s", WARN))
		t.FailNow()
	}

	testVal := strings.Contains(bufStr, testStr)
	if !testVal {
		t.Log(fmt.Sprintf("Buffer string does not contain string: %s", testStr))
		t.FailNow()
	}
}

func TestLog24goError(t *testing.T) {
	var buf bytes.Buffer
	var testStr string = "TestLog24goError"
	t.Log(fmt.Sprintf("Testing with testStr: %s", testStr))

	logger := NewLog24go(&buf)
	logger.Error(testStr)
	bufStr := buf.String()
	t.Log(bufStr)

	tagVal := strings.Contains(bufStr, ERROR)
	if !tagVal {
		t.Log(fmt.Sprintf("Buffer string does not contain string: %s", ERROR))
		t.FailNow()
	}

	testVal := strings.Contains(bufStr, testStr)
	if !testVal {
		t.Log(fmt.Sprintf("Buffer string does not contain string: %s", testStr))
		t.FailNow()
	}
}

func TestLog24goDebug(t *testing.T) {
	var buf bytes.Buffer
	var testStr string = "TestLog24goDebug"
	t.Log(fmt.Sprintf("Testing with testStr: %s", testStr))

	logger := NewLog24go(&buf)
	logger.Debug(testStr)
	bufStr := buf.String()
	t.Log(bufStr)

	tagVal := strings.Contains(bufStr, DEBUG)
	if !tagVal {
		t.Log(fmt.Sprintf("Buffer string does not contain string: %s", DEBUG))
		t.FailNow()
	}

	testVal := strings.Contains(bufStr, testStr)
	if !testVal {
		t.Log(fmt.Sprintf("Buffer string does not contain string: %s", testStr))
		t.FailNow()
	}
}
