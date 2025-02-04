package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strconv"
    "testing"
    "backend/internal/handler"
    "backend/internal/task"
    "backend/internal/repository"
    "backend/pkg/database"
    "backend/config"
    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func setupTestEnvironment() (*mux.Router, repository.TaskRepository) {
    cfg := config.LoadConfig()
    db, err := database.NewDB(cfg)
    if err != nil {
        panic(err)
    }

    taskRepo := repository.NewTaskRepository(db)
    taskService := task.NewService(taskRepo)
    taskHandler := handler.NewTaskHandler(taskService)

    r := mux.NewRouter()
    r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
    r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
    r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")

    return r, taskRepo
}

func TestCreateTask(t *testing.T) {
    router, repo := setupTestEnvironment()

    newTask := task.Task{
        Title:       "Test Task",
        Description: "This is a test task",
        Completed:   false,
    }

    payload, err := json.Marshal(newTask)
    assert.NoError(t, err)

    req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(payload))
    assert.NoError(t, err)

    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)

    var createdTask task.Task
    err = json.Unmarshal(rr.Body.Bytes(), &createdTask)
    assert.NoError(t, err)
    assert.NotZero(t, createdTask.ID)
    assert.Equal(t, newTask.Title, createdTask.Title)
    assert.Equal(t, newTask.Description, createdTask.Description)
    assert.Equal(t, newTask.Completed, createdTask.Completed)


    dbTask, err := repo.FindByID(createdTask.ID)
    assert.NoError(t, err)
    assert.Equal(t, createdTask.ID, dbTask.ID)
    assert.Equal(t, createdTask.Title, dbTask.Title)
    assert.Equal(t, createdTask.Description, dbTask.Description)
    assert.Equal(t, createdTask.Completed, dbTask.Completed)
}

func TestGetTask(t *testing.T) {
    router, repo := setupTestEnvironment()

    newTask := task.Task{
        Title:       "Test Task",
        Description: "This is a test task",
        Completed:   false,
    }
    err := repo.Create(&newTask)
    assert.NoError(t, err)

    req, err := http.NewRequest("GET", "/tasks/"+strconv.Itoa(newTask.ID), nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var fetchedTask task.Task
    err = json.Unmarshal(rr.Body.Bytes(), &fetchedTask)
    assert.NoError(t, err)
    assert.Equal(t, newTask.ID, fetchedTask.ID)
    assert.Equal(t, newTask.Title, fetchedTask.Title)
    assert.Equal(t, newTask.Description, fetchedTask.Description)
    assert.Equal(t, newTask.Completed, fetchedTask.Completed)
}