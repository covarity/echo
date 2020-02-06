package adapter

import (
	"fmt"
	"time"
)

type CheckResult struct {
	Status string
	// ValidDuration represents amount of time for which this result can be considered valid.
	ValidDuration time.Duration
	// ValidUseCount represents the number of uses for which this result can be considered valid.
	ValidUseCount int32
}
// String
func (r *CheckResult) String() string {
	return fmt.Sprintf("CheckResult: status:%s, duration:%d, usecount:%d", r.Status)
}

// IsDefault returns true if the CheckResult is in its zero state
func (r *CheckResult) IsDefault() bool {
	return r.Status == "OK" && r.ValidDuration == 0 && r.ValidUseCount == 0
}

