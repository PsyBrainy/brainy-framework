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
	nm := framework.NewNeuronManager()

	neuralNetwork := container.NewNeuralNetwork()
	db, _ := sql.Open("sqlite3", ":memory:")
	neuralNetwork.Register(db, framework.Singleton)

	tm := transaction.NewTransactionManager(db)
	neuralNetwork.Register(tm, framework.Singleton)

	exampleService := &service.ExampleService{}
	neuralNetwork.Register(exampleService, framework.Prototype)
	resolvedService := neuralNetwork.Resolve(exampleService).(*service.ExampleService)
	nm.RegisterNeuron("ExampleService", resolvedService)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	authMiddleware := middleware.NewAuthMiddleware(mux)

	loggingMiddleware := middleware.NewLoggingMiddleware(authMiddleware)

	if err := nm.Start(); err != nil {
		panic(err)
	}
	defer nm.Stop()

	http.ListenAndServe(":8080", loggingMiddleware)
}
