package main

import (
    "backend/config"
    "backend/internal/auth"
    "backend/internal/handler"
    "backend/internal/repository"
    "backend/internal/task"
    "backend/pkg/database"
    "backend/pkg/middleware"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/handlers"

    "github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()


	taskRepo := repository.NewTaskRepository(db)
	taskService := task.NewService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	
	r := mux.NewRouter()

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		
		token, err := auth.GenerateToken("usuario_teste")
		if err != nil {
			http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
			return
		}

		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}).Methods("GET","POST", "OPTIONS") 

	
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware) 

	api.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST", "OPTIONS")
	api.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET", "OPTIONS")
	api.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT", "OPTIONS")
	api.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE", "OPTIONS")
	api.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET", "OPTIONS")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}