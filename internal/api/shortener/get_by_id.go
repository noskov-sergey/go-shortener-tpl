package shortener

import (
	"net/http"
)

func (i *Implementation) GetByID(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortUrl := req.PathValue("id")
	if shortUrl == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := i.service.GetByID(shortUrl)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
