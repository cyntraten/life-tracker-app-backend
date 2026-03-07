package database

type Task struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	Timestamp int    `json:"timestamp"`
}

type Mood struct {
	ID        string  `gorm:"primaryKey" json:"id"`
	Note      *string `json:"done"`
	Mood      int8    `json:"mood"`
	Timestamp int     `json:"timestamp"`
}

type Habit struct {
	ID              string `gorm:"primaryKey" json:"id"`
	Name            string `json:"name"`
	Streak          int    `json:"streak"`
	LastCompleted   int    `json:"lastCompleted"`
	ProgressPercent int8   `json:"progressPercent"`
}
