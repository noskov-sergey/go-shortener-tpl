package shortener

import (
	"encoding/json"
	"io"
	"net/http"
)

func (i *Implementation) shortenHandler(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var data shortenRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}

	s, err := i.service.Create(data.URL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	got, err := json.Marshal(ToResponse(i.cfg.baseURL + "/" + s))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(got)
}
