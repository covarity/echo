package adapter

import (
	"fmt"
)

type CheckResult struct {
	Status string
}

func (r *CheckResult) String() string {
	return fmt.Sprintf("CheckResult: status:%s, duration:%d, usecount:%d", r.Status)
}
