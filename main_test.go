package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	comment := &Comment{
		Description: "Some comment",
		Files: []File{
			{Name: "file-1", DataURL: "htpp://file-1"},
			{Name: "file-2", DataURL: "htpp://file-2"},
		},
	}

	DB.Create(comment)

	var result Comment
	if err := DB.First(&result, comment.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	//Delete comment, it will delete associated files and should trigger BeforeDelete.
	//Files deleted successfully but BeforeDelete hook is not triggered.
	commentId := comment.ID
	DB.Delete(&Comment{}, commentId)

	// To confirm file has deleted or not
	var files []File
	DB.Where("comment_id = ?", commentId).Find(&files)
	if len(files) == 0 {
		fmt.Println("Associated files has deleted")
	}
}
