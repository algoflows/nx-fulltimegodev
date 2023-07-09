package main

import (
	"fmt"
)

var x = 10 // global variable
var (
	y = 20 // global variable
	z = 30 // global variable
)

const (
	version = 1
	pi      = 3.14159
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

type WeaponType int

const (
	Axe WeaponType = iota // 0
	Sword
	WoodenStick
	Knife
)

var (
	floatVar  float32 = 1.0       // 4 bytes (32 bits) of memory
	floatVar2 float64 = 1.0       // default type for floating point numbers is float64 - 8 bytes (64 bits) of memory
	name      string  = "introgo" // string is a sequence of bytes
	intVar32  int32   = 1         // 4 bytes (32 bits) of memory
	intVar64  int64   = 1         // 8 bytes (64 bits) of memory
	intVar    int     = 1         // default type for integers is int - 4 bytes (32 bits) of memory
	uintVar   uint    = 1         // 4 bytes (32 bits) of memory
	uintVar32 uint32  = 1         // 4 bytes (32 bits) of memory
	uintVar64 uint64  = 1         // 8 bytes (64 bits) of memory
	uint8Var  uint8   = 1         // 1 byte (8 bits) of memory
	uint16Var uint16  = 1         // 2 bytes (16 bits) of memory
	uint32Var uint32  = 1         // 4 bytes (32 bits) of memory
	uint64Var uint64  = 1         // 8 bytes (64 bits) of memory
	byteVar   byte    = 1         // alias for uint8 - same as uint8Var
	runeVar   rune    = 1         // alias for int32 - same as intVar32
	runVar    rune    = 'a'       // rune is an alias for int32 and represents a Unicode code point
)

type Player struct {
	name        string
	health      int
	attackPower float64
}

func getPlayerHealth(player Player) int {
	return player.health
}

func (player Player) getHealth() int {
	return player.health
}

func main() {
	var version int
	const pi = 3.14159
	// pi = 3.14 // error: cannot assign to pi

	version = 1
	version = 2
	version = 10

	fmt.Println(Hello("introgo"))
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("Version:", version)
	fmt.Println("Pi:", pi)

	player := Player{
		name:        "Captain America",
		health:      100,
		attackPower: 45.1,
	}
	fmt.Printf("Player: %+v\n", player)
	fmt.Printf("Player health: %d\n", getPlayerHealth(player))
	fmt.Printf("Player health: %d\n", player.getHealth())

	// maps
	users := map[string]int{
		"John": 28,
		"Jane": 36,
	}

	age, ok := users["John"]
	if ok {
		fmt.Printf("John's age: %d\n", age)
	}
	age2, ok := users["John2"]
	if !ok {
		fmt.Println("John not found", age2)
	}

	products := make(map[string]float64)
	products["shoes"] = 300.0
	products["hat"] = 20.0
	products = map[string]float64{
		"shoes": 300.0,
		"hat":   20.0,
	}

	users["Jack"] = 32
	users["Jill"] = 30

	fmt.Printf("Users: %+v\n", users)
	fmt.Printf("Users: %+v\n", users["John"])
	fmt.Printf("Users: %+v\n", len(users))
	fmt.Printf("Users: %+v\n", products)

	keys := make([]string, 0, len(products))
	for k := range products {
		keys = append(keys, k)
	}
	fmt.Printf("Keys: %+v\n", keys)

	delete(users, "John")
	fmt.Printf("Users: %+v\n", users)
	for k, v := range users {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
	// products
	for k, v := range products {
		fmt.Printf("Key: %s, Value: %f\n", k, v)
	}

	// arrays
	numbers := []int{1, 2, 3, 4, 5}
	otherNumbers := make([]int, 5)
	fmt.Printf("Numbers: %+v\n", numbers)
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Number: %d\n", numbers[i])
	}
	fmt.Printf("Other numbers: %+v\n", otherNumbers)
	clear(otherNumbers)

	// enums
	fmt.Printf("Damage of weapon (%s) (%d):\n ", Axe, getDamage(Axe))
	fmt.Printf("Damage of weapon (%s) (%d):\n ", Sword, getDamage(Sword))
	fmt.Printf("Damage of weapon (%s) (%d):\n ", WoodenStick, getDamage(WoodenStick))

	// loops
	for i := 0; i < 10; i++ {
		fmt.Printf("it: %d\n", i)
	}

	numbers = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Number: %d\n", numbers[i])
	}

	names := []string{"John", "Jane", "Jack"}
	for i, v := range names {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	for _, v := range names {
		if v == "Jack" {
			break
		}
		fmt.Printf("Value: %s\n", v)
	}
	fmt.Println("Done")

	for k, v := range users {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	// switch case
	name := "John"
	switch name {
	case "John":
		fmt.Println("John")
	case "Jane":
		fmt.Println("Jane")
	default:
		fmt.Println("Unknown")
	}
}

func (w WeaponType) String() string {
	switch w {
	case Axe:
		return "AXE"
	case Sword:
		return "SWORD"
	case WoodenStick:
		return "WOODEN_STICK"
	case Knife:
		return "KNIFE"
	default:
		return "UNKNOWN"
	}
}

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 40
	default:
		panic("weapon does not exist")
	}
}
