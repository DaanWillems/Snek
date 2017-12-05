// Package models contains the types for schema '23558'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Task represents a row from '23558.tasks'.
type Task struct {
	ID           int            `json:"id"`           // id
	Creatoruuid  string         `json:"creatorUUID"`  // creatorUUID
	Assigneduuid sql.NullString `json:"assignedUUID"` // assignedUUID
	Solved       bool           `json:"solved"`       // solved
	Title        string         `json:"title"`        // title
	Description  string         `json:"description"`  // description
	Location     string         `json:"location"`     // location

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Task exists in the database.
func (t *Task) Exists() bool {
	return t._exists
}

// Deleted provides information if the Task has been deleted from the database.
func (t *Task) Deleted() bool {
	return t._deleted
}

// Insert inserts the Task to the database.
func (t *Task) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO 23558.tasks (` +
		`creatorUUID, assignedUUID, solved, title, description, location` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, t.Creatoruuid, t.Assigneduuid, t.Solved, t.Title, t.Description, t.Location)
	res, err := db.Exec(sqlstr, t.Creatoruuid, t.Assigneduuid, t.Solved, t.Title, t.Description, t.Location)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	t.ID = int(id)
	t._exists = true

	return nil
}

// Update updates the Task in the database.
func (t *Task) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE 23558.tasks SET ` +
		`creatorUUID = ?, assignedUUID = ?, solved = ?, title = ?, description = ?, location = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, t.Creatoruuid, t.Assigneduuid, t.Solved, t.Title, t.Description, t.Location, t.ID)
	_, err = db.Exec(sqlstr, t.Creatoruuid, t.Assigneduuid, t.Solved, t.Title, t.Description, t.Location, t.ID)
	return err
}

// Save saves the Task to the database.
func (t *Task) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Delete deletes the Task from the database.
func (t *Task) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM 23558.tasks WHERE id = ?`

	// run query
	XOLog(sqlstr, t.ID)
	_, err = db.Exec(sqlstr, t.ID)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// TaskByID retrieves a row from '23558.tasks' as a Task.
//
// Generated from index 'tasks_id_pkey'.
func TaskByID(db XODB, id int) (*Task, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, creatorUUID, assignedUUID, solved, title, description, location ` +
		`FROM tasks ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	t := Task{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&t.ID, &t.Creatoruuid, &t.Assigneduuid, &t.Solved, &t.Title, &t.Description, &t.Location)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func TaskAll(db XODB) ([]Task, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, creatorUUID, assignedUUID, solved, title, description, location ` +
		`FROM tasks `

	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	var tasks []Task

	for q.Next() {
		t := Task{}
		// scan
		err = q.Scan(&t.ID, &t.Creatoruuid, &t.Assigneduuid, &t.Solved, &t.Title, &t.Description, &t.Location)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}