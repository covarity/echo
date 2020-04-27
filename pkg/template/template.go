package template

import (
	"context"
	"fmt"

	adptTmpl "github.com/covarity/echo/api/adapter/model/v1"
	pb "github.com/covarity/echo/api/policy/v1alpha1"
	"github.com/covarity/echo/pkg/adapter"
	"github.com/gogo/protobuf/proto"
)

type (

	// Repository defines all the helper functions to access the generated template specific types and fields.
	Repository interface {
		GetTemplateInfo(template string) (Info, bool)
		SupportsTemplate(hndlrBuilder adapter.HandlerBuilder, tmpl string) (bool, string)
	}

	// TypeEvalFn evaluates an expression and returns the ValueType for the expression.
	TypeEvalFn func(string) (pb.ValueType, error)
	// InferTypeFn does Type inference from the Instance.params proto message.
	InferTypeFn func(proto.Message, TypeEvalFn) (proto.Message, error)
	// SetTypeFn dispatches the inferred types to handlers
	SetTypeFn func(types map[string]proto.Message, builder adapter.HandlerBuilder)
	// DispatchRequestFn dispatches the requests to the handler.
	DispatchRequestFn func(ctx context.Context, handler adapter.Handler) error

	// BuilderSupportsTemplateFn check if the handlerBuilder supports template.
	BuilderSupportsTemplateFn func(hndlrBuilder adapter.HandlerBuilder) bool

	// HandlerSupportsTemplateFn check if the handler supports template.
	HandlerSupportsTemplateFn func(hndlr adapter.Handler) bool
	// Info contains all the information related a template like
	// Default instance params, type inference method etc.
	Info struct {
		Name                    string
		Impl                    string
		Variety                 adptTmpl.TemplateVariety
		BldrInterfaceName       string
		HndlrInterfaceName      string
		DispatchRequest         DispatchRequestFn
		BuilderSupportsTemplate BuilderSupportsTemplateFn
		HandlerSupportsTemplate HandlerSupportsTemplateFn
	}

	// templateRepo implements Repository
	repo struct {
		info map[string]Info

		allSupportedTmpls  []string
		tmplToBuilderNames map[string]string
	}
)

// NewRepository returns an implementation of Repository
func NewRepository(templateInfos map[string]Info) Repository {
	if templateInfos == nil {
		return repo{
			info:               make(map[string]Info),
			allSupportedTmpls:  make([]string, 0),
			tmplToBuilderNames: make(map[string]string),
		}
	}

	allSupportedTmpls := make([]string, 0, len(templateInfos))
	tmplToBuilderNames := make(map[string]string)

	for t, v := range templateInfos {
		allSupportedTmpls = append(allSupportedTmpls, t)
		tmplToBuilderNames[t] = v.BldrInterfaceName
	}
	return repo{info: templateInfos, tmplToBuilderNames: tmplToBuilderNames, allSupportedTmpls: allSupportedTmpls}
}

func (t repo) GetTemplateInfo(template string) (Info, bool) {
	if v, ok := t.info[template]; ok {
		return v, true
	}
	return Info{}, false
}

func (t repo) SupportsTemplate(hndlrBuilder adapter.HandlerBuilder, tmpl string) (bool, string) {
	i, ok := t.GetTemplateInfo(tmpl)
	if !ok {
		return false, fmt.Sprintf("Supported template %v is not one of the allowed supported templates %v", tmpl, t.allSupportedTmpls)
	}

	if b := i.BuilderSupportsTemplate(hndlrBuilder); !b {
		return false, fmt.Sprintf("HandlerBuilder does not implement interface %s. "+
			"Therefore, it cannot support template %v", t.tmplToBuilderNames[tmpl], tmpl)
	}

	return true, ""
}
