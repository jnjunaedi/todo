package main

import (
	"log"
	"net/http"
	"os"
	"todo-app/config"
	taskHd "todo-app/internal/delivery/http/handler"
	taskRepo "todo-app/internal/repository/mysql"
	taskUc "todo-app/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan environment system")
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal inisialisasi database: %v", err)
	}
	defer db.Close()
	log.Println("Berhasil terhubung ke database MySQL!")

	taskRepo := taskRepo.NewRepository(db)
	taskUsecase := taskUc.NewUsecase(taskRepo)
	taskHandler := taskHd.NewHandler(taskUsecase)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/tasks", taskHandler.HandlerCreateTask)

	clientOrigin := os.Getenv("CLIENT_ORIGIN")

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{clientOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handlerWithCORS := corsMiddleware.Handler(mux)

	port := os.Getenv("PORT")

	log.Printf("Server berjalan di port :%s\n", port)
	if err := http.ListenAndServe(":"+port, handlerWithCORS); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
