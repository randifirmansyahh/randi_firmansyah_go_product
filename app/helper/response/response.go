package response

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/models/responseModel"
)

func Response(w http.ResponseWriter, code int, msg string, data interface{}) {
	receiver := responseModel.Response{}

	receiver.Status = cekStatus(code)
	receiver.Code = code
	receiver.Message = msg
	receiver.Data = data

	jadiJson, err := json.Marshal(receiver) // nge convert jadi json

	w.Header().Set("Content-Type", "application/json") // return type data nya

	if err != nil {
		w.WriteHeader(500) // status code
		w.Write([]byte("Error to marshall"))
	}

	// success
	w.WriteHeader(code) // status code
	w.Write(jadiJson)   // return datanya
}

func ResponseInternalServerError(w http.ResponseWriter) {
	Response(w, http.StatusInternalServerError, internal_server_error, nil)
}

func ResponseSuccess(w http.ResponseWriter, crud, handlerName string, data interface{}) {
	Response(w, http.StatusOK, cekForMassage(crud, handlerName), data)
}

func ResponseRunningService(w http.ResponseWriter) {
	Response(w, http.StatusOK, service_running, nil)
}

func ResponseBadRequest(w http.ResponseWriter) {
	Response(w, http.StatusBadRequest, bad_request, nil)
}

func ResponNotFound(w http.ResponseWriter) {
	Response(w, http.StatusNotFound, not_found, nil)
}

func ResponUnAuthorize(w http.ResponseWriter) {
	Response(w, http.StatusUnauthorized, unauthorize, nil)
}

func ResponTokenSucces(w http.ResponseWriter, token interface{}) {
	Response(w, http.StatusOK, tokencreated, token)
}
