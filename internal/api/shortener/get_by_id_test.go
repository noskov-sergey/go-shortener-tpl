package shortener

import (
	"testing"
)

func TestImplementation_GetByID_Success(t *testing.T) {
	//type want struct {
	//	code int
	//}
	//tests := []struct {
	//	name string
	//	want want
	//}{
	//	{
	//		name: "positive test #1 - handler work good",
	//		want: want{
	//			code: http.StatusTemporaryRedirect,
	//		},
	//	},
	//}
	//
	//for _, test := range tests {
	//	t.Run(test.name, func(t *testing.T) {
	//		ctrl := gomock.NewController(t)
	//		api := mock_shortener.NewMockservice(ctrl)
	//		td := New(api, "BaseURL")
	//		server := httptest.NewServer(td)
	//
	//		r := httptest.NewRequest(http.MethodGet, "/segseggs", nil)
	//
	//		api.EXPECT().
	//			GetByID(gomock.Any()).
	//			Return("http://shortener.com", nil)
	//
	//		resp, err := server.Client().Do(r)
	//		if err != nil {
	//			return
	//		}
	//
	//		assert.Equal(t, test.want.code, resp.Request.Response.StatusCode, "Code doesnt match")
	//	})
	//}
}

func TestImplementation_GetByID_Error(t *testing.T) {
	//type want struct {
	//	code int
	//	body []byte
	//}
	//tests := []struct {
	//	name string
	//	want want
	//}{
	//	{
	//		name: "negative test #1 - error work",
	//		want: want{
	//			code: http.StatusBadRequest,
	//		},
	//	},
	//	{
	//		name: "negative test #2 - lost body",
	//		want: want{
	//			code: http.StatusBadRequest,
	//		},
	//	},
	//}
	//
	//for _, test := range tests {
	//	t.Run(test.name, func(t *testing.T) {
	//		ctrl := gomock.NewController(t)
	//		api := mock_shortener.NewMockservice(ctrl)
	//		td := New(api, "BaseURL")
	//		server := httptest.NewServer(td)
	//
	//		r := httptest.NewRequest(http.MethodGet, "/", nil)
	//		w := httptest.NewRecorder()
	//
	//		api.EXPECT().
	//			GetByID(gomock.Any()).
	//			Return("http://shortener.com", nil).
	//			AnyTimes()
	//
	//		td.getByIDHandler(w, r)
	//
	//		resp, err := server.Client().Do(r)
	//		if err != nil {
	//			return
	//		}
	//
	//		assert.Equal(t, test.want.code, resp.StatusCode, "Code doesnt match")
	//	})
	//}
}
