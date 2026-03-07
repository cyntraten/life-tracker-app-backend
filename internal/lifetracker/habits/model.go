package habits

type Habit struct {
	ID              string `gorm:"primaryKey" json:"id"`
	Name            string `json:"name"`
	Streak          int    `json:"streak"`
	LastCompleted   int    `json:"lastCompleted"`
	ProgressPercent int8   `json:"progressPercent"`
}
