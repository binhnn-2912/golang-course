package domain

import "time"

// Course represents a course entity.
type Course struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Instructor  string    `json:"instructor"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateCourseRequest is the payload for creating a course.
type CreateCourseRequest struct {
	Title       string  `json:"title"       validate:"required,min=3,max=150"`
	Description string  `json:"description" validate:"required"`
	Instructor  string  `json:"instructor"  validate:"required"`
	Price       float64 `json:"price"       validate:"gte=0"`
}

// UpdateCourseRequest is the payload for updating a course.
type UpdateCourseRequest struct {
	Title       *string  `json:"title,omitempty"       validate:"omitempty,min=3,max=150"`
	Description *string  `json:"description,omitempty"`
	Instructor  *string  `json:"instructor,omitempty"`
	Price       *float64 `json:"price,omitempty"       validate:"omitempty,gte=0"`
}
