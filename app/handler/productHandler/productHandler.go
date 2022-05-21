package productHandler

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/helper/redisHelper"
	"randi_firmansyah/app/helper/requestHelper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/models/productModel"
	"randi_firmansyah/app/service"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

var (
	key_redis   = "list_product_randi"
	HandlerName = "product"
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

	if data, err := redisHelper.GetRedisData(key_redis, h.redis); err == nil {
		response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), data)
		return
	}

	// select ke service
	listProduct, err := h.service.IProductService.FindAll()
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgGetAll(false, HandlerName), nil)
		return
	}

	// convert product to product response
	var newListProduct []productModel.ProductResponse
	for _, product := range listProduct {
		newListProduct = append(newListProduct, productModel.ProductResponse{
			Id:             product.Id,
			Category_Id:    product.Category_Id,
			Nama:           product.Nama,
			Harga:          product.Harga,
			Qty:            product.Qty,
			Image:          product.Image,
			DateAuditModel: product.DateAuditModel,
		})
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), newListProduct)
}

func (h *productHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// check id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// get one data from redis
	if result, err := redisHelper.GetOneRedisData(id, key_redis, h.redis); err == nil {
		response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), result)
		return
	}

	// select ke service
	cari, err := h.service.IProductService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// convert product to product response
	newCari := productModel.ProductResponse{
		Id:             cari.Id,
		Category_Id:    cari.Category_Id,
		Nama:           cari.Nama,
		Harga:          cari.Harga,
		Qty:            cari.Qty,
		Image:          cari.Image,
		DateAuditModel: cari.DateAuditModel,
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), newCari)
}

func (h *productHandler) PostProduct(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest productModel.Product
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgCreate(false, HandlerName), nil)
		return
	}

	// insert
	created, err := h.service.IProductService.Create(datarequest)
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

func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest productModel.Product
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
	if _, err := h.service.IProductService.FindByID(newId); err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// update
	updated, err := h.service.IProductService.Update(newId, datarequest)
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

func (h *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// cek id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgDelete(false, HandlerName), nil)
		return
	}

	// cari data
	cari, err := h.service.IProductService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// delete
	deleted, err := h.service.IProductService.Delete(cari)
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
