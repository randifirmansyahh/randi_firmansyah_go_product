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

	w.WriteHeader(code) // status code
	w.Write(jadiJson)   // return datanya
}

func ResponseInternalServerError(w http.ResponseWriter) {
	Response(w, http.StatusInternalServerError, MsgServiceErr(), nil)
}

func ResponseSuccess(w http.ResponseWriter, crud, handlerName string, data interface{}) {
	Response(w, http.StatusOK, cekForMassage(crud, handlerName), data)
}

func ResponseBadRequest(w http.ResponseWriter) {
	Response(w, http.StatusBadRequest, MsgInvalidReq(), nil)
}

func cekStatus(code int) string {
	if code == http.StatusOK {
		return "success"
	} else if code == http.StatusBadRequest {
		return "bad request"
	} else if code == http.StatusInternalServerError {
		return "internal server error"
	} else {
		return "unknown"
	}
}

func cekForMassage(c, h string) string {
	if c == "c" {
		return MsgTambah(h)
	} else if c == "r" {
		return MsgGetData(h)
	} else if c == "u" {
		return MsgUpdate(h)
	} else if c == "d" {
		return MsgHapus(h)
	} else {
		return "Berhasil"
	}
}

func MsgTambah(model string) string {
	return "Berhasil menambahkan data " + model
}

func MsgGetAll(model string) string {
	return "Berhasil mengambil semua data " + model
}

func MsgGetData(model string) string {
	return "Berhasil mengambil data " + model
}

func MsgGetDetail(model string) string {
	return "Berhasil mengambil detail data " + model
}

func MsgUpdate(model string) string {
	return "Berhasil mengupdate data " + model
}

func MsgHapus(model string) string {
	return "Berhasil menghapus data " + model
}

func MsgNotFound(model string) string {
	return "Data " + model + " tidak ditemukan"
}

func MsgServiceErr() string {
	return "Terdapat masalah pada service"
}

func MsgInvalidReq() string {
	return "Invalid Request"
}
