package framework

import (
	"fmt"
	"sync"
)

type Framework struct {
	components map[string]Component
	mu         sync.Mutex
}

func NewFramework() *Framework {
	return &Framework{
		components: make(map[string]Component),
	}
}

func (f *Framework) RegisterComponent(name string, component Component) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.components[name] = component
}

func (f *Framework) Start() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	for name, component := range f.components {
		if err := component.Start(); err != nil {
			return fmt.Errorf("error starting component %s: %v", name, err)
		}
	}
	return nil
}

func (f *Framework) Stop() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	for name, component := range f.components {
		if err := component.Stop(); err != nil {
			return fmt.Errorf("error stopping component %s: %v", name, err)
		}
	}
	return nil
}
