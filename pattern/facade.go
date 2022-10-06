package main

import (
	"log"
	"math/rand"
	"strconv"
)

// Facade:
//
// Это структурный паттерн, главная цель которого создать структурную обертку (фасад/api) между сложной системой и бизнес-логикой
//
// Достоинвства:
// - Сокрытие/Изоляция сложной логики для конечного потребителя
// - Единый интерфейс для н-го списка интерфейсов подсистем
// - Если мы захотим изменить/заменить имплементацию сложной логики фасада, это не затронет конечного пользователя
// - Упращает рефакторинг кода, т.к вся логика находится в одном месте
// - Выоский уровень абстракции над сложной системой
//
// Недостатки:
// - Большая кодавая база
//
// Цели использования:
// - Инкапсулировать большой повторяющийся код для конкретных процессов и предоставлять его через интерфейс фасада
//
// ПРИМЕР РЕАЛЕЗАЦИИ:
type Order struct {
	userID   string
	products []Product
	amount   int
}
type Product struct {
	Name string
}

// Payment Service
type PaymentService struct {
	AuthToken string
}

func NewOrderService(token string) *PaymentService {
	return &PaymentService{AuthToken: token}
}

func (s PaymentService) ExecutePayent(amount int, sender string) {
	log.Printf("[Info] transcation from %s with amount %d", sender, amount)
}

// Fulfillment Service
type RealizationService struct {
	AuthToken string
}

func NewRealizationService(token string) *RealizationService {
	return &RealizationService{AuthToken: token}
}

func (s RealizationService) CreateOrder(product []Product) string {

	orderID := strconv.Itoa(rand.Intn(10000))
	log.Printf("[Info] creating new order with id - %s\n", orderID)

	return orderID
}

// Order tracking Service
type OrderTrackingService struct {
	AuthToken string
}

func NewOrderTrackingService(token string) *OrderTrackingService {
	return &OrderTrackingService{AuthToken: token}
}

func (s *OrderTrackingService) TrackOrder(orderID string) {

	log.Printf("[Info] tracking order with ID - %s", orderID)
}

// Order Facade
type OrderFacade struct {
	Track    *OrderTrackingService
	Fullfill *RealizationService
	Payment  *PaymentService
}

func NewOrderFacade() *OrderFacade {

	return &OrderFacade{
		Payment:  NewOrderService("bnr-256-112-567-obn"),
		Fullfill: NewRealizationService("zqc-256-167-123-vcq"),
		Track:    NewOrderTrackingService("abc-123-456-789-qwe"),
	}
}

func (f *OrderFacade) PayAndPrepareOrder(order Order) string {

	f.Payment.ExecutePayent(order.amount, order.userID)
	{
		// handle erros and other stuff
	}
	orderID := f.Fullfill.CreateOrder(order.products)

	return orderID
}

func main() {

	order := Order{
		amount: 5,
		userID: "123456789",
		products: []Product{
			{"Logitech G Pro Wireless"},
			{"Logitech G703"},
			{"Logitech PTZ Pro 2"},
		},
	}

	orderHandler := NewOrderFacade()

	orderID := orderHandler.PayAndPrepareOrder(order)

	orderHandler.Track.TrackOrder(orderID)

}
