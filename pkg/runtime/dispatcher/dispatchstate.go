package dispatcher

import (
	"context"

	"github.com/covarity/echo/pkg/adapter"
)

// dispatchState keeps the input/output state during the dispatch to a handler. It is used as temporary
// memory location to keep ephemeral state, thus avoiding garbage creation.
type dispatchState struct {
	session *session
	ctx     context.Context

	checkResult adapter.CheckResult

}
