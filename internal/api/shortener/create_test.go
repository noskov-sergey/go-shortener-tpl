package shortener

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_shortener "github.ru/noskov-sergey/go-shortener-tpl/internal/api/shortener/mocks"
)

func TestImplementation_Create_Success(t *testing.T) {
	type want struct {
		code int
		body []byte
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1 - handler work good",
			want: want{
				code: http.StatusCreated,
				body: []byte(`http://shortener.com`),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			api := mock_shortener.NewMockservice(ctrl)
			td := New(api, "BaseURL")

			body := bytes.NewReader(test.want.body)
			r := httptest.NewRequest(http.MethodPost, "/", body)
			w := httptest.NewRecorder()

			api.EXPECT().
				Create("http://shortener.com").
				Return("AAAbbbCC", nil)

			td.createHandler(w, r)

			assert.Equal(t, test.want.code, w.Code, "Code doesnt match")
		})
	}
}

func TestImplementation_Create_Error(t *testing.T) {
	type want struct {
		code int
		body []byte
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "negative test #1 - error work",
			want: want{
				code: http.StatusBadRequest,
				body: []byte(``),
			},
		},
		{
			name: "negative test #2 - lost body",
			want: want{
				code: http.StatusBadRequest,
				body: []byte(`http://shortener.com`),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			api := mock_shortener.NewMockservice(ctrl)
			td := New(api, "BaseURL")

			body := bytes.NewReader(test.want.body)
			r := httptest.NewRequest(http.MethodPost, "/", body)
			w := httptest.NewRecorder()

			api.EXPECT().
				Create(gomock.Any()).
				Return("", assert.AnError).
				AnyTimes()

			td.createHandler(w, r)

			assert.Equal(t, test.want.code, w.Code, "Code doesnt match")
		})
	}
}
