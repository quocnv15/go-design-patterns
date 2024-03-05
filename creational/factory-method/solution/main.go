package main

import (
	"errors"
	"fmt"
	"log"
)

type Drink interface {
	Drink()
}

type Food interface {
	Eat()
}

type Voucher struct {
	Drink
	Food
}

type Coffee struct{}

func (Coffee) Drink() {
	fmt.Println("Drinking coffee")
}

type Beer struct{}

func (Beer) Drink() {
	fmt.Println("Drinking beer")
}

type Pizza struct{}

func (Pizza) Eat() {
	fmt.Println("Eating pizza")
}

type Cake struct{}

func (Cake) Eat() {
	fmt.Println("Eating cake")
}

type GrilledOctopus struct{}

func (GrilledOctopus) Eat() {
	fmt.Println("Eating grilled octopus")
}

type VoucherAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}

func (v Voucher) applyVoucher() {
	v.Drink.Drink()
	v.Food.Eat()
}

type CoffeeMorningVoucherFactory struct{}

func (CoffeeMorningVoucherFactory) GetDrink() Drink {
	return Coffee{}
}
func (CoffeeMorningVoucherFactory) GetFood() Food {
	return Cake{}
}

type BeerNightVoucherFactory struct{}

func (BeerNightVoucherFactory) GetDrink() Drink {
	return Beer{}
}
func (BeerNightVoucherFactory) GetFood() Food {
	return GrilledOctopus{}
}

type VouccherType int

const (
	CoffeeMorning VouccherType = iota
	BeerNight
)

func GetVoucherByType(t VouccherType) (VoucherAbstractFactory, error) {
	switch t {
	case CoffeeMorning:
		return CoffeeMorningVoucherFactory{}, nil
	case BeerNight:
		return BeerNightVoucherFactory{}, nil
	default:
		return nil, errors.New("Invalid voucher type")
	}
}

func GetVoucher(factory VoucherAbstractFactory) Voucher {
	return Voucher{
		Drink: factory.GetDrink(),
		Food:  factory.GetFood(),
	}
}
func main() {
	voucherFactory, err := GetVoucherByType(CoffeeMorning)

	if err != nil {
		log.Fatal(err)
	}

	myVoucher := GetVoucher(voucherFactory)
	myVoucher.applyVoucher()

	fmt.Println("I'm happy with this voucher and come back to use it next time.", myVoucher)

	voucherFactory1, err := GetVoucherByType(BeerNight)

	if err != nil {
		log.Fatal(err)
	}

	myVoucher1 := GetVoucher(voucherFactory1)
	myVoucher1.applyVoucher()
}
