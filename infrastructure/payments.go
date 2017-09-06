package infrastructure

import (
	"fmt"
	"github.com/edwinlab/go-clean-architecture/app"
)

// MysqlPaymentsRepository implements PaymentsRepository.
type MysqlPaymentsRepository struct{}

// AddPayment implements PaymentsRepository interface.
func (r MysqlPaymentsRepository) AddPayment(payment app.Payment) error {
	fmt.Printf("adding payment %+v\n", payment)

	return nil
}
