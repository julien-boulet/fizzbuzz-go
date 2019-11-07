package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/dto"
	"github.com/jboulet/fizzbuzz-go/service"
	"log"
	"net/http"
)

var decoder = schema.NewDecoder()

type App struct {
	Router *mux.Router
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Queries("int1", "{int1:[0-9]+}", "int2", "{int2:[0-9]+}", "limit", "{limit:[0-9]+}", "str1", "{str1:[A-Za-z]+}", "str2", "{str2:[A-Za-z]+}").
		Path("/fizzbuzz").
		HandlerFunc(app.getFunction)
}

func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {

	var gameParameter dto.GameParamater

	err := decoder.Decode(&gameParameter, r.URL.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	} else {
		log.Println("GET parameters : ", gameParameter)
	}

	result := make([]string, gameParameter.Limit)
	i := 0
	for out := range service.FizzBuzz(gameParameter) {
		result[i] = out
		i += 1
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
