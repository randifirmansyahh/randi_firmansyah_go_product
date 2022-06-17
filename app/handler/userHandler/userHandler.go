package userHandler

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/helper/redisHelper"
	"randi_firmansyah/app/helper/requestHelper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/models/userModel"
	"randi_firmansyah/app/service"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
)

var (
	HandlerName = "User"
	key_redis   = "list_user_randi"
	paramName   = "id"
)

type userHandler struct {
	service service.Service
	redis   *redis.Client
}

func NewUserHandler(userService service.Service, redis *redis.Client) *userHandler {
	return &userHandler{userService, redis}
}

func (h *userHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	username := chi.URLParam(r, paramName)

	// check redis with get response
	if data, err := redisHelper.GetRedisData(key_redis, h.redis); err == nil {
		response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), data)
		return
	}

	// select ke service
	cari, err := h.service.IUserService.FindByUsername(username)
	if err != nil {
		response.Response(w, http.StatusNotFound, "Data dengan username tersebut tidak ditemukan", nil)
		return
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), cari)
}

func (h *userHandler) GetSemuaUser(w http.ResponseWriter, r *http.Request) {
	// check redis with get response
	if data, err := redisHelper.GetRedisData(key_redis, h.redis); err == nil {
		response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), data)
		return
	}

	// select ke service
	listUser, err := h.service.IUserService.FindAll()
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgGetAll(false, HandlerName), nil)
		return
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), listUser)
}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// check id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, "ID harus berupa angka", nil)
	}

	// get one data from redis
	// if result, err := redisHelper.GetOneRedisData(id, key_redis, h.redis); err == nil {
	// 	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), result)
	// 	return
	// }

	// select ke service
	cari, err := h.service.IUserService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, "Data dengan ID tersebut tidak ditemukan", nil)
		return
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), cari)
}

func (h *userHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Data harus berupa json / request kurang lengkap", nil)
		return
	}

	validate := validator.New()
	err := validate.Struct(datarequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			response.Response(w, http.StatusBadRequest, errors.Error(), nil)
			return
		}
	}

	// find username
	if _, err := h.service.IUserService.FindByUsername(datarequest.Username); err == nil {
		response.Response(w, http.StatusBadRequest, "Username sudah digunakan", nil)
		return
	}

	// insert
	if _, err := h.service.IUserService.Create(datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, response.Internal_server_error, nil)
		return
	}

	// delete cache from redis by key
	go redisHelper.ClearRedis(h.redis, key_redis)

	// response success
	response.Response(w, http.StatusOK, response.MsgCreate(true, HandlerName), nil)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	username := chi.URLParam(r, "username")

	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Data harus berupa json / request kurang lengkap", nil)
		return
	}

	// validate
	validate := validator.New()
	err := validate.Struct(datarequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			response.Response(w, http.StatusBadRequest, errors.Error(), nil)
			return
		}
	}

	// find
	findUser, err := h.service.IUserService.FindByUsername(username)
	if err != nil {
		response.Response(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// update
	if _, err := h.service.IUserService.Update(findUser.Id, datarequest); err != nil {
		response.Response(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// clear redis cache
	go redisHelper.ClearRedis(h.redis, key_redis)

	// response success
	response.Response(w, http.StatusOK, response.MsgUpdate(true, HandlerName), nil)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// check id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, "ID harus berupa angka", nil)
	}

	// cari data
	cari, err := h.service.IUserService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, "Data dengan ID tersebut tidak ditemukan", nil)
		return
	}

	// delete
	deleted, err := h.service.IUserService.Delete(cari)
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgDelete(false, HandlerName), nil)
		return
	}

	// clear redis cache
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	response.Response(w, http.StatusOK, response.MsgDelete(true, HandlerName), deleted)
}
