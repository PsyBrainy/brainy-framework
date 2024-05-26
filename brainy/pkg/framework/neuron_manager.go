package framework

import (
	"fmt"
	"sync"
)

type NeuronManager struct {
	neurons map[string]Neuron
	mu      sync.Mutex
}

func NewNeuronManager() *NeuronManager {
	return &NeuronManager{
		neurons: make(map[string]Neuron),
	}
}

func (nm *NeuronManager) RegisterNeuron(name string, neuron Neuron) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	nm.neurons[name] = neuron
}

func (nm *NeuronManager) Start() error {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	for name, neuron := range nm.neurons {
		if err := neuron.Start(); err != nil {
			return fmt.Errorf("error starting neuron %s: %v", name, err)
		}
	}
	return nil
}

func (nm *NeuronManager) Stop() error {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	for name, neuron := range nm.neurons {
		if err := neuron.Stop(); err != nil {
			return fmt.Errorf("error stopping neuron %s: %v", name, err)
		}
	}
	return nil
}
