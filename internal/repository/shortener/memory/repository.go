package memory

const (
	runeLen     = 52
	shortUrlLen = 8
)

type repository struct {
	data map[string]string
}

func New() *repository {
	return &repository{
		data: make(map[string]string),
	}
}
