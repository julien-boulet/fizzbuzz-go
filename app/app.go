package app

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/dto"
	"github.com/jboulet/fizzbuzz-go/service"
	"log"
	"net/http"
)

type App struct {
	Router   *mux.Router
	Decoder  *schema.Decoder
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Queries("int1", "{int1:[0-9]+}", "int2", "{int2:[0-9]+}", "limit", "{limit:[0-9]+}", "str1", "{str1:[A-Za-z]+}", "str2", "{str2:[A-Za-z]+}").
		Path("/fizzbuzz").
		HandlerFunc(app.getFunction)
}

func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {

	var gameParameter dto.GameParameter

	err := app.Decoder.Decode(&gameParameter, r.URL.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	} else {
		log.Println("GET parameters : ", gameParameter)
	}

	result := service.FizzBuzz(gameParameter)
	service.Save(app.Database, gameParameter)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
