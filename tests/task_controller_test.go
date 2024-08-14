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

// Mock for TaskUseCase
type MockTaskUseCase struct {
	mock.Mock
}
type TaskController struct {
	TaskUseCase *MockTaskUseCase
}

func (m *MockTaskUseCase) GetAllTasks() ([]Domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) CreateTask(task *Domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUseCase) GetTaskByID(id int) (*Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) UpdateTask(id int, task *Domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskUseCase) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestTaskController(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		body         string
		expectedCode int
		mockSetup    func(mockUseCase *MockTaskUseCase)
	}{
		{
			name:         "GetAllTasks",
			method:       http.MethodGet,
			url:          "/tasks",
			body:         "",
			expectedCode: http.StatusOK,
			mockSetup: func(mockUseCase *MockTaskUseCase) {
				mockUseCase.On("GetAllTasks").Return([]Domain.Task{}, nil)
			},
		},
		{
			name:         "CreateTask",
			method:       http.MethodPost,
			url:          "/tasks",
			body:         `{"title": "Test Task", "description": "Test Description", "dueDate": "2023-08-09", "status": "pending"}`,
			expectedCode: http.StatusCreated,
			mockSetup: func(mockUseCase *MockTaskUseCase) {
				mockUseCase.On("CreateTask", mock.AnythingOfType("*Domain.Task")).Return(nil)
			},
		},
		{
			name:         "GetTaskByID",
			method:       http.MethodGet,
			url:          "/tasks/1",
			body:         "",
			expectedCode: http.StatusOK,
			mockSetup: func(mockUseCase *MockTaskUseCase) {
				mockTask := &Domain.Task{
					ID:          1,
					Title:       "Test Task",
					Description: "Test Description",
					DueDate:     "2023-08-09",
					Status:      "pending",
				}
				mockUseCase.On("GetTaskByID", 1).Return(mockTask, nil)
			},
		},
		{
			name:         "UpdateTask",
			method:       http.MethodPut,
			url:          "/tasks/1",
			body:         `{"title": "Updated Task", "description": "Updated Description", "dueDate": "2023-08-10", "status": "completed"}`,
			expectedCode: http.StatusOK,
			mockSetup: func(mockUseCase *MockTaskUseCase) {
				mockTask := &Domain.Task{
					ID:          1,
					Title:       "Updated Task",
					Description: "Updated Description",
					DueDate:     "2023-08-10",
					Status:      "completed",
				}
				mockUseCase.On("UpdateTask", 1, mockTask).Return(nil)
			},
		},
		{
			name:         "DeleteTask",
			method:       http.MethodDelete,
			url:          "/tasks/1",
			body:         "",
			expectedCode: http.StatusOK,
			mockSetup: func(mockUseCase *MockTaskUseCase) {
				mockUseCase.On("DeleteTask", 1).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := new(MockTaskUseCase)
			tt.mockSetup(mockUseCase)

			controller := controllers.TaskController{TaskUseCase: Usecases.TaskUseCase{}}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.Handle(tt.method, tt.url, controller.GetAllTasks)

			req, _ := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			if tt.method == http.MethodPost || tt.method == http.MethodPut {
				req.Header.Set("Content-Type", "application/json")
			}
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			mockUseCase.AssertExpectations(t)
		})
	}
}
