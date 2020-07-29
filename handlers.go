package main

import (
	"io/ioutil"
	"net/http"
)
import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func handleDataflowUpdate(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
		writer.WriteHeader(401)
		return
	}

	d, _:= ioutil.ReadAll(request.Body)
	DataflowLogger.Printf("%v", string(d))
	//var dataflowUpdate models.DataFlowUpdate
	//err := json.NewDecoder(request.Body).Decode(&dataflowUpdate)
	//
	//if err != nil {
	//	ErrorLogger.Printf("Error in decoding dataflow update; err=%v \n", err)
	//	writer.WriteHeader(400)
	//	return
	//}
	//
	//DataflowLogger.Printf("%+v", dataflowUpdate)
	writer.WriteHeader(200)
}

func handleConsentUpdate(writer http.ResponseWriter, request *http.Request) {


	if request.Method != http.MethodPost {
		writer.WriteHeader(401)
		return
	}

	c, _:= ioutil.ReadAll(request.Body)
	ConsentLogger.Printf("%v", string(c))
	//var consentUpdate models.ConsentUpdate
	//err := json.NewDecoder(request.Body).Decode(&consentUpdate)
	//
	//if err != nil {
	//	ErrorLogger.Printf("Error in decoding consent update; err=%v \n", err)
	//	writer.WriteHeader(400)
	//	return
	//}
	//
	//ConsentLogger.Printf("%+v", consentUpdate)
	writer.WriteHeader(200)

}
