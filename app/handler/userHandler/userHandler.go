package userHandler

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/helper/redisHelper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/models/userModel"
	"randi_firmansyah/app/service"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

var (
	HandlerName = "User"
	key_redis   = "list_user_randi"
	paramName   = "id"
	read        = "read"
	create      = "create"
	update      = "update"
	delete      = "delete"
	detail      = "detail"
)

type userHandler struct {
	service service.Service
	redis   *redis.Client
}

func NewUserHandler(userService service.Service, redis *redis.Client) *userHandler {
	return &userHandler{userService, redis}
}

func (h *userHandler) GetSemuaUser(w http.ResponseWriter, r *http.Request) {
	// check redis with get response
	go func() {
		if data, err := redisHelper.GetRedisData(key_redis, h.redis); err == nil {
			response.ResponseSuccess(w, read, HandlerName, data)
			return
		}
	}()

	// select ke service
	listUser, err := h.service.IUserService.FindAll()
	if err != nil {
		response.ResponseInternalServerError(w)
		return
	}

	// success response
	response.ResponseSuccess(w, read, HandlerName, listUser)
}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// get one data from redis
	go func() {
		if result, err := redisHelper.GetOneRedisData(id, key_redis, h.redis); err == nil {
			response.ResponseSuccess(w, detail, HandlerName, result)
			return
		}
	}()

	// select ke service
	cari, err := h.service.IUserService.FindByID(id)
	if err != nil {
		response.ResponseBadRequest(w)
		return
	}

	// success response
	response.ResponseSuccess(w, detail, HandlerName, cari)
}

func (h *userHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		response.ResponseInternalServerError(w)
		return
	}

	// insert
	created, err := h.service.IUserService.Create(datarequest)
	if err != nil {
		response.Response(w, http.StatusBadRequest, "Username sudah digunakan", nil)
		return
	}

	// delete cache from redis by key
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	// response success
	response.ResponseSuccess(w, create, HandlerName, created)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		response.ResponseBadRequest(w)
		return
	}

	// update
	updated, err := h.service.IUserService.Update(id, datarequest)
	if err != nil {
		response.ResponseInternalServerError(w)
		return
	}

	// clear redis cache
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	// response success
	response.ResponseSuccess(w, update, HandlerName, updated)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// delete
	deleted, err := h.service.IUserService.Delete(id)
	if err != nil {
		response.ResponseBadRequest(w)
		return
	}

	// clear redis cache
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	response.ResponseSuccess(w, delete, HandlerName, deleted)
}
