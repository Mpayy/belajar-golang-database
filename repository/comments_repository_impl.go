package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error) {
	//TODO implement me
	query := "INSERT INTO comments(email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(lastId)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comments, error) {
	//TODO implement me
	query := "SELECT id, email, comment FROM comments WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	comment := entity.Comments{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comment, err
		}
		return comment, nil
	}
	return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comments, error) {
	//TODO implement me
	query := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comments
	for rows.Next() {
		comment := entity.Comments{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
