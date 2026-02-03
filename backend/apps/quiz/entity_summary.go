package quiz

import (
	"time"

	"github.com/google/uuid"
)

type SummaryEntity struct {
	Id              int       `db:"id"`
	SummaryPublicId uuid.UUID    `db:"summary_public_id"`
	ChapterPublicId uuid.UUID `db:"chapter_public_id"`
	SummaryJSON string `db:"summary_json"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}