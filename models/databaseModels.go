package models

type Questions struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	Status_id   uint   `json:"status_id"`
	Subject_id  uint   `json:"subject_id"`
}

type Statuses struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Subjects struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
