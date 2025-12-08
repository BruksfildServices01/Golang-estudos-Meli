package main

import (
	"log"
	"net/http"

	_ "api-campeonato/docs"
	"api-campeonato/handler"
	"api-campeonato/service"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func main() {
	// TORNEIOS
	torneioService := service.NewTorneioService()
	torneioHandler := handler.NewTorneioHandler(torneioService)
	http.Handle("/torneios", torneioHandler)
	http.Handle("/torneios/", torneioHandler)

	// TIMES
	timeService := service.NewTimeService()
	timeHandler := handler.NewTimeHandler(timeService)
	http.Handle("/times", timeHandler)
	http.Handle("/times/", timeHandler)

	// JOGADORES
	jogadorService := service.NewJogadorService()
	jogadorHandler := handler.NewJogadorHandler(jogadorService)
	http.Handle("/jogadores", jogadorHandler)
	http.Handle("/jogadores/", jogadorHandler)

	// SWAGGER UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Servidor rodando em http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
