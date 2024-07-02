package handlers

import (
	"CRM-Service/internal/models"
	"CRM-Service/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BidHandler struct {
	bidService *services.BidService
}

func NewBidHandler(bidService *services.BidService) *BidHandler {
	return &BidHandler{bidService: bidService}
}

func RegisterBidRoutes(api *mux.Router, h *BidHandler) {
	api.HandleFunc("/bid", h.GetAllBids).Methods("GET")
	api.HandleFunc("/partner/{id}", h.GetBid).Methods("GET")
	api.HandleFunc("/bid", h.CreateBid).Methods("POST")
	api.HandleFunc("/bid/{id}", h.UpdateBid).Methods("PUT")
	api.HandleFunc("/bid/{id}", h.DeleteBid).Methods("DELETE")
}

func (h *BidHandler) GetAllBids(w http.ResponseWriter, r *http.Request) {
	bids, err := h.bidService.GetAllBids()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(bids); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *BidHandler) GetBid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bid, err := h.bidService.GetBid(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(bid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *BidHandler) CreateBid(w http.ResponseWriter, r *http.Request) {
	var bid models.Bid
	if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.bidService.CreateBid(&bid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(bid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *BidHandler) UpdateBid(w http.ResponseWriter, r *http.Request) {
	var bid models.Bid
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.bidService.UpdateBid(id, &bid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(bid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *BidHandler) DeleteBid(w http.ResponseWriter, r *http.Request) {
	var bid models.Bid
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.bidService.DeleteBid(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(bid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
