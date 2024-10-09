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
	updateItem(item *Item)
}

// Utility Methods
func (item *Item) decrementItemQuality() {
	if item.Quality > MINQUAL {
		item.Quality -= 1
	}
}

func (item *Item) incrementItemQuality() {
	if item.Quality < MAXQUAL {
		item.Quality += 1
	}
}

func (item *Item) decrementItemSellIn() {
	item.SellIn -= 1
}

// Item updaters
func (item *Item) updateItem() {
	item.decrementItemSellIn()
	item.decrementItemQuality()
}

// note - These do nothing, may be removed later
func (SulfurasItem *Sulfuras) updateItem() {}

func (PassesItem *Brie) updateItem() {
	PassesItem.incrementItemQuality()
	PassesItem.decrementItemSellIn()
	if PassesItem.SellIn < 0 {
		PassesItem.incrementItemQuality()
	}

}

func (PassesItem *Passes) updateItem() {
	PassesItem.incrementItemQuality()
	if PassesItem.SellIn <= 10 {
		PassesItem.incrementItemQuality()
	}
	if PassesItem.SellIn <= 5 {
		PassesItem.incrementItemQuality()
		// PassesItem.incrementItemQuality()
	}
	PassesItem.decrementItemSellIn()
	if PassesItem.SellIn < 0 {
		PassesItem.Quality = PassesItem.Quality - PassesItem.Quality
	}

}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		var wrappedItem ItemWrapper
		switch {
		case strings.Contains(item.Name, "Sulfuras"):
			wrappedItem = Sulfuras{item}
			// SulfurasItem := Sulfuras{item}
			// SulfurasItem.updateItem()
		case strings.Contains(item.Name, "Backstage passes"):
			PassesItem := Passes{item}
			PassesItem.updateItem()
		case strings.Contains(item.Name, "Brie"):
			BrieItem := Brie{item}
			BrieItem.updateItem()
		default:
			item.updateItem()
		}
		wrappedItem.updateItem()

	}
}
