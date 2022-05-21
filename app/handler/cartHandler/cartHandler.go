package cartHandler

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/helper/redisHelper"
	"randi_firmansyah/app/helper/requestHelper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/models/cartModel"
	"randi_firmansyah/app/service"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
)

var (
	key_redis   = "list_cart_randi"
	HandlerName = "cart"
	paramName   = "id"
)

type cartHandler struct {
	service service.Service
	redis   *redis.Client
}

func NewCartHandler(cartService service.Service, redis *redis.Client) *cartHandler {
	return &cartHandler{cartService, redis}
}

func (h *cartHandler) GetSemuaCart(w http.ResponseWriter, r *http.Request) {
	// check redis with get response
	if data, err := redisHelper.GetRedisData(key_redis, h.redis); err == nil {
		response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), data)
		return
	}

	// select ke service
	listCart, err := h.service.ICartService.FindAll()
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgGetAll(false, HandlerName), nil)
		return
	}

	// convert cart to cartResponse
	var newListCart []cartModel.CartResponse
	for _, cart := range listCart {
		newListCart = append(newListCart, cartModel.CartResponse{
			Id:             cart.Id,
			User_Id:        cart.User_Id,
			Product_Id:     cart.Product_Id,
			Qty:            cart.Qty,
			Total:          cart.Total,
			DateAuditModel: cart.DateAuditModel,
		})
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), newListCart)
}

func (h *cartHandler) GetCartByID(w http.ResponseWriter, r *http.Request) {
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
	cari, err := h.service.ICartService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// convert cart to cartResponse
	newCart := cartModel.CartResponse{
		Id:             cari.Id,
		User_Id:        cari.User_Id,
		Product_Id:     cari.Product_Id,
		Qty:            cari.Qty,
		Total:          cari.Total,
		DateAuditModel: cari.DateAuditModel,
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), newCart)
}

func (h *cartHandler) PostCart(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest cartModel.Cart
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgCreate(false, HandlerName), nil)
		return
	}

	// insert
	created, err := h.service.ICartService.Create(datarequest)
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

func (h *cartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest cartModel.Cart
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
	if _, err := h.service.ICartService.FindByID(newId); err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// update
	updated, err := h.service.ICartService.Update(newId, datarequest)
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

func (h *cartHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// cek id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgDelete(false, HandlerName), nil)
		return
	}

	// cari data
	cari, err := h.service.ICartService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, HandlerName), nil)
		return
	}

	// delete
	deleted, err := h.service.ICartService.Delete(cari)
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
