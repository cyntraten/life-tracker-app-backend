package moods

import "gorm.io/gorm"

type MoodsRepository interface {
	GetAllMoods() ([]Mood, error)
	CreateMood(mood Mood) error
	DeleteMood(id string) error
	UpdateMood(mood Mood) error
	GetMoodById(id string) (Mood, error)
}

type MoodsRepo struct {
	db *gorm.DB
}

func NewMoodsRepository(db *gorm.DB) MoodsRepository {
	return &MoodsRepo{db: db}
}

func (r *MoodsRepo) GetAllMoods() ([]Mood, error) {
	var moods []Mood
	err := r.db.Find(&moods).Error
	return moods, err
}

func (r *MoodsRepo) CreateMood(mood Mood) error {
	return r.db.Create(&mood).Error
}

func (r *MoodsRepo) DeleteMood(id string) error {
	return r.db.Delete(&Mood{}, "id = ?", id).Error
}

func (r *MoodsRepo) UpdateMood(mood Mood) error {
	return r.db.Save(&mood).Error
}

func (r *MoodsRepo) GetMoodById(id string) (Mood, error) {
	var mood Mood

	err := r.db.First(&Mood{}, "id = ?", id).Error

	return mood, err
}
