package handlers

import (
	"encoding/json"
	"log"
	"net/http"

)

type Response struct {
	Msg  string
	Code int
}

func Health(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Msg:  "health",
		Code: 200,
	}
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonResponse)
}

