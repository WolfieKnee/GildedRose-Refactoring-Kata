package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Item_oneDay(t *testing.T) {
	items := []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Elixir of the Mongoose", 5, 7},
	}

	expected := []*gildedrose.Item{
		{"+5 Dexterity Vest", 9, 19},
		{"Elixir of the Mongoose", 4, 6},
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

func Test_Item_tenDay(t *testing.T) {
	items := []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Elixir of the Mongoose", 5, 7},
	}

	expected := []*gildedrose.Item{
		{"+5 Dexterity Vest", 0, 10},
		{"Elixir of the Mongoose", -5, 0},
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

func Test_Tickets_Q_increases_by_one(t *testing.T) {
	items := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	}

	expected := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
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

func Test_Tickets_Q_increases_by_two(t *testing.T) {
	items := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 10, 20},
	}

	expected := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 9, 22},
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

func Test_Tickets_Q_increases_by_3(t *testing.T) {
	items := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 5, 20},
	}

	expected := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 4, 23},
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

func Test_Tickets_Q_zero(t *testing.T) {
	items := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 0, 20},
	}

	expected := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
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

func Test_Brie_oneDay(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", 2, 0},
	}

	expected := []*gildedrose.Item{
		{"Aged Brie", 1, 1},
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

func Test_Brie_tenDay(t *testing.T) {
	items := []*gildedrose.Item{
		{"Aged Brie", 2, 0},
	}

	expected := []*gildedrose.Item{
		{"Aged Brie", -8, 18},
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

func Test_Conjured(t *testing.T) {
	// Arrange
	var items = []*gildedrose.Item{

		{"Conjured Mana Cake", 3, 6},
	}

	var expected = []*gildedrose.Item{
		{"Conjured Mana Cake", 2, 4},
	}
	// Act
	gildedrose.UpdateQuality(items)
	// Assert
	for i, item := range expected {
		if items[i].Name != item.Name {
			t.Errorf("Name: Expected %s but got %s ", item.Name, items[i].Name)
		}
		if items[i].SellIn != item.SellIn {
			t.Errorf("%s - SellIn: Expected %d but got %d ", item.Name, item.SellIn, items[i].SellIn)
		}
		if items[i].Quality != item.Quality {
			t.Errorf("%s - Quality: Expected %d but got %d ", item.Name, item.Quality, items[i].Quality)
		}
	}
}
