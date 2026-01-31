package course

type NewCourseRequestPayload struct {
	UserPublicId string `db:"user_public_id"`
	Title        string `db:"title"`
	Images       string `db:"images"`
}