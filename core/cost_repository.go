package core

type CostRepository interface {
	Insert(dp *Cost) error
	Retrieve() ([]Cost, error)
}
