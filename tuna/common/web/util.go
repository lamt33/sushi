package web

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadBody(requestBody io.ReadCloser, pointerThing interface{}) error {
	b, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return fmt.Errorf("could not ready body (ReadAll): %s", err.Error())
	}

	err = json.Unmarshal(b, pointerThing)
	if err != nil {
		return fmt.Errorf("could not ready body (Unmarshal): %s", err.Error())
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	res, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("cannot marshal response: %v", err)
	}
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Printf("cannot generate error message : %s", writeErr)
	}
	return nil
}

type ErrResponse struct {
	Error string `json:"error"`
}

func GenerateError(w http.ResponseWriter, errorMessage string, code int) {
	// log_plum.Error(fmt.Sprintf("%d : %s", code, errorMessage))
	err := WriteJSON(w, code, ErrResponse{errorMessage})
	if err != nil {
		log.Printf("could not write json message: %v", err)
	}
}

type OkResponse struct {
	Message string `json:"message"`
}

func GenerateSuccess(w http.ResponseWriter, successMessage string) {
	err := WriteJSON(w, http.StatusOK, &OkResponse{successMessage})
	if err != nil {
		log.Printf("could not write json message: %v", err)
	}
}
