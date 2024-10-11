package core

type DpRepository interface {
	Insert(dp *Dp) error
	Retrieve() ([]Dp, error)
}
