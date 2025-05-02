package transport_test

import (
	"learn/rest-api/internal/book/service"
	"learn/rest-api/internal/book/transport"
	"learn/rest-api/internal/errdefs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// Мок для теста
type mockService struct{}

func (m *mockService) GetBook(input service.BookInputDTO) (service.BookOutputDTO, error) {
	if input.ID == 9 {
		return service.BookOutputDTO{}, errdefs.ErrNotFound
	}
	return service.BookOutputDTO{ID: input.ID, Name: "Test Book", Author: "Test Author"}, nil
}

// TODO доделать
func (m *mockService) CreateBook(service.CreateBookInputDTO) (service.CreateBookOutputDTO, error) {
	return service.CreateBookOutputDTO{}, nil
}

func TestGetBookHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := transport.NewBookHandler(&mockService{})
	router.GET("/books/:id", handler.GetBook)

	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "valid id",
			url:          "/books/1",
			expectedCode: 200,
			expectedBody: `{"id":1,"name":"Test Book","author":"Test Author"}`,
		},
		{
			name:         "invalid id (negative)",
			url:          "/books/-1",
			expectedCode: 400,
			expectedBody: `{"error":"Key: 'BookRequestDTO.ID' Error:Field validation for 'ID' failed on the 'min' tag"}`,
		},
		{
			name:         "not found id",
			url:          "/books/9",
			expectedCode: 404,
			expectedBody: `{"error":"resource not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status %d, got %d", tt.expectedCode, w.Code)
			}
			body := strings.TrimSpace(w.Body.String())
			if body != tt.expectedBody {
				t.Errorf("unexpected body: got %s, want %s", body, tt.expectedBody)
			}
		})
	}
}
