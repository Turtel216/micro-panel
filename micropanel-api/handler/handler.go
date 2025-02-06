package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Turtel216/micro-panel/micropanel-api/server"
	"github.com/go-chi/chi"
)

type handler struct {
	ctx    context.Context
	server *server.Server
}

func NewHandler(server *server.Server) *handler {
	return &handler{
		ctx:    context.Background(),
		server: server,
	}
}

func (h *handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var p ProductReq
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	product, err := h.server.CreateProduct(h.ctx, toStorerProduct(p))
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(product)

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) getProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	product, err := h.server.GetProduct(h.ctx, i)
	if err != nil {
		http.Error(w, "Error getting product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(product)
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) listProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.server.ListProducts(h.ctx)
	if err != nil {
		http.Error(w, "Error listening products", http.StatusInternalServerError)
		return
	}

	var res []ProductRes
	for _, p := range products {
		res = append(res, toProductRes(&p))
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) updateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	var p ProductReq
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	product, err := h.server.GetProduct(h.ctx, i)
	if err != nil {
		http.Error(w, "Error getting product", http.StatusInternalServerError)
		return
	}

	patchProductReq(product, p)

	updatedProd, err := h.server.UpdateProduct(h.ctx, product)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(updatedProd)
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	if err := h.server.DeleteProduct(h.ctx, i); err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) createOrder(w http.ResponseWriter, r *http.Request) {
	var o OrderReq
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	created, err := h.server.CreateOrder(h.ctx, toStorerOrder(o))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := toOrderRes(created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) getOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	order, err := h.server.GetOrder(h.ctx, i)
	if err != nil {
		http.Error(w, "Error getting order", http.StatusInternalServerError)
		return
	}

	res := toOrderRes(order)
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) listOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.server.ListOrder(h.ctx)
	if err != nil {
		http.Error(w, "Error listening orders", http.StatusInternalServerError)
		return
	}

	var res []OrderRes
	for _, p := range orders {
		res = append(res, toOrderRes(&p))
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) deleteOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	if err := h.server.DeleteOrder(h.ctx, i); err != nil {
		http.Error(w, "Error deleting order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
