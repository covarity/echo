package config

import (
	"fmt"

	"github.com/covarity/echo/pkg/adapter"
)

type adapterInfoRegistry struct {
	adapterInfosByName map[string]*adapter.Info
}

// newRegistry returns a new adapterInfoRegistry.
func newRegistry(infos []adapter.InfoFn) *adapterInfoRegistry {
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
			r.adapterInfosByName[adptInfo.Name] = &adptInfo
		}
	}
	return r
}

// AdapterInfoMap returns the known adapter.Infos, indexed by their names.
func AdapterInfoMap(handlerRegFns []adapter.InfoFn) map[string]*adapter.Info {
	return newRegistry(handlerRegFns).adapterInfosByName
}
