package adapter


type Info struct {
	Name string
	Impl string
	Description string
}

type InfoFn func() Info