package handlers

import (
	"CRM-Service/internal/models"
	"CRM-Service/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ContactHandler struct {
	ContactService *services.ContactService
}

func NewContactHandler(contactService *services.ContactService) *ContactHandler {
	return &ContactHandler{ContactService: contactService}
}

func RegisterContactsRoutes(api *mux.Router, h *ContactHandler) {
	api.HandleFunc("/contact", h.GetAllContacts).Methods("GET")
	api.HandleFunc("/contact/{id}", h.GetContact).Methods("GET")
	api.HandleFunc("/contact", h.CreateContact).Methods("POST")
	api.HandleFunc("/contact/{id}", h.UpdateContact).Methods("PUT")
	api.HandleFunc("/contact/{id}", h.DeleteContact).Methods("DELETE")
}

func (contactHandler *ContactHandler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := contactHandler.ContactService.GetAllContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (contactHandler *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	contact, err := contactHandler.ContactService.GetContactById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (contactHandler *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := contactHandler.ContactService.CreateContact(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (contactHandler *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = contactHandler.ContactService.UpdateContact(id, &contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (contactHandler *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = contactHandler.ContactService.DeleteContact(id)
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
