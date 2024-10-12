package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type TodoRepo struct {
	//mu sync.Mutex
	db                           *sql.DB
	insert, selectAll            *sql.Stmt
	selectByCompleted, deleteAll *sql.Stmt
	updateCompleted              *sql.Stmt
	updateCompletedByName        *sql.Stmt
	selectByName                 *sql.Stmt
}

const createTodoTable = `
CREATE TABLE IF NOT EXISTS todo (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE
);
`

func NewTodoRepo(path string) (*TodoRepo, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(createTodoTable); err != nil {
		return nil, err
	}
	insert, err := db.Prepare("INSERT INTO todo (name, description, completed) VALUES (?,?,?);")
	if err != nil {
		return nil, err
	}
	selectAll, err := db.Prepare("SELECT * FROM todo;")
	if err != nil {
		return nil, err
	}
	selectByCompleted, err := db.Prepare("SELECT * FROM todo WHERE completed=?;")
	if err != nil {
		return nil, err
	}

	deleteAll, err := db.Prepare("DELETE FROM todo;")
	if err != nil {
		return nil, err
	}

	updateCompleted, err := db.Prepare("UPDATE todo SET completed=TRUE WHERE id=?;")
	if err != nil {
		return nil, err
	}

	updateCompletedByName, err := db.Prepare("UPDATE todo SET completed=TRUE WHERE name=?;")
	if err != nil {
		return nil, err
	}

	selectByName, err := db.Prepare("SELECT * FROM todo WHERE name=?;")
	if err != nil {
		return nil, err
	}

	return &TodoRepo{
		db:                    db,
		insert:                insert,
		selectAll:             selectAll,
		selectByCompleted:     selectByCompleted,
		deleteAll:             deleteAll,
		updateCompleted:       updateCompleted,
		updateCompletedByName: updateCompletedByName,
		selectByName:          selectByName,
	}, nil
}

func (r *TodoRepo) Insert(todo Todo) (int, error) {
	res, err := r.insert.Exec(todo.Name, todo.Description, todo.Completed)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *TodoRepo) SelectAll() ([]Todo, error) {
	rows, err := r.selectAll.Query()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	data := make([]Todo, 0)
	for rows.Next() {
		t := Todo{}
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Completed); err != nil {
			return nil, err
		}
		data = append(data, t)
	}
	return data, nil
}

func (r *TodoRepo) SelectByCompletedStatus(completed bool) ([]Todo, error) {
	rows, err := r.selectByCompleted.Query(completed)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	data := make([]Todo, 0)
	for rows.Next() {
		t := Todo{}
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Completed); err != nil {
			return nil, err
		}
		data = append(data, t)
	}
	return data, nil
}

func (r *TodoRepo) Clear() error {
	_, err := r.deleteAll.Exec()
	return err
}

func (r *TodoRepo) Complete(todoID int) error {
	_, err := r.updateCompleted.Exec(todoID)
	return err
}

func (r *TodoRepo) CompleteByName(name string) error {
	_, err := r.updateCompletedByName.Exec(name)
	return err
}

func (r *TodoRepo) GetByName(name string) (Todo, error) {
	todo := Todo{}
	err := r.selectByName.QueryRow(name).Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Completed)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Todo{}, nil
		}
		return Todo{}, err
	}

	return todo, nil
}
