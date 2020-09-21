package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
	"time"
)

type Entity struct {
	Id int `gorm:"unique"`
}

func main() {

	//setup the db
	dialect := "sqlite3"
	name := "db.db?cache=shared&_journal=WAL"
	db, err := gorm.Open(dialect, name)
	if err != nil {
		fmt.Println(err)
	}
	//migrate
	db.AutoMigrate(&Entity{})

	//to simulate 2 different requests/ processes working on db with different transactions,
	//this part starts a goroutine with a non wrong managed transaction that lock the db
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go wrongTransaction(db, wg)
	wg.Wait()

	//after the wrong transaction, this code try to read and write in the db
	//using another transaction, but find that db is locked
	ticker := time.NewTicker(1 * time.Second)
	for {

		select {
		case <-ticker.C:
			t2 := db.Begin()

			entity2 := &Entity{Id: 2}
			var users2 []Entity
			err = t2.Find(&users2).Error
			fmt.Print(err)
			err =t2.Create(entity2).Error
			fmt.Print(err)
			err =t2.Commit().Error
			fmt.Print(err)
		}

	}

}

func wrongTransaction(db *gorm.DB, wg *sync.WaitGroup) {

	//the transaction is open, but not closed with commit nor rollback
	//in this case the db is locked and to unlock it the process must be restarted
	t := db.Begin()

	entity := &Entity{Id: 1}
	var entities []Entity
	err := t.Find(&entities).Error
	fmt.Print(err)

	err = t.Create(entity).Error
	fmt.Print(err)
	wg.Done()
}
