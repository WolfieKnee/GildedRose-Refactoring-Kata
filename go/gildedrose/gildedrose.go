package gildedrose

import (
	"strings"
)

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

// Methods for Item
func (item *Item) decrementItemQuality() {
	if item.Quality > minQual {
		item.Quality -= 1
	}
}

func (item *Item) incrementItemQuality() {
	if item.Quality < maxQual {
		item.Quality += 1
	}
}

func (item *Item) decrementItemSellIn() {
	item.SellIn -= 1
}

func (item *Item) updateItem() {
	item.SellIn -= 1
	if item.Quality > minQual {
		item.Quality -= 1
	}
}

// Methods for Sulfuras
// note - These do nothing, may be removed later
func (SulfurasItem *Sulfuras) updateItem() {}

// Methods for Passes
func (PassesItem *Passes) updateQuality() {
	PassesItem.incrementItemQuality()
	if PassesItem.SellIn <= 10 {
		PassesItem.incrementItemQuality()
	}
	if PassesItem.SellIn <= 5 {
		PassesItem.incrementItemQuality()
		// PassesItem.incrementItemQuality()
	}
}

// magic numbers
var maxQual int = 50
var minQual int = 0

// functions

func BaselineUpdateItem(item *Item) {

	// update the quality
	if item.Name == "Aged Brie" {
		item.incrementItemQuality()
	} else if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		item.decrementItemQuality()
	}

	// decrement the sell in
	item.decrementItemSellIn()

	// behaviours for SellIn < 0
	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			// set Q for expired tickets to zero
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				item.Quality = item.Quality - item.Quality
			} else {
				item.decrementItemQuality()
			}
		} else {
			// if item.Quality < maxQual {
			// 	item.Quality += 1
			// }
			item.incrementItemQuality()
		}
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {

		if strings.Contains(item.Name, "Sulfuras") {
			SulfurasItem := Sulfuras{item}
			SulfurasItem.updateItem()
		} else if strings.Contains(item.Name, "Backstage passes") {
			PassesItem := Passes{item}
			PassesItem.updateQuality()
			BaselineUpdateItem(item)
		} else {
			BaselineUpdateItem(item)
		}

	}
}
