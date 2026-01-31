package chapter

type NewChapterRequestPayload struct {
	CoursePublicId string `db:"course_public_id"`
	Title          string `db:"title"`
	PdfUrl         string `db:"pdf_url"`
}