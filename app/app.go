package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/dto"
	"github.com/jboulet/fizzbuzz-go/service"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"net/http"
)

type App struct {
	Router   *mux.Router
	Decoder  *schema.Decoder
	Database *sql.DB
	Producer *kafka.Producer
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Queries("int1", "{int1:[0-9]+}", "int2", "{int2:[0-9]+}", "limit", "{limit:[0-9]+}", "str1", "{str1:[A-Za-z]+}", "str2", "{str2:[A-Za-z]+}").
		Path("/fizzbuzz").
		HandlerFunc(app.playFizzBuzz)

	app.Router.
		Methods("GET").
		Path("/oneTopStatistic").
		HandlerFunc(app.oneTopStatistic)
}

func (app *App) playFizzBuzz(w http.ResponseWriter, r *http.Request) {

	gameParameter := dto.GameParameter{}
	if err := app.Decoder.Decode(&gameParameter, r.URL.Query()); err != nil {
		log.Fatal("Error in GET parameters : ", err)
	}

	for value := range service.FizzBuzz(&gameParameter) {
		fmt.Fprintln(w, value)
	}
	service.Save(app.Database, &gameParameter)
	service.PushtoKafka(app.Producer, &gameParameter)
}

func (app *App) oneTopStatistic(w http.ResponseWriter, r *http.Request) {
	result := service.FindMax(app.Database)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
