package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

// Struct should have exported type
type Item struct {
	Title string
	Body  string
}

var db []Item

// RPCs should always
// 1. Be Go methods (should have receiver argument)
// 2. Be Exported Functions
// 3. Have two arguments of exported types
// 4. Have the second argument as a pointer
// 5. Have return type as error type

func (a *API) GetDB(title string, response *[]Item) error {
	*response = db
	return nil
}

func (a *API) GetByName(title string, response *Item) error {
	var retItem Item

	for _, item := range db {
		if item.Title == title {
			retItem = item
		}
	}
	*response = retItem
	return nil
}

func (a *API) CreateItem(props []string, response *Item) error {
	*response = Item{Title: props[0], Body: props[1]}
	return nil
}

func (a *API) AddItem(item Item, response *Item) error {
	db = append(db, item)
	*response = item
	return nil
}

func (a *API) EditItem(item Item, response *Item) error {
	for i, existing := range db {
		if existing.Title == item.Title {
			db[i] = Item{Title: item.Title, Body: item.Body}
			*response = db[i]
			break
		}
	}
	return nil
}

func (a *API) DeleteItem(item Item, response *Item) error {
	for i, existing := range db {
		if existing.Title == item.Title {
			db = append(db[:i], db[i+1:]...)
			break
		}
	}
	*response = item
	return nil
}

// Server enabling simple CRUD operation
func main() {
	// fmt.Println("Initial: ", db)
	// a := CreateItem("First", "First Item")
	// b := CreateItem("Second", "Second Item")
	// c := CreateItem("Third", "Third Item")

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)

	// fmt.Println("Three items A, B, C added: ", db)

	// DeleteItem(b)

	// fmt.Println("Item B deleted: ", db)

	// EditItem("Third", CreateItem("Second", "This is now the second Item"))

	// fmt.Println("Final database: ", db)

	// Register API
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		fmt.Println("Error registering the API: ", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		fmt.Println("Error listening: ", err)
	}

	log.Printf("Serving rpc on port %d", 4040)

	err = http.Serve(listener, nil)

	if err != nil {
		fmt.Println("Error serving rpc: ", err)
	}

}
