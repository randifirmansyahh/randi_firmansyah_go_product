package response

import "net/http"

var (
	success               = "SUCCESS"
	bad_request           = "BAD REQUEST"
	internal_server_error = "INTERNAL SERVER ERROR"
	not_found             = "NOT FOUND"
	unknown               = "UNKNOWN"
	unauthorize           = "UNAUTHORIZE"
	tokencreated          = "TOKEN CREATED"
	service_running       = "SERVICE RUNNING"
)

func cekStatus(code int) string {
	if code == http.StatusOK {
		return success
	} else if code == http.StatusBadRequest {
		return bad_request
	} else if code == http.StatusInternalServerError {
		return internal_server_error
	} else {
		return unknown
	}
}

func cekForMassage(c, h string) string {
	if c == "create" {
		return msgTambah(h)
	} else if c == "read" {
		return msgGetData(h)
	} else if c == "update" {
		return msgUpdate(h)
	} else if c == "delete" {
		return msgHapus(h)
	} else if c == "detail" {
		return msgGetDetail(h)
	} else {
		return "Berhasil"
	}
}

func msgTambah(model string) string {
	return "Berhasil menambahkan data " + model
}

func msgGetData(model string) string {
	return "Berhasil mengambil data " + model
}

func msgGetDetail(model string) string {
	return "Berhasil mengambil detail data " + model
}

func msgUpdate(model string) string {
	return "Berhasil mengupdate data " + model
}

func msgHapus(model string) string {
	return "Berhasil menghapus data " + model
}
