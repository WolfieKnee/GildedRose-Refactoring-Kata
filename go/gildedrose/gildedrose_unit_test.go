package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Sulfuras_oneDay(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	expected := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	// Act
	gildedrose.UpdateQuality(items)

	// Assert
	for i, item := range expected {
		if items[i].SellIn != item.SellIn {
			t.Errorf("%s - SellIn: Expected %d but got %d ", item.Name, item.SellIn, items[i].SellIn)
		}
		if items[i].Quality != item.Quality {
			t.Errorf("%s - Quality: Expected %d but got %d ", item.Name, item.Quality, items[i].Quality)
		}
	}
}

func Test_Sulfuras_tenDay(t *testing.T) {
	// Arrange
	items := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	expected := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	// Act
	for day := 1; day <= 10; day++ {
		gildedrose.UpdateQuality(items)
	}

	// Assert
	for i, item := range expected {
		if items[i].SellIn != item.SellIn {
			t.Errorf("%s - SellIn: Expected %d but got %d ", item.Name, item.SellIn, items[i].SellIn)
		}
		if items[i].Quality != item.Quality {
			t.Errorf("%s - Quality: Expected %d but got %d ", item.Name, item.Quality, items[i].Quality)
		}
	}
}
