package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

// RespondJson - Respnds with json resp header as part of HTTP Response
func RespondJson(w http.ResponseWriter, JsonType interface{}) {
	js, err := json.Marshal(JsonType)
	if err != nil {
		RespondError(w, err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// RespondError - HTTP Response with error
func RespondError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

// GetIOStream - Read entire IO String in the format of JSON
func GetIOStream(body io.ReadCloser) (interface{}, error) {
	contents, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	jsonStr, err := GetJson(contents)
	return jsonStr, err
}

// GetJson - Converts byte stream into JSON Format
func GetJson(body []byte) (jsonSource interface{}, err error) {
	if string(body) != "" && body != nil {
		err := json.Unmarshal(body, &jsonSource)
		if err != nil {
			return nil, err
		}
	}
	return jsonSource, err
}

// WriteJson - Converts JSON struct into byte stream
func WriteJson(JsonType interface{}) ([]byte, error) {
	js, err := json.Marshal(JsonType)
	if err != nil {
		return nil, err
	}
	return js, err
}

// WriteJsonToFile - Write JSON formatted string to file
func WriteJsonToFile(JSONType interface{}, file string) error {
	jsonData, err := WriteJson(JSONType)
	if err != nil {
		return err
	}
	return WriteBytesToFile(jsonData, file)
}

// ReturnPrettyPrintJson - Returns byte stream into pretty json
func ReturnPrettyPrintJson(body []byte) string {
	jsonStr, err := GetJson(body)
	output, err := json.MarshalIndent(jsonStr, "", "   ")
	if err != nil {
		return err.Error()
	}
	return string(output)
}

// CheckError - Prints the error details
func CheckError(err error) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		log.Fatalf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	}
}

// ConvertToInt -
func ConvertToInt(dataType interface{}) int {
	reflectType := reflect.TypeOf(dataType)
	var returnType int
	if reflectType.Kind() == reflect.Int {
		returnType = dataType.(int)
	} else if reflectType.Kind() == reflect.Float64 {
		returnType = int(dataType.(float64))
	}
	return returnType
}
