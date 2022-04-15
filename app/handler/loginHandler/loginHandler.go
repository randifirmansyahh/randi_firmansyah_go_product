package loginHandler

import (
	"log"
	"net/http"
	"randi_firmansyah/app/helper/helper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/helper/tokenHelper"
	"randi_firmansyah/app/models/tokenModel"
	"strconv"
)

var (
	WAKTU        = tokenHelper.WAKTU
	AUD          = tokenHelper.AUD
	ISS          = tokenHelper.ISS
	LOGIN_SECRET = tokenHelper.LOGIN_SECRET
)

// func Login(w http.ResponseWriter, r *http.Request) {
// 	// cek user dan pass
// 	// decode from json
// 	decoder := json.NewDecoder(r.Body)

// 	// fill to model
// 	var datarequest userModel.User
// 	if err := decoder.Decode(&datarequest); err != nil {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// select ke db
// 	log.Println(dbGettingData)
// 	cari, err := userRepository.FindByID(datarequest.Id)
// 	if err != nil {
// 		log.Println(err)
// 		response.Response(w, http.StatusBadRequest, response.MsgNotFound("User"), cari)
// 		return
// 	}

// 	// bandingkan
// 	if cari.Id != datarequest.Id || cari.Username != datarequest.Username {
// 		log.Println(err)
// 		response.Response(w, http.StatusBadRequest, response.MsgNotFound("User"), cari)
// 	}

// 	// buat expired time nya
// 	expiredTime := helper.ExpiredTime(WAKTU)

// 	// fill ke jwt
// 	token, err := tokenHelper.BuatJWT(ISS, AUD, LOGIN_SECRET, expiredTime)

// 	// cek jika generate gagal
// 	if err != nil {
// 		log.Println("Error", err)
// 		response.Response(w, http.StatusInternalServerError, "Gagal membuat JWT !!", nil)
// 		return
// 	}

// 	// masukin ke model, kirim respon
// 	var tokensMaps tokenModel.Token
// 	tokensMaps.FullToken = token

// 	response.Response(w, http.StatusOK, "Berhasil generate token", token)
// }

func GenerateTokens(w http.ResponseWriter, r *http.Request) {
	// convert
	newWaktu, _ := strconv.Atoi(WAKTU)
	// buat expired time nya
	expiredTime := helper.ExpiredTime(newWaktu)

	// fill ke jwt
	token, err := tokenHelper.BuatJWT(ISS, AUD, LOGIN_SECRET, expiredTime)

	// cek jika generate gagal
	if err != nil {
		log.Println("Error", err)
		response.Response(w, http.StatusInternalServerError, "Gagal membuat JWT !!", nil)
		return
	}

	// masukin ke model, kirim respon
	var tokensMaps tokenModel.Token
	tokensMaps.FullToken = token
	response.Response(w, http.StatusOK, "Sukses membuat JWT !!", tokensMaps)
}
