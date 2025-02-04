package handler

import (
    "backend/internal/task"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type TaskHandler struct {
    service *task.Service
}

func NewTaskHandler(service *task.Service) *TaskHandler {
    return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var t task.Task
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
        return
    }

    if err := h.service.CreateTask(&t); err != nil {
        http.Error(w, "Erro ao criar a tarefa", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    t, err := h.service.GetTaskByID(id)
    if err != nil {
        http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(t)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    var t task.Task
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
        return
    }

    t.ID = id
    if err := h.service.UpdateTask(&t); err != nil {
        http.Error(w, "Erro ao atualizar a tarefa", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    if err := h.service.DeleteTask(id); err != nil {
        http.Error(w, "Erro ao deletar a tarefa", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.service.GetAllTasks()
    if err != nil {
        http.Error(w, "Erro ao buscar tarefas", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(tasks)
}