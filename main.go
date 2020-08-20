package main

import (
	"fmt"
	"strconv"
)

func main() {
	car1 := Car{23.65, 0.25}
	truck1 := Car{349, 0.5}

	house1 := House{"Jono", "Jalan Sukamaju no. 6 Jakarta Pusat, Jakarta"}
	house2 := House{"John Smith", "2267 Buckhannan Avenue, North Syracuse, NY"}

	store1 := Store{
		"CV. Berkah Mandiri",
		"Gang Mangga 2 no. 334, Jakarta Barat, Jakarta",
		[]Product{{"Roti Buaya", 100_000, "IDR", 8.6}},
	}

	store2 := Store{
		"George's Clothing",
		"56F Tail Ends Road, Cape Girardeau, Missouri",
		[]Product{{"Suit", 40, "USD", 8.6}},
	}

	foodTruck1 := FoodTruck{
		"Great Awesome Cake",
		450,
		0.33,
		[]Product{{"Strawberry Cake", 5, "USD", 9.2}, {"Orange Cake", 6, "USD", 8.9}},
	}

	streetVendor1 := StreetVendor{
		"Mr. Roger",
		[]Product{{"Super Duper Awesome Hot Dogs", 10, "USD", 9.4}},
	}

	// Drivable collection
	drivables := []Drivable{car1, truck1, foodTruck1}

	for _, drivable := range drivables {
		fuelConsumption := drivable.fuelConsumption()
		fmt.Println(fuelConsumption)
	}

	// Addressable collection
	addressables := []Addressable{house1, house2, store1, store2}

	for _, addressable := range addressables {
		description := addressable.description()
		fmt.Println(description)
	}

	// Shoppable collection
	shoppables := []Shoppable{store1, store2, foodTruck1, streetVendor1}

	for _, shoppable := range shoppables {
		description := shoppable.description()
		fmt.Println(description)
	}

	// shoppableDescription(store1)
	// shoppableDescription(foodTruck1)
}

func shoppableDescription(shoppable Shoppable) {
	desc := shoppable.description()
	fmt.Println(desc)
}

// Drivable sets an interface for drivable vehicles.
type Drivable interface {
	fuelConsumption() float64
}

// Addressable sets an interface for anything that can have address.
type Addressable interface {
	description() string
}

// Shoppable sets an interface for anything we can shop from.
type Shoppable interface {
	description() string
	// averageRating() float64
}

type Car struct {
	Mileage          float64
	FuelLiterPerMile float64
}

func (car Car) fuelConsumption() float64 {
	return car.Mileage * car.FuelLiterPerMile
}

type House struct {
	OwnerName string
	Address   string
}

func (house House) description() string {
	return "This house belongs to " + house.OwnerName + ", it is located at: " + house.Address
}

type Product struct {
	Name     string
	Price    float64
	Currency string
	Rating   float64
}

type Store struct {
	Name     string
	Address  string
	Products []Product
}

func (store Store) description() string {
	return store.Name + " is located at " + store.Address + ", and it sells: " + productStrings(store.Products)
}

type FoodTruck struct {
	ShopName         string
	Mileage          float64
	FuelLiterPerMile float64
	Products         []Product
}

func (ft FoodTruck) description() string {
	mileageString := strconv.FormatFloat(ft.Mileage, 'f', 2, 64)
	return "Food Truck: " + ft.ShopName + " has traveled " + mileageString + ", and it sells: " + productStrings(ft.Products)
}

func (ft FoodTruck) fuelConsumption() float64 {
	return ft.Mileage * ft.FuelLiterPerMile
}

type StreetVendor struct {
	OwnerName string
	Products  []Product
}

func (sv StreetVendor) description() string {
	return "Street Vendor with owner: " + sv.OwnerName + " sells: " + productStrings(sv.Products)
}

func productStrings(products []Product) string {
	productString := ""
	for _, product := range products {
		priceString := strconv.FormatFloat(product.Price, 'f', 2, 64)
		ratingString := strconv.FormatFloat(product.Rating, 'f', 1, 64)
		productString += product.Name + " price: " + priceString + " with rating: " + ratingString + ", "
	}
	return productString
}
