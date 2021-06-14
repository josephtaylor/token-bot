package price

import (
	"token-bot/uniswap"
)

type Item interface {
	GetMessage(token *uniswap.Token) string
}

var items []Item

func init() {
	items = make([]Item, 0)
	registerItem(&PriceItem{})
	registerItem(&BuyItem{})
	registerItem(&ChartItem{})
	registerItem(&WebsiteItem{})
	registerItem(&TwitterItem{})
	registerItem(&ContractItem{})
}

func registerItem(item Item) {
	items = append(items, item)
}

func GetItems() []Item {
	return items
}
