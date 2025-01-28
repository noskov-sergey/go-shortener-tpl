package model

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

func ToResponse(result string) *ShortenResponse {
	return &ShortenResponse{
		Result: result,
	}
}

type ShortenUserRequest struct {
	ShortURL string `json:"short_url"`
	URL      string `json:"original_url"`
}

func ToUserResponse(data []Shortener) []ShortenUserRequest {
	var res []ShortenUserRequest
	for _, d := range data {
		res = append(res, ShortenUserRequest{
			URL:      d.URL,
			ShortURL: d.ShortURL,
		})
	}
	return res
}

type Shortener struct {
	URL      string
	ShortURL string
	Username string
}
