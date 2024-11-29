package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"simpleCloudService/internal/model"
	"simpleCloudService/internal/repository"
)

type ServerConfig struct {
	Address string `yaml:"address"`
}

type API struct {
	DBConnection interface{}
}

func NewAPI(dbConnection interface{}) *API {
	return &API{
		DBConnection: dbConnection,
	}
}

func (a *API) Muxer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/save", a.postRequests)
	mux.HandleFunc("/", a.getRequests)
	return mux
}

func WriteResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if response == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		res, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a *API) postRequests(w http.ResponseWriter, r *http.Request) {
	println("postRequests")
	var user model.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		WriteResponse(w, nil)
		return
	}
	log.Printf("Received body: %s", string(body))

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&user); err != nil {
		fmt.Printf(err.Error())
		WriteResponse(w, nil)
	}

	err = a.DBConnection.(*repository.PostgresRepository).CreateUser(&user)
	if err != nil {
		return
	}
}

func (a *API) getRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rawQuery := r.URL.Path
		userID := rawQuery[1:]
		user, err := a.DBConnection.(*repository.PostgresRepository).GetUser(uuid.MustParse(userID))
		if err != nil {
			fmt.Printf(err.Error())
			WriteResponse(w, nil)
		}
		WriteResponse(w, user)
	} else {
		WriteResponse(w, nil)
	}
}
