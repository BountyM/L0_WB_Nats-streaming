package myrandom

import (
	"math/rand"

	"github.com/BountyM/L0_WB_Nats-streaming/models"
)

func RandomOrder() models.Order {
	var order models.Order
	order.Order_uid = RandStringRunes(19)
	order.Track_number = RandStringRunes(14)
	order.Entry = RandStringRunes(4)
	order.Locale = RandStringRunes(2)
	order.Internal_signature = RandStringRunes(14)
	order.Customer_id = RandStringRunes(4)
	order.Delivery_service = RandStringRunes(5)
	order.Shardkey = RandStringRunes(1)
	order.Sm_id = rand.Intn(100)
	order.Date_created = RandStringRunes(20)
	order.Oof_shard = RandStringRunes(1)
	// delivery
	order.Delivery.Name = RandStringRunes(11)
	order.Delivery.Phone = RandStringRunes(11)
	order.Delivery.Zip = RandStringRunes(7)
	order.Delivery.City = RandStringRunes(13)
	order.Delivery.Address = RandStringRunes(15)
	order.Delivery.Region = RandStringRunes(6)
	order.Delivery.Email = RandStringRunes(14)
	// payment
	order.Payment.Transaction = order.Order_uid
	order.Payment.Request_id = RandStringRunes(4)
	order.Payment.Currency = RandStringRunes(3)
	order.Payment.Provider = RandStringRunes(5)
	order.Payment.Amount = rand.Intn(4000)
	order.Payment.Payment_dt = rand.Intn(123456789)
	order.Payment.Bank = RandStringRunes(5)
	order.Payment.Delivery_cost = rand.Intn(3000)
	order.Payment.Goods_total = rand.Intn(634)
	order.Payment.Custom_fee = rand.Intn(144)
	// items
	items := make([]models.Item, 1)
	ni := rand.Intn(2) + 1
	for i := 0; i < ni; i++ {
		var item models.Item
		item.Chrt_id = rand.Intn(3000)
		item.Track_number = order.Track_number
		item.Price = rand.Intn(1000)
		item.Rid = RandStringRunes(14)
		item.Name = RandStringRunes(8)
		item.Sale = rand.Intn(100)
		item.Size = RandStringRunes(1)
		item.Total_price = rand.Intn(1000)
		item.Nm_id = rand.Intn(1000000)
		item.Brand = RandStringRunes(14)
		item.Status = rand.Intn(500)
	}
	order.Items = items
	return order
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
