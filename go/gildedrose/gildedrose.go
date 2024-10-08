package gildedrose

import (
	"fmt"
	"strings"
)

// Type definitions
type Item struct {
	Name            string
	SellIn, Quality int
}

type Sulfuras struct {
	Item
}

type Passes struct {
	Item
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
func (PassesItem *Passes) updateItem() {
	fmt.Print("hello from passes \n")
}

func (PassesItem *Passes) UpdatePassQuality() {
	if PassesItem.SellIn < 11 {
		if PassesItem.Quality < maxQual {
			PassesItem.incrementItemQuality()
		}
	}
	if PassesItem.SellIn < 6 {
		if PassesItem.Quality < maxQual {
			PassesItem.incrementItemQuality()
		}
	}
}

// magic numbers
var maxQual int = 50
var minQual int = 0

// functions

func BaselineUpdateItem(item *Item) {

	// update the quality
	if item.Name == "Aged Brie" {
		if item.Quality < maxQual {
			item.Quality += 1
		}
	} else if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.Quality < maxQual {
			item.Quality += 1
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {

				if item.SellIn < 11 {
					if item.Quality < maxQual {
						item.incrementItemQuality()
					}
				}
				if item.SellIn < 6 {
					if item.Quality < maxQual {
						item.incrementItemQuality()
					}
				}
			}
		}

	} else {
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
			if item.Quality < maxQual {
				item.Quality += 1
			}
		}
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		if strings.Contains(item.Name, "Sulfuras") {
			SulfurasItem := Sulfuras{Item: *item}
			SulfurasItem.updateItem()
		} else if strings.Contains(item.Name, "Backstage passes") {
			PassesItem := Passes{Item: *item}
			PassesItem.updateItem()
			BaselineUpdateItem(item)
		} else {
			BaselineUpdateItem(item)
		}
	}
}
