package moods

type Mood struct {
	ID        string  `gorm:"primaryKey" json:"id"`
	Note      *string `json:"done"`
	Mood      int8    `json:"mood"`
	Timestamp int     `json:"timestamp"`
}
