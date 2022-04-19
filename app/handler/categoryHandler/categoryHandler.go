package categoryHandler

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/helper/redisHelper"
	"randi_firmansyah/app/helper/requestHelper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/models/categoryModel"
	"randi_firmansyah/app/service"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

var (
	key_redis   = "list_category_randi"
	HandlerName = "category"
	paramName   = "id"
)

type categoryHandler struct {
	service service.Service
	redis   *redis.Client
}

func NewCategoryHandler(categoryService service.Service, redis *redis.Client) *categoryHandler {
	return &categoryHandler{categoryService, redis}
}

func (h *categoryHandler) GetSemuaCategory(w http.ResponseWriter, r *http.Request) {
	// check redis with get response
	go func() {
		if data, err := redisHelper.GetRedisData(key_redis, h.redis); err == nil {
			response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), data)
			return
		}
	}()

	// select ke service
	listCategory, err := h.service.ICategoryService.FindAll()
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgGetAll(false, HandlerName), nil)
		return
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), listCategory)
}

func (h *categoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// check id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// get one data from redis
	go func() {
		if result, err := redisHelper.GetOneRedisData(id, key_redis, h.redis); err == nil {
			response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), result)
			return
		}
	}()

	// select ke service
	cari, err := h.service.ICategoryService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), cari)
}

func (h *categoryHandler) PostCategory(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest categoryModel.Category
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgCreate(false, HandlerName), nil)
		return
	}

	// insert
	created, err := h.service.ICategoryService.Create(datarequest)
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgCreate(false, HandlerName), nil)
		return
	}

	// delete cache from redis by key
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	// response success
	response.Response(w, http.StatusOK, response.MsgCreate(true, HandlerName), created)
}

func (h *categoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest categoryModel.Category
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgUpdate(false, HandlerName), nil)
		return
	}

	// cek id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgUpdate(false, HandlerName), nil)
		return
	}

	// cari data
	if _, err := h.service.ICategoryService.FindByID(newId); err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// update
	updated, err := h.service.ICategoryService.Update(newId, datarequest)
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgUpdate(false, HandlerName), nil)
		return
	}

	// clear redis cache
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	// response success
	response.Response(w, http.StatusOK, response.MsgUpdate(true, HandlerName), updated)
}

func (h *categoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// cek id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgDelete(false, HandlerName), nil)
		return
	}

	// cari data
	cari, err := h.service.ICategoryService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// delete
	deleted, err := h.service.ICategoryService.Delete(cari)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgDelete(false, HandlerName), nil)
		return
	}

	// clear redis cache
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	response.Response(w, http.StatusOK, response.MsgDelete(true, HandlerName), deleted)
}
