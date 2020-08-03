package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestCreateData(t *testing.T) {

	users, _, _ := createSomeTestdata()

	var result User
	if err := DB.First(&result, users[0].ID).Error; err != nil {

		t.Errorf("Failed, got error: %v", err)
	}
}

func TestReadData(t *testing.T) {

	users, _, _ := createSomeTestdata()

	queryConnection := DB
	queryConnection = queryConnection.Joins("left join document_fulltexts on documents.id = document_fulltexts.document_id")
	queryConnection = queryConnection.Where(`"documents"."user_id" = ?`, users[0].ID).Order(`created_at desc`)

	queryConnection = queryConnection.Select(
		`"documents"."user_id",
		"documents"."name",
		CASE WHEN "document_fulltexts"."fulltext" IS NOT NULL THEN TRUE ELSE FALSE END AS "fulltext_exists"`,
	)

	documents := []DocumentListEntry{}

	queryConnection.Table("documents").Scan(&documents)

	if len(documents) != 3 {
		t.Errorf("Failed to get all documents from the DB")
	}
}

func createSomeTestdata() ([]User, []Document, []DocumentFulltext) {

	var users []User
	var documents []Document
	var documentFulltexts []DocumentFulltext

	user1 := User{Name: "user1"}

	DB.Create(&user1)
	document1 := Document{
		UserID: user1.ID,
		Name:   "document1",
	}
	DB.Create(&document1)

	document2 := Document{
		UserID: user1.ID,
		Name:   "document2",
	}
	DB.Create(&document2)

	document3 := Document{
		UserID: user1.ID,
		Name:   "document3",
	}
	DB.Create(&document3)

	documentFulltext1 := DocumentFulltext{
		DocumentID: document1.ID,
		Name:       "documentFulltext1",
	}
	DB.Create(&documentFulltext1)
	documentFulltext3 := DocumentFulltext{
		DocumentID: document3.ID,
		Name:       "documentFulltext3",
	}
	DB.Create(&documentFulltext3)

	users = append(users, user1)
	documents = append(documents, document1)
	documents = append(documents, document2)
	documents = append(documents, document3)
	documentFulltexts = append(documentFulltexts, documentFulltext1)
	documentFulltexts = append(documentFulltexts, documentFulltext3)

	return users, documents, documentFulltexts
}
