package course

import "github.com/google/uuid"

type NewCourseRequestPayload struct {
	UserPublicId   uuid.UUID `db:"user_public_id"`
	Title        string `db:"title"`
	ImagesUrl    string `db:"images_url"`
}