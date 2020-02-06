package dispatcher

import (
	"context"
	"fmt"

	"github.com/covarity/echo/pkg/adapter"
)

// dispatchState keeps the input/output state during the dispatch to a handler. It is used as temporary
// memory location to keep ephemeral state, thus avoiding garbage creation.
type dispatchState struct {
	session *session
	ctx     context.Context
	destination *routing.Destination
	checkResult adapter.CheckResult
	err         error
}

func (ds *dispatchState) clear() {
	ds.session = nil
	ds.ctx = nil
	ds.checkResult = adapter.CheckResult{}
}

func (ds *dispatchState) invokeHandler(interface{}) {
	reachedEnd := false

	defer func() {
		if reachedEnd {
			return
		}
		r := recover()
		ds.err = fmt.Errorf("panic during handler dispatch: %v", r)
		ds.session.completed <- ds
	}()
	fmt.Printf("invokeHandler")
	ds.des
	ds.session.completed <- ds
	reachedEnd = true

}
