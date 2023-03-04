package main

import "fmt"

func main() {
	var ages map[string]int
	fmt.Println(ages)

	// before using the map it has to be instatieted using 'make' - otherwise it will err
	// make(make[Type]Type, initailSize) ...initSize is optional
	ages = make(map[string]int, 5)

	// first way to initialize a map
	ages["Frank"] = 24

	// Second way to initialize a map
	ages = map[string]int{
		"Frank": 24,
		"Jesus": 33,
	}

	fmt.Println(ages, len(ages))

	birthdays := map[string]string{
		"Frank":             "21-12-1997",
		"Jesus":             "25-12-0000",
		"LeBron James":      "30-12-1984",
		"Kylian Mbappe":     "20-12-1998",
		"Denzel Washington": "28-12-1954",
		"Jay-Z":             "04-12-1969",
		"John Snow":         "26-12-1986",
		"Pablo Escobar":     "01-12-1949",
		"Karim Benzema":     "19-12-1987",
		"Raheem Sterling":   "08-12-1994",
	}

	fmt.Println(birthdays, len(birthdays))

	// returns the expected value and true if item was found
	xmas, found := birthdays["Jesus"]
	fmt.Println(xmas, found)

	// to delete an item from the map
	delete(birthdays, "Jesus")
	fmt.Println("birthdays", birthdays, len(birthdays))

	_, found = birthdays["Jesus"]
	fmt.Println("Is Xmass available?", found)

	// adding an item
	birthdays["Jesus"] = "25-12-0000"
	fmt.Println("birthdays", birthdays)

	// to iterate a map use 'range' ...same rules as in slices
	// range in maps returns the current key and value in no particular order
	fmt.Println("Name\tBirthday")
	for name, birthday := range birthdays{
		fmt.Printf("%s\t\t%s\n", name, birthday)
	}

}
