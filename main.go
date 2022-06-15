package main

import (
	"fmt"
	"go-with-db/database"
	"go-with-db/user"
	"log"
)

var (
	id                 int
	usersCount         int
	username, password string
	isActive           int
)

func main() {

	db, err := database.OpenDatabase()
	checkError(err)
	defer db.Close()

	database.CreateAndUseDB(db, "testdb")

	_, err = database.CreateTableDB(db, "users")
	checkError(err)

	choice := 0
	for {
		fmt.Printf(
			`Tanlovingizni kiriting >>
		1. User qo'shish
		2. Saqlangan userlarni ko'rish
		3. Username o'qali qidirish
		4. Limit bilan chiqarish
		5. Bazadan username orqali o'chirish
		6. User passwordni o'zgartirish
		>> `,
		)
		fmt.Scan(&choice)
		if choice > 0 && choice < 7 {
			break
		}
	}

	switch choice {
	case 1:
		usersArr := CreateUsersArray()
		err = database.InserMultipleDataToDB(db, usersArr)
		checkError(err)
		break
	case 2:
		row, err := database.SelectDataFromToDB(db, "users")
		checkError(err)
		for row.Next() {
			err = row.Scan(&id, &username, &password, &isActive)
			checkError(err)
			log.Printf("id: %v, username: %q, password: %q, isActive: %v", id, username, password, isActive)
		}
		break
	case 3:
		username := ""
		fmt.Printf("Usernameni kiriting >> ")
		fmt.Scan(&username)
		row, err := database.SelectDataByUsernameFromToDB(db, "users", username)
		checkError(err)
		for row.Next() {
			err = row.Scan(&id, &username, &password, &isActive)
			checkError(err)
			log.Printf("id: %v, username: %q, password: %q, isActive: %v", id, username, password, isActive)
		}
		break
	case 4:
		limit := 1
		fmt.Printf("Limitnini kiriting >> ")
		fmt.Scan(&limit)
		checkError(err)
		row, err := database.SelectDataWithLimitFromToDB(db, "users", limit)
		checkError(err)
		for row.Next() {
			err = row.Scan(&id, &username, &password, &isActive)
			checkError(err)
			log.Printf("id: %v, username: %q, password: %q, isActive: %v", id, username, password, isActive)
		}
		break
	case 5:
		username := ""
		fmt.Printf("Usernameni kiriting >> ")
		fmt.Scan(&username)
		_, err := database.DeleteDataByUsernameFromToDB(db, "users", username)
		checkError(err)
		log.Println("Data deleted")
		break
	case 6:
		username, password := "", ""
		fmt.Printf("Usernameni kiriting >> ")
		fmt.Scan(&username)
		fmt.Printf("Yangi passwordni kiriting >> ")
		fmt.Scan(&password)
		_, err := database.UpdateDataByUsernameFromToDB(db, "users", username, "password", password)
		checkError(err)
		log.Println("password changed")
		break
	}
}

func CreateUsersArray() (usersArr []user.User) {
	fmt.Printf("Nechta user qo'shmoqchisiz? >> ")
	fmt.Scan(&usersCount)

	for i := 0; i < usersCount; i++ {
		fmt.Printf("username, password, isActive [example: admin pas1 1] => ")
		fmt.Scan(&username, &password, &isActive)
		u, err := user.AddUser(username, password, isActive)
		checkError(err)

		usersArr = append(usersArr, u)
	}
	return
}

// check error
func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
