package handlers

import (
	"CRM-Service/internal/models"
	"CRM-Service/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PartnerHandler struct {
	partnerService *services.PartnerService
}

func NewPartnerHandler(partnerService *services.PartnerService) *PartnerHandler {
	return &PartnerHandler{partnerService: partnerService}
}

func (handler *PartnerHandler) GetAllPartners(w http.ResponseWriter, r *http.Request) {
	partners, err := handler.partnerService.GetAllPartners()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(partners); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (handler *PartnerHandler) GetPartner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	partner, err := handler.partnerService.GetPartner(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(partner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *PartnerHandler) CreatePartner(w http.ResponseWriter, r *http.Request) {
	var partner models.Partner
	err := json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.partnerService.CreatePartner(&partner)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(partner); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (handler *PartnerHandler) UpdatePartner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var partner models.Partner
	if err := json.NewDecoder(r.Body).Decode(&partner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.partnerService.UpdatePartner(id, &partner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(partner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (handler *PartnerHandler) DeletePartner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.partnerService.DeletePartner(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
