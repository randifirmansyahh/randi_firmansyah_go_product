package loginHandler

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/helper/helper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/helper/tokenHelper"
	"randi_firmansyah/app/models/tokenModel"
	"randi_firmansyah/app/models/userModel"
	"randi_firmansyah/app/service"
	"strconv"
)

var (
	WAKTU        = tokenHelper.WAKTU
	AUD          = tokenHelper.AUD
	ISS          = tokenHelper.ISS
	LOGIN_SECRET = tokenHelper.LOGIN_SECRET
)

type loginHandler struct {
	service service.Service
}

func NewLoginHandler(loginService service.Service) *loginHandler {
	return &loginHandler{loginService}
}

func (l *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	// cek user dan pass
	// decode from json
	decoder := json.NewDecoder(r.Body)

	// fill to model
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Gagal login", nil)
		return
	}

	// select ke db
	cari, err := l.service.IUserService.FindByUsername(datarequest.Username)
	if err != nil {
		response.Response(w, http.StatusOK, "Username tidak ditemukan", nil)
		return
	}

	// hash password from request
	newPassword := helper.Encode([]byte(datarequest.Password))
	datarequest.Password = string(newPassword)

	// bandingkan
	if cari.Username != datarequest.Username || cari.Password != datarequest.Password {
		response.Response(w, http.StatusOK, "Password salah", nil)
		return
	}

	// buat expired time nya
	// convert
	newWaktu, _ := strconv.Atoi(WAKTU)
	expiredTime := helper.ExpiredTime(newWaktu)

	// fill ke jwt
	token, err := tokenHelper.BuatJWT(ISS, AUD, LOGIN_SECRET, expiredTime)

	// cek jika generate gagal
	if err != nil {
		response.Response(w, http.StatusInternalServerError, "Gagal login", nil)
		return
	}

	// masukin ke model, kirim respon
	var tokensMaps tokenModel.Token
	tokensMaps.FullToken = token

	response.Response(w, http.StatusOK, "Login berhasil", tokensMaps)
}

func (l *loginHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	// convert
	newWaktu, _ := strconv.Atoi(WAKTU)

	// buat expired time nya
	expiredTime := helper.ExpiredTime(newWaktu)

	// fill ke jwt
	token, err := tokenHelper.BuatJWT(ISS, AUD, LOGIN_SECRET, expiredTime)

	// cek jika generate gagal
	if err != nil {
		response.Response(w, http.StatusInternalServerError, "Gagal login", nil)
		return
	}

	// masukin ke model, kirim respon
	var tokensMaps tokenModel.Token
	tokensMaps.FullToken = token
	response.Response(w, http.StatusOK, "Login berhasil", tokensMaps)
}
