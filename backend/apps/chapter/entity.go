package chapter

import (
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type ChapterEntity struct {
	Id              int       `db:"id"`
	ChapterPublicId uuid.UUID `db:"chapter_public_id"`
	CoursePublicId  string    `db:"course_public_id"`
	Title           string    `db:"title"`
	PdfKey          string    `db:"pdf_key"`
	OriginalFilename string    `db:"original_filename"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt 		time.Time `db:"updated_at"`
}

func NewFormCreateChapterRequest(req NewChapterRequestPayload) ChapterEntity {
	return ChapterEntity{
		ChapterPublicId: uuid.New(),
		CoursePublicId: req.CoursePublicId,
		Title: req.Title,
		PdfKey: req.PdfKey,
		OriginalFilename: req.OriginalFilename,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func GenerateUniqueKey(originalName string) string {
    ext := filepath.Ext(originalName) 
    uniqueID := uuid.New().String()    
    return uniqueID + ext             
}