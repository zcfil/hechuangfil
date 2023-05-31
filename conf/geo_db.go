package conf

type GeoDB struct {
	Path string
}

func NewGeoDB(path string) func() *GeoDB {
	return func() *GeoDB {
		return &GeoDB{
			path,
		}
	}
}
