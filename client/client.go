package main

import (
	"fmt"
	"net/rpc"
)

// Struct should be of exported type
type Item struct {
	Title string
	Body  string
}

// CRUD operations using RPCs from server running on 4040
func main() {
	var response Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		fmt.Println("Error contacting server: ", err)
	}

	item1 := make([]string, 2)
	item1[0] = "First"
	item1[1] = "First item"

	item2 := make([]string, 2)
	item2[0] = "Second"
	item2[1] = "Second item"

	item3 := make([]string, 2)
	item3[0] = "Third"
	item3[1] = "Third item"

	// Initial database - GetDB() API
	client.Call("API.GetDB", "", &db)
	fmt.Println("Initial Database: ", db)

	// Creating Items using the CreateItem() API
	client.Call("API.CreateItem", &item1, &response)
	a := response
	fmt.Println("Created Item: ", a)
	client.Call("API.CreateItem", &item2, &response)
	b := response
	fmt.Println("Created Item: ", b)
	client.Call("API.CreateItem", &item3, &response)
	c := response
	fmt.Println("Created Item: ", c)

	// Adding items to database using AddItem() API
	// a := Item{"First", "First Item"}
	// b := Item{"Second", "Second Item"}
	// c := Item{"Third", "Third Item"}

	client.Call("API.AddItem", a, &response)
	client.Call("API.AddItem", b, &response)
	client.Call("API.AddItem", c, &response)

	// // Database after adding three items
	client.Call("API.GetDB", "", &db)
	fmt.Println("Inserted three records: ", db)

	// // Delete second item using DeleteItem() API
	toDelete := Item{Title: "Second", Body: " "}
	client.Call("API.DeleteItem", &toDelete, &response)
	fmt.Println("Deleted Item: ", response.Title)

	// // Database after deleting the second item
	client.Call("API.GetDB", "", &db)
	fmt.Println("Deleted second record: ", db)

	// // Edit the second item's title and body
	toEdit := Item{Title: "Third", Body: "This is now the second item"}
	client.Call("API.EditItem", &toEdit, &response)
	fmt.Println("Edited item: ", response.Title)

	// // Database after editing the second item
	client.Call("API.GetDB", "", &db)
	fmt.Println("Edited second record: ", db)
}
