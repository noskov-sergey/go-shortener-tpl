package file

type Shorten struct {
	Uuid        int8   `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
