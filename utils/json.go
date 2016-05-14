package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"go-http-api/client"
)

func RespondJson(w http.ResponseWriter, JsonType interface{}) {
	js, err := json.Marshal(JsonType)
	if err != nil {
		client.RespondError(w, err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetJson(body []byte) (jsonSource interface{}) {
	if string(body) != "" && body != nil {
		err := json.Unmarshal(body, &jsonSource)
		if err != nil {
			log.Print("template executing error: ", err)
		}
	}
	return
}

func WriteJson(JsonType interface{}) []byte {
	js, err := json.Marshal(JsonType)
	if err != nil {
		log.Print("Invalid JSON format", err)
	}
	return js
}

func WriteJsonToFile(JsonType interface{}, file string) {
	 WriteBytesToFile(WriteJson(JsonType), file)
}

func ReturnPrettyPrintJson(body []byte) string {
	 output, err := json.MarshalIndent(GetJson(body), "", "   ")
	 if err != nil {
	         log.Print("Json Intending Error: ", err)
	 }
	 return string(output)
}
