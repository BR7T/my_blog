package structs

import "time"

type GetPost struct{
	ID int
	Title string
	Content string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

