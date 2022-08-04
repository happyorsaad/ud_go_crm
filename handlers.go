package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type withId func(string, http.ResponseWriter, *http.Request)

func parseCustomerFromRequest(r *http.Request) (Customer, error) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (s *server) extractIdFromPath(f withId) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if id, ok := mux.Vars(r)["id"]; ok {
			if s.db.hasId(id) {
				f(id, w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("id not found in the backend"))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id is missing in the request path"))
		}
	}
}

func (s *server) getAllCustomers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.db.customers)
	}
}

func (s *server) getCustomer(id string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.db.customers[id])
}

func (s *server) createCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customer, err := parseCustomerFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.db.createCustomer(customer)
	}
}

func (s *server) updateCustomer(id string, w http.ResponseWriter, r *http.Request) {
	customer, err := parseCustomerFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.db.updateCustomer(id, customer)
}

func (s *server) deleteCustomer(id string, w http.ResponseWriter, r *http.Request) {
	s.db.deleteCustomer(id)
}
