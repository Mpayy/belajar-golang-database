package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentRepositoryImpl_Insert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comments{
		Email:   "repository@test.com",
		Comment: "Test CommentRepository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentRepositoryImpl_FindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()

	id := int32(100)
	comments, err := commentRepository.FindById(ctx, id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(comments)
}

func TestCommentRepositoryImpl_FindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
