package dispatcher

import (
	"context"

	adptTmpl "github.com/covarity/echo/api/adapter/model/v1"
)

type Requester interface {
	Request(destination string) error

	Flush() error

	Done()
}

type requester struct {
	impl        *Impl
	ctx         context.Context
	state       *dispatchState
	destination string
}

var _ Requester = &requester{}

func (r *requester) clear() {
	r.impl = nil
	r.ctx = nil
	r.state = nil
}

func (r *requester) Request(destination string) error {
	s := r.impl.getSession(r.ctx, adptTmpl.TemplateVariety_TEMPLATE_VARIETY_REQUEST, destination)
	s.requestState = r.state
	err := s.dispatch()
	if err == nil {
		err = s.err
	}
	r.impl.putSession(s)
	return err
}

func (r *requester) Flush() error {
	s := r.impl.getSession(r.ctx, adptTmpl.TemplateVariety_TEMPLATE_VARIETY_REQUEST, "")
	s.requestState = r.state
	err := s.err
	r.impl.putSession(s)
	return err
}

func (r *requester) Done() {
	r.impl.putRequester(r)
}
