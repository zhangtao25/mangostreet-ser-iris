package repositories

import (
	"github.com/jinzhu/gorm"
	"mangostreet-ser-iris/model"
)

var NoteRepository = newNoteRepository()

func newNoteRepository() *noteRepository {
	return &noteRepository{}
}

type noteRepository struct {
}

func (this *noteRepository) Create(db *gorm.DB, t *model.Note) (err error) {
	err = db.Create(t).Error
	return
}