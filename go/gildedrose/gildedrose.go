package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func (item *Item) incrementItemQuality() {
	if item.Quality > 0 {
		if item.Name != "Sulfuras, Hand of Ragnaros" {
			item.Quality -= 1
		}
	}
}

func (item *Item) incrementItemSellIn() {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn -= 1
	}

}

func UpdateItem(item *Item) {

	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		item.incrementItemQuality()
	} else {
		if item.Quality < 50 {
			item.Quality += 1
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.SellIn < 11 {
					if item.Quality < 50 {
						item.Quality += 1
					}
				}
				if item.SellIn < 6 {
					if item.Quality < 50 {
						item.Quality += 1
					}
				}
			}
		}
	}
	item.incrementItemSellIn()

	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				item.incrementItemQuality()
			} else {
				item.Quality = item.Quality - item.Quality
			}
		} else {
			if item.Quality < 50 {
				item.Quality += 1
			}
		}
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		UpdateItem(item)
	}
}
