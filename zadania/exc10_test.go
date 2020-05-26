package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRouter(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	tests := []struct {
		name    string
		surname string
		want    PersonalInformation
	}{
		{name: "admin", surname: "adnim1", want: PersonalInformation{
			Name:    "admin",
			Surname: "admin1",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, server := Router()

			vars := map[string]string{
				"name":    tt.name,
				"surname": tt.surname,
			}

			req, _ := http.NewRequest("GET", "/path/"+tt.name+"/"+tt.surname+"", nil)
			req = mux.SetURLVars(req, vars)

			server.NameHandler(httptest.NewRecorder(), req)

			userList := server.usersList
			got := userList[0]
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiHandler(t *testing.T) {

	tests := []struct {
		name string
		key  string
		want string
	}{
		{name: "basic test", key: "123", want: "my key is 123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, server := Router()

			vars := map[string]string{
				"key": tt.key,
			}

			req, _ := http.NewRequest("GET", "/api/"+tt.key+"", nil)
			req = mux.SetURLVars(req, vars)

			got := httptest.NewRecorder()

			server.ApiHandler(got, req)

			if !reflect.DeepEqual(got.Body.String(), tt.want) {
				t.Errorf("ApiHandler() got = '%v', want '%v'", got.Body.String(), tt.want)
			}
		})
	}
}
