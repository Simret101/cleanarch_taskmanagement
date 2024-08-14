package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"task/Delivery/controllers"
	"task/Domain"
	"task/Usecases"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) Register(user *Domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) Login(credentials *Domain.Credentials) (string, error) {
	args := m.Called(credentials)
	return args.String(0), args.Error(1)
}

func TestUserController(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		body         string
		expectedCode int
		expectedBody string
		mockSetup    func(mockUseCase *MockUserUseCase)
	}{
		{
			name:         "Register",
			method:       http.MethodPost,
			url:          "/register",
			body:         `{"username": "testuser", "password": "password", "email": "test@example.com"}`,
			expectedCode: http.StatusCreated,
			expectedBody: "",
			mockSetup: func(mockUseCase *MockUserUseCase) {
				mockUseCase.On("Register", mock.AnythingOfType("*Domain.User")).Return(nil)
			},
		},
		{
			name:         "Login",
			method:       http.MethodPost,
			url:          "/login",
			body:         `{"username": "testuser", "password": "password"}`,
			expectedCode: http.StatusOK,
			expectedBody: `{"token":"mock-token"}`,
			mockSetup: func(mockUseCase *MockUserUseCase) {
				mockUseCase.On("Login", mock.AnythingOfType("*Domain.Credentials")).Return("mock-token", nil)
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			mockUseCase := new(MockUserUseCase)
			tt.mockSetup(mockUseCase)

			controller := controllers.UserController{UserUseCase: Usecases.UserUseCase{}}

			gin.SetMode(gin.TestMode)
			r := gin.Default()

			r.Handle(tt.method, tt.url, controller.Register)
			r.Handle(tt.method, tt.url, controller.Register)

			req, _ := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}
