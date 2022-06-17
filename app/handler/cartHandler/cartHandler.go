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
			Username:       cart.User.Username,
			Product:        cart.Product,
			Qty:            cart.Qty,
			Total:          cart.Total,
			DateAuditModel: cart.DateAuditModel,
		})
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), newListCart)
}

func (h *cartHandler) GetCartByUsername(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	username := chi.URLParam(r, "username")

	// get one data from redis
	// if result, err := redisHelper.GetOneRedisData(id, key_redis, h.redis); err == nil {
	// 	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), result)
	// 	return
	// }

	// find user
	user, err := h.service.IUserService.FindByUsername(username)
	if err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, "username"), nil)
		return
	}

	// select ke service
	cari, err := h.service.ICartService.FindByUserID(user.Id)
	if err != nil {
		response.Response(w, http.StatusNotFound, response.MsgGetDetail(false, "username"), nil)
		return
	}

	// convert cart to cartResponse
	var cartRes []cartModel.CartResponse
	for _, val := range cari {
		cartRes = append(cartRes, cartModel.CartResponse{
			Id:             val.Id,
			Username:       val.User.Username,
			Product:        val.Product,
			Qty:            val.Qty,
			Total:          val.Total,
			DateAuditModel: val.DateAuditModel,
		})
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), cartRes)
}

func (h *cartHandler) PostCart(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest cartModel.CartRequest
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgCreate(false, HandlerName), nil)
		return
	}

	// find User
	user, err := h.service.IUserService.FindByUsername(datarequest.Username)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgGetDetail(false, "user"), nil)
		return
	}

	// find product
	product, err := h.service.IProductService.FindByID(datarequest.ProductId)
	if err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgGetDetail(false, "product"), nil)
		return
	}

	if datarequest.Qty > product.Qty {
		response.Response(w, http.StatusBadRequest, "Order melebihi stock", nil)
		return
	}

	// cartRequest to cartModel
	cart := cartModel.Cart{
		User_Id:    user.Id,
		Product_Id: datarequest.ProductId,
		Qty:        datarequest.Qty,
		Total:      datarequest.Qty * product.Harga,
	}

	// insert
	if _, err := h.service.ICartService.Create(cart); err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgCreate(false, HandlerName), nil)
		return
	}

	// delete cache from redis by key
	go redisHelper.ClearRedis(h.redis, key_redis)

	// response success
	response.Response(w, http.StatusOK, response.MsgCreate(true, HandlerName), nil)
}

func (h *cartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)
	qty := chi.URLParam(r, "qty")

	if id == "" {
		response.Response(w, http.StatusBadRequest, "id tidak boleh kosong", nil)
		return
	}

	if qty == "" {
		response.Response(w, http.StatusBadRequest, "qty tidak boleh kosong", nil)
		return
	}

	// cek id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	newQty, err := requestHelper.CheckIDInt(qty)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// cari data
	if _, err := h.service.ICartService.FindByID(newId); err != nil {
		response.Response(w, http.StatusNotFound, "cart tidak ditemukan", nil)
		return
	}

	// update
	if err = h.service.ICartService.Update(newId, newQty); err != nil {
		response.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// clear redis cache
	go redisHelper.ClearRedis(h.redis, key_redis)

	// response success
	response.Response(w, http.StatusOK, response.MsgUpdate(true, HandlerName), nil)
}

func (h *cartHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, paramName)

	// cek id
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// cari data
	cari, err := h.service.ICartService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, "cart tidak ditemukan", nil)
		return
	}

	// delete
	if _, err := h.service.ICartService.Delete(cari); err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgDelete(false, HandlerName), nil)
		return
	}

	// clear redis cache
	go redisHelper.ClearRedis(h.redis, key_redis)

	response.Response(w, http.StatusOK, response.MsgDelete(true, HandlerName), nil)
}
