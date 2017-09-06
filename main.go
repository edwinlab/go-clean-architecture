package main

import (
	"github.com/edwinlab/go-clean-architecture/app"
	"github.com/edwinlab/go-clean-architecture/domain"
	"github.com/edwinlab/go-clean-architecture/infrastructure"
)

func main() {
	payment := app.Payment{Amount: 42.16, Order: domain.Order{}}
	repo := infrastructure.MysqlPaymentsRepository{}

	err := repo.AddPayment(payment)
	if err != nil {
		panic(err)
	}
}
