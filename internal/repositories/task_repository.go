package repositories

import (
	"database/sql"
	"errors"
	"simple-todo-app/internal/models"
)

type TaskRepository interface {
	GetAllByUser(userId uint) ([]models.Task, error)
	Create(task models.Task) error
	Update(id int, task models.Task) error
	Delete(id int) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) GetAllByUser(userId uint) ([]models.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, is_finished, user_id FROM tasks WHERE user_id = $1", userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks = []models.Task{}
	for rows.Next() {
		var each models.Task
		if err := rows.Scan(&each.Id, &each.Title, &each.Description, &each.IsFinished, &each.UserId); err != nil {
			return nil, err
		}

		tasks = append(tasks, each)
	}

	return tasks, nil
}

func (r *taskRepository) Create(task models.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, description, user_id) values ($1, $2, $3)", task.Title, task.Description, task.UserId)

	if err != nil {
		return errors.New("Error create task : " + err.Error())
	}

	return nil
}

func (r *taskRepository) Update(id int, data models.Task) error {
	res, err := r.db.Exec("UPDATE tasks SET title = $1, description = $2, is_finished = $3 WHERE id = $4", data.Title, data.Description, data.IsFinished, id)

	if err != nil {
		return errors.New("Error Update task : " + err.Error())
	}

	// optional. just check if data with id is exists
	rows, _ := res.RowsAffected()

	if rows == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *taskRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE from tasks where id = $1", id)

	if err != nil {
		return errors.New("Error Delete task : " + err.Error())
	}

	// optional. just check if data with id is exists
	rows, _ := res.RowsAffected()

	if rows == 0 {
		return errors.New("task not found")
	}

	return nil
}
