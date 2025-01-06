package shortener

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Result string `json:"result"`
}

func ToResponse(result string) *shortenResponse {
	return &shortenResponse{
		Result: result,
	}
}
