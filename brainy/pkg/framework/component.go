package framework

type Component interface {
	Start() error
	Stop() error
}
