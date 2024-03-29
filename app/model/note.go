package model

import (
	"fmt"
	"time"

	"github.com/josephspurrier/gowebapp/app/shared/database"

	"gopkg.in/mgo.v2/bson"
)

// *****************************************************************************
// Note
// *****************************************************************************

// Note table contains the information for each note
type Note struct {
	ObjectID  bson.ObjectId `bson:"_id"`
	ID        uint32        `db:"id" bson:"id,omitempty"` // Don't use Id, use NoteID() instead for consistency with MongoDB
	Content   string        `db:"content" bson:"content"`
	UserID    bson.ObjectId `bson:"user_id"`
	UID       uint32        `db:"user_id" bson:"userid,omitempty"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	Deleted   uint8         `db:"deleted" bson:"deleted"`
}

// NoteID returns the note id
func (u *Note) NoteID() string {
	r := ""
	r = fmt.Sprintf("%v", u.ID)
	return r
}

// NoteByID gets note by ID
func NoteByID(userID string, noteID string) (Note, error) {
	var err error
	result := Note{}
	err = database.SQL.Get(&result, "SELECT id, content, user_id, created_at, updated_at, deleted FROM note WHERE id = ? AND user_id = ? LIMIT 1", noteID, userID)
	return result, standardizeError(err)
}

// NotesByUserID gets all notes for a user
func NotesByUserID(userID string) ([]Note, error) {
	var err error
	var result []Note
	err = database.SQL.Select(&result, "SELECT id, content, user_id, created_at, updated_at, deleted FROM note WHERE user_id = ?", userID)
	return result, standardizeError(err)
}

// NoteCreate creates a note
func NoteCreate(content string, userID string) error {
	var err error
	_, err = database.SQL.Exec("INSERT INTO note (content, user_id) VALUES (?,?)", content, userID)
	return standardizeError(err)
}

// NoteUpdate updates a note
func NoteUpdate(content string, userID string, noteID string) error {
	var err error
	_, err = database.SQL.Exec("UPDATE note SET content=? WHERE id = ? AND user_id = ? LIMIT 1", content, noteID, userID)
	return standardizeError(err)
}

// NoteDelete deletes a note
func NoteDelete(userID string, noteID string) error {
	var err error
	_, err = database.SQL.Exec("DELETE FROM note WHERE id = ? AND user_id = ?", noteID, userID)
	return standardizeError(err)
}
