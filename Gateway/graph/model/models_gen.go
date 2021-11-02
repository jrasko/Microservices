// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Customer struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Order struct {
	ID       string           `json:"id"`
	Time     string           `json:"time"`
	State    OrderState       `json:"state"`
	Shop     *Shop            `json:"shop"`
	Customer *Customer        `json:"customer"`
	Products []*OrderProducts `json:"products"`
}

type OrderProducts struct {
	Order    *Order   `json:"order"`
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
	Price    float64  `json:"price"`
}

type Product struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description *string          `json:"description"`
	Shops       []*ProductInShop `json:"shops"`
	Image       string           `json:"image"`
}

type ProductInShop struct {
	Product   *Product `json:"Product"`
	Shop      *Shop    `json:"Shop"`
	Price     float64  `json:"price"`
	Currency  string   `json:"currency"`
	SalesUnit int      `json:"sales_unit"`
}

type Shop struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Location *string          `json:"location"`
	Products []*ProductInShop `json:"products"`
}

type OrderState string

const (
	OrderStateAccepted       OrderState = "ACCEPTED"
	OrderStateInProgress     OrderState = "IN_PROGRESS"
	OrderStateDone           OrderState = "DONE"
	OrderStateReadyForPickUp OrderState = "READY_FOR_PICK_UP"
	OrderStateRetrieved      OrderState = "RETRIEVED"
)

var AllOrderState = []OrderState{
	OrderStateAccepted,
	OrderStateInProgress,
	OrderStateDone,
	OrderStateReadyForPickUp,
	OrderStateRetrieved,
}

func (e OrderState) IsValid() bool {
	switch e {
	case OrderStateAccepted, OrderStateInProgress, OrderStateDone, OrderStateReadyForPickUp, OrderStateRetrieved:
		return true
	}
	return false
}

func (e OrderState) String() string {
	return string(e)
}

func (e *OrderState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderState", str)
	}
	return nil
}

func (e OrderState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}