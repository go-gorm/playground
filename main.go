package main

import "fmt"

func main() {
	//
	// create data
	//

	user := User{
		Name: "John Random",
		Age:  48,
	}

	DB.Create(&user)

	toy := Toy{
		Name: "Elly",
	}

	pet := Pet{
		UserID: &user.ID,
		Name:   "Max",
		Toy:    toy,
	}

	DB.Create(&pet)

	//
	// this query works
	//

	var test User

	DB.Preload("Pets").Preload("Pets.Toy").Where("Age > ?", 30).First(&test)

	fmt.Println(test.Name)
	fmt.Println(test.Age)

	for _, pet := range test.Pets {

		fmt.Println(pet.Name)
		fmt.Println(pet.Toy.Name)
	}

	//
	// this query fails
	//

	rows, _ := DB.Model(&User{}).
		Joins("Pets").
		Joins("Pets.Toy").
		Where("Age > ?", 30).
		Order("id ASC").
		Limit(1024).
		Rows()

	defer rows.Close()

	for rows.Next() {
		var user User

		DB.ScanRows(rows, &user)

		fmt.Println(user.Name)
		fmt.Println(user.Age)

		for _, pet := range user.Pets {

			fmt.Println(pet.Name)
			fmt.Println(pet.Toy.Name)
		}
	}
}
