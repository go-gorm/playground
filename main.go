package main

import (
	"fmt"

	"gorm.io/gen/examples/dal"
)

func main() {
	dal.ConnectDB(`sqlite.db`)
	generate()
	fmt.Println("vim-go")
}
