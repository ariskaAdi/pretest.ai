package chapter

import (
	"time"

	"github.com/google/uuid"
)

type ChapterEntity struct {
	Id              int       `db:"id"`
	ChapterPublicId uuid.UUID `db:"chapter_public_id"`
	CoursePublicId  string    `db:"course_public_id"`
	Title           string    `db:"title"`
	PdfUrl          string    `db:"pdf_url"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt 		time.Time `db:"updated_at"`
}

func NewFormCreateChapterRequest(req NewChapterRequestPayload) ChapterEntity {
	return ChapterEntity{
		ChapterPublicId: uuid.New(),
		CoursePublicId: req.CoursePublicId,
		Title: req.Title,
		PdfUrl: req.PdfUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}