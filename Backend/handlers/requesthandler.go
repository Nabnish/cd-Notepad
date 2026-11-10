package handlers

import(
	"encoding/json"
	"log"
	"net/http"

	"github.com/cd-Notepad/Backend/services"
)



func createRequest(w http.ResponseWriter, r *http.Request) {
	var req services.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Fatal(err)
	}
	err = req.InsertRequest(req)
	if err != nil {
		resp := Response{
			Msg:  "ERROR",
			Code: 304,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	res := Response{
		Msg:  "succesfully created a request",
		Code: 200,
	}

	jsonResponse, err := json.Marshal(res)
	jsonResponse_req, e := json.Marshal(req)
	if err != nil || e != nil {
		log.Fatal(err)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jsonResponse)
	w.Write(jsonResponse_req)
}

func GetAllRequest(w http.ResponseWriter,r *http.Request){

}

func GetRequestID(w http.ResponseWriter,r *http.Request){

}

func DeleteRequest(w http.ResponseWriter, r *http.Request){
	
}
