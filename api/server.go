package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kamalesh889/spendingCalculator/config"
	"github.com/kamalesh889/spendingCalculator/model"
)

type server struct {
	db         *model.Database
	httpClient *http.Client
	config     *config.Config
	router     *mux.Router
}

func NewServer(cfg *config.Config) (*server, error) {

	var err error

	s := &server{}

	s.config = cfg

	s.db, err = model.InitializeDB(s.config)
	if err != nil {
		log.Println("Error in creating connection to db", err)
		return nil, err
	}
	// configure http client for global usage
	s.httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	s.router = mux.NewRouter()

	return s, nil

}
