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
