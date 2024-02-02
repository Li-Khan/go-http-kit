package go_http_kit

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpKit_Methods(t *testing.T) {
	hk := New()

	hk.GET("/get/hello", helloHandle)
	hk.HEAD("/head/hello", helloHandle)
	hk.POST("/post/hello", helloHandle)
	hk.PUT("/put/hello", helloHandle)
	hk.PATCH("/patch/hello", helloHandle)
	hk.DELETE("/delete/hello", helloHandle)
	hk.CONNECT("/connect/hello", helloHandle)
	hk.OPTIONS("/options/hello", helloHandle)
	hk.TRACE("/trace/hello", helloHandle)

	paths := map[string]string{
		"/get/hello":     http.MethodGet,
		"/head/hello":    http.MethodHead,
		"/post/hello":    http.MethodPost,
		"/put/hello":     http.MethodPut,
		"/patch/hello":   http.MethodPatch,
		"/delete/hello":  http.MethodDelete,
		"/connect/hello": http.MethodConnect,
		"/options/hello": http.MethodOptions,
		"/trace/hello":   http.MethodTrace,
	}

	server := httptest.NewServer(hk.Mux())
	defer server.Close()

	for path, method := range paths {
		request, err := http.NewRequest(method, server.URL+path, nil)
		if err != nil {
			t.Fatal(err)
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			t.Fatal(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}
		_ = response.Body.Close()
	}
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, testing!"))
}
