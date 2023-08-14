//go:build !functional
// +build !functional

package sarama

import (
	"flag"
	"log"
	"os"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	defer goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick"))
	flag.Parse()
	if f := flag.Lookup("test.v"); f != nil && f.Value.String() == "true" {
		Logger = log.New(os.Stderr, "[DEBUG] ", log.Lmicroseconds|log.Ltime)
	}
	return m.Run()
}
