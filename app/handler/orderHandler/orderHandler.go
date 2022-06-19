package orderHandler

import (
	"encoding/json"
	"net/http"
	"randi_firmansyah/app/helper/redisHelper"
	"randi_firmansyah/app/helper/requestHelper"
	"randi_firmansyah/app/helper/response"
	"randi_firmansyah/app/models/orderModel"
	"randi_firmansyah/app/service"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
)

var (
	key_redis   = "list_order_randi"
	HandlerName = "order"
	paramName   = "username"
)

type orderHandler struct {
	service service.Service
	redis   *redis.Client
}

func NewOrderHandler(orderService service.Service, redis *redis.Client) *orderHandler {
	return &orderHandler{orderService, redis}
}

func (h *orderHandler) GetSemuaOrder(w http.ResponseWriter, r *http.Request) {
	// check redis with get response
	if data, err := redisHelper.GetRedisData(key_redis, h.redis); err == nil {
		response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), data)
		return
	}

	// select ke service
	listOrder, err := h.service.IOrderService.FindAll()
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgGetAll(false, HandlerName), nil)
		return
	}

	// convert order to order response
	var newListOrder []orderModel.OrderResponse
	for _, order := range listOrder {
		newListOrder = append(newListOrder, orderModel.OrderResponse{
			Id:             order.Id,
			Username:       order.Username,
			Product:        order.Product,
			Qty:            order.Qty,
			Total:          order.Total,
			OrderStatus:    order.OrderStatus,
			DateAuditModel: order.DateAuditModel,
		})
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), newListOrder)
}

func (h *orderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, "id")

	// get one data from redis
	// if result, err := redisHelper.GetOneRedisData(id, key_redis, h.redis); err == nil {
	// 	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), result)
	// 	return
	// }

	// to int
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// select ke service
	cari, err := h.service.IOrderService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// convert order to order response
	orderResponse := orderModel.OrderResponse{
		Id:             cari.Id,
		Username:       cari.Username,
		Product:        cari.Product,
		Qty:            cari.Qty,
		Total:          cari.Total,
		OrderStatus:    cari.OrderStatus,
		DateAuditModel: cari.DateAuditModel,
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), orderResponse)
}

func (h *orderHandler) GetOrderByUsername(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	username := chi.URLParam(r, paramName)

	// get one data from redis
	// if result, err := redisHelper.GetOneRedisData(id, key_redis, h.redis); err == nil {
	// 	response.Response(w, http.StatusOK, response.MsgGetDetail(true, HandlerName), result)
	// 	return
	// }

	// select ke service
	cari, err := h.service.IOrderService.FindByUsername(username)
	if err != nil {
		response.Response(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// convert order to order response
	var orderResponse []orderModel.OrderResponse
	for _, val := range cari {
		orderResponse = append(orderResponse, orderModel.OrderResponse{
			Id:             val.Id,
			Username:       val.Username,
			Product:        val.Product,
			Qty:            val.Qty,
			Total:          val.Total,
			OrderStatus:    val.OrderStatus,
			DateAuditModel: val.DateAuditModel})
	}

	// success response
	response.Response(w, http.StatusOK, response.MsgGetAll(true, HandlerName), orderResponse)
}

func (h *orderHandler) PostOrder(w http.ResponseWriter, r *http.Request) {
	// decode and fill to model
	decoder := json.NewDecoder(r.Body)
	var datarequest orderModel.OrderReq
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, response.MsgCreate(false, HandlerName), nil)
		return
	}

	// check data request
	if err := h.CheckDatarequest(datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// find user by username
	if _, err := h.service.IUserService.FindByUsername(datarequest.Username); err != nil {
		response.Response(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// find product by id
	findProduct, err := h.service.IProductService.FindByID(datarequest.ProductId)
	if err != nil {
		response.Response(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// calculate total
	if datarequest.Qty > findProduct.Qty {
		response.Response(w, http.StatusBadRequest, "Stock tidak mencukupi", nil)
		return
	}
	request := orderModel.Order{
		Username:   datarequest.Username,
		Product_Id: datarequest.ProductId,
		Qty:        datarequest.Qty,
		Total:      findProduct.Harga * datarequest.Qty,
	}

	// insert
	if _, err := h.service.IOrderService.Create(request); err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgCreate(false, HandlerName), nil)
		return
	}

	// delete cache from redis by key
	go redisHelper.ClearRedis(h.redis, key_redis)

	// response success
	response.Response(w, http.StatusOK, response.MsgCreate(true, HandlerName), nil)
}

func (s *orderHandler) PayOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	saldo := chi.URLParam(r, "saldo")

	// check id int
	newId, err := requestHelper.CheckIDInt(id)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	newSaldo, err := requestHelper.CheckIDInt(saldo)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	findOrder, err := s.service.IOrderService.FindByID(newId)
	if err != nil {
		response.Response(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	if findOrder.OrderStatus {
		response.Response(w, http.StatusBadRequest, "Order tersebut telah dibayar", nil)
		return
	}

	if findOrder.Total > newSaldo {
		response.Response(w, http.StatusBadRequest, "Saldo tidak mencukupi", nil)
		return
	}

	findProduct, err := s.service.IProductService.FindByID(findOrder.Product_Id)
	if err != nil {
		response.Response(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	if findOrder.Qty > findProduct.Qty {
		response.Response(w, http.StatusBadRequest, "Stock tidak mencukupi", nil)
		return
	}

	findProduct.Qty -= findOrder.Qty
	_, err = s.service.IProductService.Update(findProduct.Id, findProduct)
	if err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgUpdate(false, HandlerName), nil)
		return
	}

	findOrder.OrderStatus = true
	if _, err := s.service.IOrderService.Update(findOrder.Id, findOrder); err != nil {
		response.Response(w, http.StatusInternalServerError, response.MsgUpdate(false, HandlerName), nil)
		return
	}

	response.Response(w, http.StatusOK, "Order telah berhasil di bayar", nil)
}

func (s *orderHandler) CheckDatarequest(datarequest orderModel.OrderReq) error {
	validate := validator.New()
	if err := validate.Struct(datarequest); err != nil {
		if errors := err.(validator.ValidationErrors); errors != nil {
			return errors
		}
	}
	return nil
}
