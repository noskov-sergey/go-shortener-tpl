package shortener

import (
	"io"
	"net/http"
)

func (i *Implementation) createHandler(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	s, err := i.service.Create(string(body))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(i.baseURL + "/" + s))
}
