package gildedrose

import "strings"

//  Type definitions
type Item struct {
	Name            string
	SellIn, Quality int
}

type Sulfuras struct {
	Item
}

// Methods for Item
func (item *Item) decrementItemQuality() {
	if item.Quality > minQual {
		item.Quality -= 1
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

// note - These do nothing, may be removed later
func (SulfurasItemitem *Sulfuras) updateItem() {}

// magic numbers
var maxQual int = 50
var minQual int = 0

// functions

func BaselineUpdateItem(item *Item) {
	if item.Name == "Aged Brie" || item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.Quality < maxQual {
			item.Quality += 1
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.SellIn < 11 {
					if item.Quality < maxQual {
						item.Quality += 1
					}
				}
				if item.SellIn < 6 {
					if item.Quality < maxQual {
						item.Quality += 1
					}
				}
			}
		}
	} else {
		item.decrementItemQuality()
	}
	item.decrementItemSellIn()

	if item.SellIn < minQual {
		if item.Name != "Aged Brie" {
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
		} else {
			BaselineUpdateItem(item)
		}
	}
}
