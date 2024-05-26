package main

import (
	"brainy-framework/brainy/internal/container"
	"brainy-framework/brainy/internal/middleware"
	"brainy-framework/brainy/internal/service"
	"brainy-framework/brainy/internal/transaction"
	"brainy-framework/brainy/pkg/framework"
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fw := framework.NewFramework()

	container := container.NewContainer()
	db, _ := sql.Open("sqlite3", ":memory:")
	container.Register(db)

	tm := transaction.NewTransactionManager(db)
	container.Register(tm)

	exampleService := &service.ExampleService{}
	fw.RegisterComponent("ExampleService", exampleService)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	authMiddleware := middleware.NewAuthMiddleware(mux)

	loggingMiddleware := middleware.NewLoggingMiddleware(authMiddleware)

	if err := fw.Start(); err != nil {
		panic(err)
	}
	defer fw.Stop()

	http.Handle("/", authMiddleware)
	http.ListenAndServe(":8080", loggingMiddleware)
}
