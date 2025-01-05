package shortener

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_shortener "github.ru/noskov-sergey/go-shortener-tpl/internal/api/shortener/mocks"
)

//func TestImplementation_GetByID_Success(t *testing.T) {
//	t.Parallel()
//	ctrl := gomock.NewController(t)
//	api := mock_shortener.NewMockservice(ctrl)
//	td := New(api)
//
//	r := httptest.NewRequest(http.MethodGet, "/sXaeNgYs", nil)
//	w := httptest.NewRecorder()
//
//	api.EXPECT().
//		GetByID(gomock.Any()).
//		Return("github.ru/synoskov/", nil)
//
//	td.GetByID(w, r)
//
//	assert.Equal(t, http.StatusTemporaryRedirect, w.Code, "Код ответа не совпадает с ожидаемым")
//}

func TestImplementation_GetByID_Error(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	api := mock_shortener.NewMockservice(ctrl)
	td := New(api)

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	td.GetByID(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Код ответа не совпадает с ожидаемым")
}
