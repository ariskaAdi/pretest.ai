package course

import (
	"time"

	"github.com/google/uuid"
)

type CourseEntity struct {
	Id             	int 		`db:"id"`
	CoursePublicId 	uuid.UUID 	`db:"course_public_id"`
	UserPublicId 	string 		`db:"user_public_id"`
	Title           string 		`db:"title"`
	Images 			string 		`db:"images"`
	CreatedAt 		time.Time 	`db:"created_at"`
	UpdatedAt 		time.Time 	`db:"updated_at"`
}

func NewFormCreateCourseRequest(req NewCourseRequestPayload) CourseEntity {
	return CourseEntity{
		CoursePublicId: uuid.New(),
		UserPublicId: req.UserPublicId,
		Title: req.Title,
		Images: req.Images,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}