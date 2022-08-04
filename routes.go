package main

func (s *server) routes() {
	// Create
	s.router.HandleFunc("/customers", s.createCustomer()).Methods("POST")
	// Read
	s.router.HandleFunc("/customers", s.getAllCustomers()).Methods("GET")
	s.router.HandleFunc("/customers/{id}", s.extractIdFromPath(s.getCustomer)).Methods("GET")
	//Update
	s.router.HandleFunc("/customers/{id}", s.extractIdFromPath(s.updateCustomer)).Methods("PUT")
	//Delete
	s.router.HandleFunc("/customers/{id}", s.extractIdFromPath(s.deleteCustomer)).Methods("DELETE")
}
