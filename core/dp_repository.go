package core

type DpRepository interface {
	Create(dp *Dp) error

	Get() ([]Dp, error)

	Search(provider string, product string, tier string, region string, customer string) []Dp
}
