package file

type Shorten struct {
	UUID        int8    `json:"uuid"`
	ShortURL    string  `json:"short_url"`
	OriginalURL string  `json:"original_url"`
	Username    *string `json:"username"`
}
