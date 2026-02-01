package course

import (
	"time"

	"github.com/google/uuid"
)

type CourseEntity struct {
	Id             	int 		`db:"id"`
	CoursePublicId 	uuid.UUID 	`db:"course_public_id"`
	UserPublicId 	uuid.UUID 		`db:"user_public_id"`
	Title           string 		`db:"title"`
	ImagesUrl 		string 		`db:"images_url"`
	CreatedAt 		time.Time 	`db:"created_at"`
	UpdatedAt 		time.Time 	`db:"updated_at"`
}

func NewFormCreateCourseRequest(req NewCourseRequestPayload) CourseEntity {
	return CourseEntity{
		CoursePublicId: uuid.New(),
		UserPublicId: req.UserPublicId,
		Title: req.Title,
		ImagesUrl: req.ImagesUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}