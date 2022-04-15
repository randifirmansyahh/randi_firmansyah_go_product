package productHandler

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/helper/redisHelper"
	res "randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/models/productModel"
	"randi_firmansyah/app/service"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

var (
	key_redis   = "list_product_randi"
	HandlerName = "Product"
	paramName   = "id"
)

type productHandler struct {
	service service.Service
	redis   *redis.Client
}

func NewProductHandler(productService service.Service, redis *redis.Client) *productHandler {
	return &productHandler{productService, redis}
}

func (h *productHandler) GetSemuaProduct(w http.ResponseWriter, r *http.Request) {
	// check redis with get response
	go func() {
		if data, err := redisHelper.GetRedisData(key_redis, h.redis); err == nil {
			res.ResponseSuccess(w, "r", HandlerName, data)
			return
		}
	}()

	// select ke service
	listProduct, err := h.service.IProductService.FindAll()
	if err != nil {
		res.ResponseInternalServerError(w)
		return
	}

	// success response
	res.ResponseSuccess(w, "r", HandlerName, listProduct)
}

func (h *productHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// get one data from redis
	go func() {
		if result, err := redisHelper.GetOneRedisData(id, key_redis, h.redis); err == nil {
			res.ResponseSuccess(w, "r", HandlerName, result)
			return
		}
	}()

	// select ke service
	cari, err := h.service.IProductService.FindByID(id)
	if err != nil {
		res.ResponseBadRequest(w)
		return
	}

	// success response
	res.ResponseSuccess(w, "r", HandlerName, cari)
}

func (h *productHandler) PostProduct(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest productModel.Product
	if err := decoder.Decode(&datarequest); err != nil {
		res.ResponseInternalServerError(w)
		return
	}

	// insert
	create, err := h.service.IProductService.Create(datarequest)
	if err != nil {
		res.ResponseInternalServerError(w)
		return
	}

	// delete cache from redis by key
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	// response success
	res.ResponseSuccess(w, "c", HandlerName, create)
}

func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest productModel.Product
	if err := decoder.Decode(&datarequest); err != nil {
		res.ResponseBadRequest(w)
		return
	}

	// update
	updated, err := h.service.IProductService.Update(id, datarequest)
	if err != nil {
		res.ResponseInternalServerError(w)
		return
	}

	// clear redis cache
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	// response success
	res.ResponseSuccess(w, "u", HandlerName, updated)
}

func (h *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// delete
	deleted, err := h.service.IProductService.Delete(id)
	if err != nil {
		res.ResponseBadRequest(w)
		return
	}

	// clear redis cache
	go func() {
		redisHelper.ClearRedis(h.redis, key_redis)
	}()

	res.ResponseSuccess(w, "d", HandlerName, deleted)
}
