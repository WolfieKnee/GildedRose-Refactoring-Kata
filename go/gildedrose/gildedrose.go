package gildedrose

import (
	"strings"
)

var MAXQUAL int = 50
var MINQUAL int = 0

// Type definitions
type Item struct {
	Name            string
	SellIn, Quality int
}

type Sulfuras struct {
	*Item
}

type Passes struct {
	*Item
}

type Brie struct {
	*Item
}

type ItemWrapper interface {
	updateItem()
}

// Utility Methods
func (item *Item) decrementQuality(value ...int) {
	decrement := 1
	if len(value) > 0 {
		decrement = value[0]
	}
	if item.Quality > MINQUAL {
		item.Quality -= decrement
	}
}

func (item *Item) incrementQuality(value ...int) {
	increment := 1
	if len(value) > 0 {
		increment = value[0]
	}
	if item.Quality < MAXQUAL {
		item.Quality += increment
	}
}

func (item *Item) decrementSellIn(value ...int) {
	decrement := 1
	if len(value) > 0 {
		decrement = value[0]
	}
	item.SellIn -= decrement
}

// Item updaters
func (item *Item) updateItem() {
	item.decrementSellIn()
	item.decrementQuality()
}

// note - These do nothing, may be removed later
func (SulfurasItem *Sulfuras) updateItem() {}

func (PassesItem *Brie) updateItem() {
	PassesItem.decrementSellIn()
	if PassesItem.SellIn < 0 {
		PassesItem.incrementQuality(2)
	} else {
		PassesItem.incrementQuality()
	}
}

func (PassesItem *Passes) updateItem() {
	PassesItem.incrementQuality()
	if PassesItem.SellIn <= 10 {
		PassesItem.incrementQuality()
	}
	if PassesItem.SellIn <= 5 {
		PassesItem.incrementQuality()
		// PassesItem.incrementItemQuality()
	}
	PassesItem.decrementSellIn()
	if PassesItem.SellIn < 0 {
		PassesItem.Quality = 0
	}

}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		var wrappedItem ItemWrapper
		switch {
		case strings.Contains(item.Name, "Sulfuras"):
			wrappedItem = &Sulfuras{item}
		case strings.Contains(item.Name, "Backstage passes"):
			wrappedItem = &Passes{item}
		case strings.Contains(item.Name, "Brie"):
			wrappedItem = &Brie{item}
		default:
			wrappedItem = item
		}
		wrappedItem.updateItem()
	}
}
