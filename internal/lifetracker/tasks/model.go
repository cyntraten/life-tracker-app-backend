package tasks

type Task struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	Timestamp int    `json:"timestamp"`
}
