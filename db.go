package main

import "github.com/google/uuid"

type database struct {
	customers map[string]Customer
}

func seedDatabase(db *database) {
	c1 := Customer{
		Name:      "saad",
		Role:      "admin",
		Email:     "saad@gmail.com",
		Contacted: true,
		Phone:     "1122",
	}

	db.createCustomer(c1)

	c2 := Customer{
		Name:      "alina",
		Role:      "admin",
		Email:     "alina@gmail.com",
		Contacted: false,
		Phone:     "1133",
	}

	db.createCustomer(c2)

	c3 := Customer{
		Name:      "ali",
		Role:      "customer",
		Email:     "ali@gmail.com",
		Contacted: true,
		Phone:     "113344",
	}

	db.createCustomer(c3)
}
func NewDatabase() *database {
	db := database{}
	db.customers = make(map[string]Customer)
	return &db
}

func (db *database) hasId(id string) bool {
	_, ok := db.customers[id]
	return ok
}

func (db *database) getCustomer(id string) Customer {
	return db.customers[id]
}

func (db *database) deleteCustomer(id string) {
	delete(db.customers, id)
}

func (db *database) createCustomer(c Customer) string {
	id := uuid.New().String()
	db.customers[id] = c
	return id
}

func (db *database) updateCustomer(
	id string, c Customer) string {
	db.customers[id] = c
	return id
}
