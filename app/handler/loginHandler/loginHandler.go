package loginHandler

import (
	"encoding/json"
	"log"
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
		response.ResponseBadRequest(w)
		return
	}

	// select ke db
	newId := strconv.Itoa(datarequest.Id)
	cari, err := l.service.IUserService.FindByID(newId)
	if err != nil {
		log.Println(err)
		response.ResponseBadRequest(w)
		return
	}

	// bandingkan
	if cari.Id != datarequest.Id || cari.Username != datarequest.Username {
		log.Println(err)
		response.ResponseBadRequest(w)
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
		response.ResponseInternalServerError(w)
		return
	}

	// masukin ke model, kirim respon
	var tokensMaps tokenModel.Token
	tokensMaps.FullToken = token

	response.ResponTokenSucces(w, tokensMaps)
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
		response.ResponseInternalServerError(w)
		return
	}

	// masukin ke model, kirim respon
	var tokensMaps tokenModel.Token
	tokensMaps.FullToken = token
	response.ResponTokenSucces(w, tokensMaps)
}
