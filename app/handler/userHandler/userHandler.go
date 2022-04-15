package userHandler

import (
	"net/http"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/service"

	"github.com/go-redis/redis/v8"
)

var (
	HandlerName = "User"
)

type userHandler struct {
	service service.Service
	redis   *redis.Client
}

func NewUserHandler(userService service.Service, redis *redis.Client) *userHandler {
	return &userHandler{userService, redis}
}

func (h *userHandler) GetSemuaUser(w http.ResponseWriter, r *http.Request) {
	// check redis
	// go func() {
	// 	if redisData, err := redisHelper.GetRedis(key_redis); err == nil {
	// 		// unmarshall from redis
	// 		var data []userModel.User
	// 		helper.UnMarshall(redisData, &data)

	// 		response.Response(w, http.StatusOK, response.MsgGetAll(HandlerName), data)
	// 		return
	// 	}
	// }()

	// select ke db
	listUser, err := h.service.IUserService.FindAll()
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgServiceErr(), nil)
		return
	}

	// // nge jadiin json
	// result, err := json.Marshal(listUser)
	// helper.CheckError(err)

	// ngeset ke redis secara async dan ngecek nya
	// go func() {
	// 	redisHelper.SetRedis(key_redis, result)
	// }()

	response.Response(w, http.StatusOK, response.MsgGetAll(HandlerName), listUser)
}

// func GetUserById(w http.ResponseWriter, r *http.Request) {
// 	// ambil parameter
// 	id := chi.URLParam(r, idParam)

// 	// check id
// 	if id == kosong {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// conv to int
// 	newId, err := strconv.Atoi(id)
// 	if err != nil {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// check redis
// 	go func() {
// 		if redisData, err := redisHelper.GetRedis(key_redis); err == nil {
// 			// unmarshall from redis
// 			var data []userModel.User
// 			helper.UnMarshall(redisData, &data)

// 			// search from redis
// 			if oneData, err := helper.SearchUser(data, newId); !err {
// 				response.Response(w, http.StatusOK, response.MsgGetAll(HandlerName), oneData)
// 				return
// 			}
// 		}
// 	}()

// 	// select ke db
// 	log.Println(dbGettingData)
// 	cari, err := userRepository.FindByID(newId)
// 	if err != nil {
// 		log.Println(err)
// 		response.Response(w, http.StatusBadRequest, response.MsgNotFound(HandlerName), nil)
// 		return
// 	}

// 	response.Response(w, http.StatusOK, response.MsgGetDetail(HandlerName), cari)
// }

// func PostUser(w http.ResponseWriter, r *http.Request) {
// 	// decode from json
// 	decoder := json.NewDecoder(r.Body)

// 	// fill to model
// 	var datarequest userModel.User
// 	if err := decoder.Decode(&datarequest); err != nil {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// hash password
// 	newPassword, err := helper.HashPassword(datarequest.Password)
// 	if err != nil {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// hash password
// 	datarequest.Password = newPassword
// 	create, err := userRepository.Create(datarequest)
// 	if err != nil {
// 		response.Response(w, http.StatusInternalServerError, response.MsgServiceErr(), nil)
// 		return
// 	}

// 	// clear redis cache
// 	redisHelper.ClearRedis(key_redis)

// 	response.Response(w, http.StatusOK, response.MsgTambah(HandlerName), create)
// }

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	// ambil parameter
// 	id := chi.URLParam(r, idParam)

// 	// check id
// 	if id == kosong {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// convert to int
// 	newId, err := strconv.Atoi(id)
// 	if err != nil {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// cari
// 	search, err := userRepository.FindByID(newId)
// 	if err != nil {
// 		log.Println(err)
// 		response.Response(w, http.StatusBadRequest, response.MsgNotFound(HandlerName), nil)
// 		return
// 	}

// 	// set id
// 	var datarequest userModel.User
// 	datarequest.Id = newId

// 	// delete
// 	if _, err := userRepository.Delete(datarequest); err != nil {
// 		response.Response(w, http.StatusInternalServerError, response.MsgServiceErr(), nil)
// 		return
// 	}

// 	// clear redis cache
// 	redisHelper.ClearRedis(key_redis)

// 	response.Response(w, http.StatusOK, response.MsgHapus(HandlerName), search)
// }

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	// ambil parameter
// 	id := chi.URLParam(r, idParam)

// 	// check id
// 	if id == kosong {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// convert to int
// 	newId, errInt := strconv.Atoi(id)
// 	if errInt != nil {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// cari
// 	if _, err := userRepository.FindByID(newId); err != nil {
// 		log.Println(err)
// 		response.Response(w, http.StatusBadRequest, response.MsgNotFound(HandlerName), nil)
// 		return
// 	}

// 	// decode
// 	decoder := json.NewDecoder(r.Body)
// 	var datarequest userModel.User
// 	if err := decoder.Decode(&datarequest); err != nil {
// 		response.Response(w, http.StatusBadRequest, response.MsgInvalidReq(), nil)
// 		return
// 	}

// 	// update
// 	datarequest.Id = newId
// 	updated, err := userRepository.Update(newId, datarequest)
// 	if err != nil {
// 		response.Response(w, http.StatusInternalServerError, response.MsgServiceErr(), nil)
// 		return
// 	}

// 	// claer redis cache
// 	redisHelper.ClearRedis(key_redis)

// 	response.Response(w, http.StatusOK, response.MsgUpdate(HandlerName), updated)
// }
