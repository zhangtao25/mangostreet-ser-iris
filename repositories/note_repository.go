package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
	"mangostreet-ser-iris/model"
)

var NoteRepository = newNoteRepository()

func newNoteRepository() *noteRepository {
	return &noteRepository{}
}

type noteRepository struct {
}

func (this *noteRepository) Get(db *gorm.DB, id int64) *model.Note {
	ret := &model.Note{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *noteRepository) Create(db *gorm.DB, t *model.Note) (err error) {
	err = db.Create(t).Error
	return
}

func (this *noteRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Note) {
	cnd.Find(db, &list)
	return
}