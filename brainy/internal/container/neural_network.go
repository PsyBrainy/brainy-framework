package container

import (
	"brainy-framework/brainy/pkg/framework"
	"reflect"
)

type NeuralNetwork struct {
	singletonNeurons map[reflect.Type]interface{}
	prototypeNeurons map[reflect.Type]reflect.Type
}

func NewNeuralNetwork() *NeuralNetwork {
	return &NeuralNetwork{
		singletonNeurons: make(map[reflect.Type]interface{}),
		prototypeNeurons: make(map[reflect.Type]reflect.Type),
	}
}

func (nn *NeuralNetwork) Register(neuron interface{}, neuronType framework.NeuronType) {
	neuronTypeOf := reflect.TypeOf(neuron)
	if neuronType == framework.Singleton {
		nn.singletonNeurons[neuronTypeOf] = neuron
	} else if neuronType == framework.Prototype {
		nn.prototypeNeurons[neuronTypeOf] = neuronTypeOf
	}
}

func (nn *NeuralNetwork) Resolve(neuron interface{}) interface{} {
	neuronTypeOf := reflect.TypeOf(neuron)
	if singleton, exists := nn.singletonNeurons[neuronTypeOf]; exists {
		return singleton
	} else if prototype, exists := nn.prototypeNeurons[neuronTypeOf]; exists {
		return reflect.New(prototype.Elem()).Interface()
	}
	return nil
}
