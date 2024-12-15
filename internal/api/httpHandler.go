package api

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"simpleCloudService/internal/model"
	"simpleCloudService/internal/repository"
	utils "simpleCloudService/pkg/http"
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

func (a *API) postRequests(w http.ResponseWriter, r *http.Request) {
	println("postRequests")
	var user model.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response := utils.HttpResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}
		response.WriteResponse(w)
		return
	}
	log.Printf("Received body: %s", string(body))

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&user); err != nil {
		response := utils.HttpResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}
		response.WriteResponse(w)
	}

	err = a.DBConnection.(*repository.PostgresRepository).CreateUser(&user)
	if err != nil {
		response := utils.HttpResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}
		response.WriteResponse(w)
		return
	}
}

func (a *API) getRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rawQuery := r.URL.Path
		userID := rawQuery[1:]
		user, err := a.DBConnection.(*repository.PostgresRepository).GetUser(uuid.MustParse(userID))
		if err != nil {
			response := utils.HttpResponse{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			}
			response.WriteResponse(w)
		}

		response := utils.HttpResponse{
			Code: http.StatusOK,
			Msg:  user.ToString(),
		}
		response.WriteResponse(w)

	} else {
		response := utils.HttpResponse{
			Code: http.StatusBadRequest,
			Msg:  "Try different http method",
		}
		response.WriteResponse(w)
	}
}
