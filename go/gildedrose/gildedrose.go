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
func (item *Item) incrementItemQuality() {
	if item.Quality > minQual {
		item.Quality -= 1
	}
}

func (item *Item) incrementItemSellIn() {
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

	// Detect and cast type of item
	// note - Sulphuras don't do anything, so these may be removed later
	// if item.Name == "Sulfuras, Hand of Ragnaros" {
	if strings.Contains(item.Name, "Sulfuras") {
		// implement logic (based on the item's cast type)
		SulfurasItem := Sulfuras{Item: *item}
		SulfurasItem.updateItem()
	} else {
		if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
			item.incrementItemQuality()
		} else {
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
		}
		item.incrementItemSellIn()

		if item.SellIn < minQual {
			if item.Name != "Aged Brie" {
				if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
					item.incrementItemQuality()
				} else {
					item.Quality = item.Quality - item.Quality
				}
			} else {
				if item.Quality < maxQual {
					item.Quality += 1
				}
			}
		}
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		BaselineUpdateItem(item)
	}
}
