package shortener

import (
	"net/http"
)

func (i *Implementation) getByIDHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL := req.PathValue("id")
	if shortURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := i.service.GetByID(shortURL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
