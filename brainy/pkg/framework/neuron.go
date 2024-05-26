package framework

type Neuron interface {
	Start() error
	Stop() error
}

type NeuronType int

const (
	Singleton NeuronType = iota
	Prototype
)
