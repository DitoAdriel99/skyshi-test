package repositories

import (
	"database/sql"
	"fmt"
	"go-learn/config"
	"go-learn/entities"
)

type TodoRepo struct {
	conn *sql.DB
}

func NewTodoRepositories() *TodoRepo {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &TodoRepo{
		conn: conn,
	}
}

func (x *TodoRepo) Create(todo *entities.Todo) error {
	query := `INSERT INTO todo (title, activity_group_id,isActive ,priority, updatedAt, createdAt) VALUES (?,?,?,?,?,?)`

	res, err := x.conn.Exec(query, todo.Title, todo.ActivityID, todo.IsActive, todo.Priority, todo.UpdatedAt, todo.CreatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	todo.ID, err = res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("getting last inserted id: %w", err)
		return err
	}

	return nil
}

func (x *TodoRepo) GetAll() ([]entities.Todo, error) {

	rows, err := x.conn.Query(`SELECT * FROM todo`)
	if err != nil {
		err = fmt.Errorf("error querying: %w", err)
		return nil, err
	}
	var todoObjects []entities.Todo

	for rows.Next() {
		var todoObject entities.Todo

		err = rows.Scan(
			&todoObject.ID,
			&todoObject.Title,
			&todoObject.ActivityID,
			&todoObject.IsActive,
			&todoObject.Priority,
			&todoObject.UpdatedAt,
			&todoObject.CreatedAt,
		)

		if err != nil {
			err = fmt.Errorf("error scanning: %w", err)
			return nil, err
		}
		todoObjects = append(todoObjects, todoObject)
	}
	return todoObjects, nil
}

func (x *TodoRepo) FindById(id int64) (*entities.Todo, error) {
	query := `SELECT * FROM todo WHERE id = ?`

	var todoObject entities.Todo

	err := x.conn.QueryRow(query, id).Scan(
		&todoObject.ID,
		&todoObject.Title,
		&todoObject.ActivityID,
		&todoObject.IsActive,
		&todoObject.Priority,
		&todoObject.UpdatedAt,
		&todoObject.CreatedAt,
	)

	if err != nil {
		err = fmt.Errorf("scanning todo objects: %w", err)

		return nil, err
	}

	return &todoObject, nil
}

func (x *TodoRepo) UpdateTodo(todo *entities.Todo) (*entities.Todo, error) {
	query := `
		UPDATE todo
		SET 
			title = ?,
			activity_group_id = ?,
			isActive = ?,
			priority = ?,
			updatedAt = ?,
			createdAt = ?
		WHERE id = ?	
	`

	_, err := x.conn.Exec(query,
		todo.Title,
		todo.ActivityID,
		todo.IsActive,
		todo.Priority,
		todo.UpdatedAt,
		todo.CreatedAt,
		todo.ID,
	)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)

		return nil, err
	}

	return todo, nil
}

func (x *TodoRepo) DeleteTodo(id int64) error {
	query := `DELETE FROM todo WHERE id = ?`

	_, err := x.conn.Exec(query, id)
	if err != nil {
		err = fmt.Errorf("executing query delete: %w", err)
		return err
	}
	return nil
}
