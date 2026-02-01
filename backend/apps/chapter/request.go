package chapter

type NewChapterRequestPayload struct {
	CoursePublicId   string `db:"course_public_id"`
	Title            string `db:"title"`
	PdfKey           string `db:"pdf_key"`
	OriginalFilename string `db:"original_filename"`
}