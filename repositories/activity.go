package repositories

import (
	"database/sql"
	"fmt"
	"go-learn/config"
	"go-learn/entities"
)

type ActivityRepo struct {
	conn *sql.DB
}

func NewActivityRepositories() *ActivityRepo {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &ActivityRepo{
		conn: conn,
	}
}

func (c *ActivityRepo) Create(activity *entities.Activity) error {
	query := `INSERT INTO activity (title, email, updatedAt, createdAt) VALUES (?,?,?,?)`

	res, err := c.conn.Exec(query, activity.Title, activity.Email, activity.UpdatedAt, activity.CreatedAt)

	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	activity.ID, err = res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("getting last inserted id: %w", err)
		return err
	}

	return nil
}

func (c *ActivityRepo) GetAll() ([]entities.Activity, error) {

	rows, err := c.conn.Query(`SELECT id, title, email, updatedAt, createdAt FROM activity`)
	if err != nil {
		err = fmt.Errorf("error querying: %w", err)
		return nil, err
	}
	var activityObjects []entities.Activity

	for rows.Next() {
		var activityObject entities.Activity

		err = rows.Scan(
			&activityObject.ID,
			&activityObject.Title,
			&activityObject.Email,
			&activityObject.UpdatedAt,
			&activityObject.CreatedAt,
		)

		if err != nil {
			err = fmt.Errorf("error scanning: %w", err)
			return nil, err
		}
		activityObjects = append(activityObjects, activityObject)
	}
	return activityObjects, nil
}

func (c *ActivityRepo) FindById(id int64) (*entities.Activity, error) {
	query := `SELECT * FROM activity WHERE id = ?`

	var activityObject entities.Activity

	err := c.conn.QueryRow(query, id).Scan(
		&activityObject.ID,
		&activityObject.Title,
		&activityObject.Email,
		&activityObject.UpdatedAt,
		&activityObject.CreatedAt,
	)

	if err != nil {
		err = fmt.Errorf("scanning activity objects: %w", err)

		return nil, err
	}

	return &activityObject, nil
}

func (c *ActivityRepo) UpdateActivity(activity *entities.Activity) (*entities.Activity, error) {
	query := `
		UPDATE activity
		SET 
			title = ?,
			email = ?,
			updatedAt = ?,
			createdAt = ?
		WHERE id = ?	
	`

	_, err := c.conn.Exec(query,
		activity.Title,
		activity.Email,
		activity.UpdatedAt,
		activity.CreatedAt,
		activity.ID,
	)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)

		return nil, err
	}

	return activity, nil
}

func (c *ActivityRepo) DeleteActivity(id int64) error {
	query := `DELETE FROM activity WHERE id = ?`

	_, err := c.conn.Exec(query, id)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)
		return err
	}
	return nil
}
