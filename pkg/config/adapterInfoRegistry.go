package config

import (
	"fmt"
	"strings"

	"github.com/covarity/echo/pkg/adapter"
)

type adapterInfoRegistry struct {
	adapterInfosByName map[string]*adapter.Info
}

type handlerBuilderValidator func(hndlrBuilder adapter.HandlerBuilder, t string) (bool, string)

// newRegistry returns a new adapterInfoRegistry.
func newRegistry(infos []adapter.InfoFn, hndlrBldrValidator handlerBuilderValidator) *adapterInfoRegistry {
	r := &adapterInfoRegistry{make(map[string]*adapter.Info)}
	for _, info := range infos {
		adptInfo := info()
		if a, ok := r.adapterInfosByName[adptInfo.Name]; ok {
			// panic only if 2 different adapter.Info objects are trying to identify by the
			// same Name.
			msg := fmt.Sprintf("Duplicate registration for '%s' : old = %v new = %v", a.Name, adptInfo, a)
			panic(msg)
		} else {
			if adptInfo.NewBuilder == nil {
				// panic if adapter has not provided the NewBuilder func.
				msg := fmt.Sprintf("Adapter info %v from adapter %s has nil NewBuilder", adptInfo, adptInfo.Name)
				panic(msg)
			}
			if ok, errMsg := doesBuilderSupportsTemplates(adptInfo, hndlrBldrValidator); !ok {
				// panic if an Adapter's HandlerBuilder does not implement interfaces that it says it wants to support.
				msg := fmt.Sprintf("HandlerBuilder from adapter %s does not implement the required interfaces"+
					" for the templates it supports: %s", adptInfo.Name, errMsg)
				// log.Error(msg)
				panic(msg)
			}
			r.adapterInfosByName[adptInfo.Name] = &adptInfo
		}
	}
	return r
}

// AdapterInfoMap returns the known adapter.Infos, indexed by their names.
func AdapterInfoMap(handlerRegFns []adapter.InfoFn, hndlrBldrValidator handlerBuilderValidator) map[string]*adapter.Info {
	return newRegistry(handlerRegFns, hndlrBldrValidator).adapterInfosByName
}

func doesBuilderSupportsTemplates(info adapter.Info, hndlrBldrValidator handlerBuilderValidator) (bool, string) {
	handlerBuilder := info.NewBuilder()
	resultMsgs := make([]string, 0)
	for _, t := range info.SupportedTemplates {
		if ok, errMsg := hndlrBldrValidator(handlerBuilder, t); !ok {
			resultMsgs = append(resultMsgs, errMsg)
		}
	}
	if len(resultMsgs) != 0 {
		return false, strings.Join(resultMsgs, "\n")
	}
	return true, ""
}
