package template

import (
	"context"
	"fmt"

	adapter_model_v1 "github.com/covarity/echo/api/adapter/model/v1"
	"github.com/covarity/echo/pkg/adapter"
	"github.com/covarity/echo/pkg/template"
	"github.com/covarity/echo/templates/synthetic"
)

var (
	// SupportedTmplInfo main struct containing support templates
	SupportedTmplInfo = map[string]template.Info{
		synthetic.TemplateName: {
			Name:               synthetic.TemplateName,
			Impl:               "synthetic",
			Variety:            adapter_model_v1.TemplateVariety_TEMPLATE_VARIETY_REQUEST,
			BldrInterfaceName:  synthetic.TemplateName + "." + "HandlerBuilder",
			HndlrInterfaceName: synthetic.TemplateName + "." + "Handler",
			BuilderSupportsTemplate: func(hndlrBuilder adapter.HandlerBuilder) bool {
				_, ok := hndlrBuilder.(synthetic.HandlerBuilder)
				return ok
			},
			HandlerSupportsTemplate: func(hndlr adapter.Handler) bool {
				_, ok := hndlr.(synthetic.Handler)
				return ok
			},
			// DispatchReport dispatches the instances to the handler.
			DispatchRequest: func(ctx context.Context, handler adapter.Handler) error {
				// Invoke the handler.
				if err := handler.(synthetic.Handler).HandleSynthetic(ctx); err != nil {
					fmt.Print("error")
					return fmt.Errorf("failed to report all values: %v", err)
				}
				return nil
			},
		},
	}
)
