package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
)

type Comment struct {
    Message string `validate:"required,min=1,max=140"`
    UserName string `validate:"required,min=1,max=15"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{Name: "John Doe", Email: "john@example.com"},
	{Name: "Jane Doe", Email: "jane@example.com"},
	// Add more users as needed
}

func main(){

    var mutex = &sync.RWMutex{}
    comments := make([]Comment,0,100)

    http.HandleFunc("/comments",func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type","application/json")

        switch r.Method{
            
        case http.MethodGet:
            mutex.RLock()

            if err := json.NewEncoder(w).Encode(comments); err != nil {
                http.Error(w,fmt.Sprintf(`{"status":"%s"}`,err),http.StatusInternalServerError)
                return
            }
            mutex.RUnlock()

        case http.MethodPost:
            var c Comment
            if err := json.NewDecoder(r.Body).Decode(&c); err != nil{
                http.Error(w,fmt.Sprintf(`{"status":"%s"}`,err),http.StatusInternalServerError)
                return
            }
            validate := validator.New()
            if err:=validate.Struct(c); err != nil{
                http.Error(w,fmt.Sprintf(`{"status":"%s"}`,err),http.StatusBadRequest)
                return
            }
            mutex.Lock()
            comments = append(comments, c)
            mutex.Unlock()

            w.WriteHeader(http.StatusCreated)
            w.Write([]byte(`{"status":"created"}`))

        default:
            http.Error(w,`{"status":"permits only GET or POST"}`, http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			// Return the list of users as JSON
			if err := json.NewEncoder(w).Encode(users); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, `{"status":"permits only GET"}`, http.StatusMethodNotAllowed)
		}
	})
    http.ListenAndServe(":8888",nil)
}

