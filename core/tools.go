package core

type Tool interface {
	Name() string
	Run()
}
