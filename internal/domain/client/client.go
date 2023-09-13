package cliententity

import (
	personentity "github.com/willjrcom/sales-backend-go/internal/domain/person"
)

type Client struct {
	personentity.Person
	TotalOrders int `bun:"total_orders"`
}
