package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	redis2 "github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	_ "github.com/jboulet/fizzbuzz-go/docs"
	"github.com/jboulet/fizzbuzz-go/dto"
	"github.com/jboulet/fizzbuzz-go/service"
	kafka "github.com/segmentio/kafka-go"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type App struct {
	Router   *mux.Router
	Decoder  *schema.Decoder
	Database *sql.DB
	Producer *kafka.Writer
	Redis    *redis2.Client
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

	app.Router.
		Methods("GET").
		Path("/docs/swagger.json").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "docs/swagger.json")
		})

	// Swagger
	app.Router.
		PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

// @Summary play fizzbuzz game
// @Description play fizzbuzz game with specifics params
// @ID fizz-buzz-game
// @Accept  json
// @Produce  json
// @Param int1 query int true "first int for game"
// @Param int2 query int true "second for game"
// @Param limit query int true "limit of calculation"
// @Param str1 query string true "first word for game"
// @Param str2  query string true "second word for game"
// @Success 200 {string} string "One calculation"
// @Router /fizzbuzz [get]
func (app *App) playFizzBuzz(w http.ResponseWriter, r *http.Request) {

	gameParameter := dto.GameParameter{}
	if err := app.Decoder.Decode(&gameParameter, r.URL.Query()); err != nil {
		log.Fatal("Error in GET parameters : ", err)
	}

	for value := range service.FizzBuzz(&gameParameter) {
		fmt.Fprintln(w, value)
	}
	service.Save(app.Database, &gameParameter)
	service.PushtoKafka(app.Producer, &gameParameter, r)
	service.PushtoRedis(app.Redis, &gameParameter, r)
}

// @Summary ask the best statistics
// @Description return the params most used to play
// @ID fizz-buzz-statistic
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.StatisticResult
// @Router /oneTopStatistic [get]
func (app *App) oneTopStatistic(w http.ResponseWriter, r *http.Request) {
	result := service.FindMax(app.Database)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
