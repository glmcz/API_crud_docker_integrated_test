package serviceLayer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"simpleCloudService/internal/model"
	"simpleCloudService/internal/repository"
	utils "simpleCloudService/pkg/http"
	"strings"
)

type ServerConfig struct {
	Address string `yaml:"address"`
}

type ServiceLayer struct {
	DBConnection interface{}
	Template     *template.Template
}

func NewServiceLayer(dbConnection interface{}, templatePath string) *ServiceLayer {
	files, err := filepath.Glob(templatePath + "/*.html")
	if err != nil {
		return nil
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		return nil
	}

	return &ServiceLayer{
		DBConnection: dbConnection,
		Template:     t,
	}
}

// Muxer
// Muxer route each request to handlers, where each handler is started as a new go routine, which
// gives us parallel handling of incoming requests natively.
func (a *ServiceLayer) Muxer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/save", a.postUserRequest)
	mux.HandleFunc("/user", a.getUserRequest)
	mux.HandleFunc("/render", a.postRenderRequest)
	mux.HandleFunc("/", a.mainPage)
	return mux
}

// mainPage
// render main page
func (a *ServiceLayer) mainPage(w http.ResponseWriter, _ *http.Request) {
	err := a.Template.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		response := utils.HttpResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}
		response.WriteResponse(w)
		return
	}
}

// postRenderRequest
// Proceed with user form request and display results.
func (a *ServiceLayer) postRenderRequest(w http.ResponseWriter, r *http.Request) {

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		response := utils.HttpResponse{
			Code: http.StatusBadRequest,
			Msg:  "Error parsing form: " + err.Error(),
		}
		response.WriteResponse(w)
		return
	}

	// Get the JSON input from the form
	jsonInput := r.Form.Get("json_input")
	if isEmpty(jsonInput) {
		response := utils.HttpResponse{
			Code: http.StatusBadRequest,
			Msg:  "No JSON input provided",
		}
		response.WriteResponse(w)
		return
	}
	var validThread model.Thread

	res, err := JsonParsing(jsonInput, validThread)
	if err != nil {
		response := utils.HttpResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}
		response.WriteResponse(w)
		return
	}

	err = a.Template.ExecuteTemplate(w, "thread.html", res)
	if err != nil {
		response := utils.HttpResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}
		response.WriteResponse(w)
		return
	}
}

func JsonParsing(jsonData string, structure model.Thread) (interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(jsonData))
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&structure)
	if err != nil {

		var unmarshalTypeError *json.UnmarshalTypeError
		var syntaxError *json.SyntaxError

		switch {
		case errors.As(err, &unmarshalTypeError):
			return structure, fmt.Errorf("type error at %v: %v", unmarshalTypeError.Field, err)
		case errors.As(err, &syntaxError):
			return structure, fmt.Errorf("json syntax error: %v", err)
		case strings.Contains(err.Error(), "json: unknown field"):
			//field := extractFieldName(err.Error())
			return structure, err
		default:
			return structure, err
		}
	}

	return structure, nil
}

func isEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

/********** End of Eset task ***********/
// postUserRequest
// endpoint for creating a new users i DB
func (a *ServiceLayer) postUserRequest(w http.ResponseWriter, r *http.Request) {
	println("postUserRequest")
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

// getUserRequest
// Return user info from DB.
func (a *ServiceLayer) getUserRequest(w http.ResponseWriter, r *http.Request) {
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

//func extractFieldName(errorMsg string) string {
//	startIndex := strings.Index(errorMsg, "\"")
//	endIndex := strings.LastIndex(errorMsg, "\"")
//
//	if startIndex != -1 && endIndex != -1 && startIndex != endIndex {
//		return errorMsg[startIndex+1 : endIndex]
//	}
//
//	return ""
//}
