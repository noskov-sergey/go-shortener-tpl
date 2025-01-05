package shortener

import (
	"io"
	"net/http"
)

func (i *Implementation) Create(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

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
	res.Write([]byte("http://" + req.Host + "/" + s))
}
