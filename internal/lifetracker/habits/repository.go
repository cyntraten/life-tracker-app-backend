package habits

import "gorm.io/gorm"

type HabitsRepository interface {
	GetAllHabits() ([]Habit, error)
	CreateHabit(habit Habit) error
	DeleteHabit(id string) error
	UpdateHabit(habit Habit) error
	GetHabitById(id string) (Habit, error)
}

type HabitsRepo struct {
	db *gorm.DB
}

func NewHabitsRepository(db *gorm.DB) HabitsRepository {
	return &HabitsRepo{db: db}
}

func (r *HabitsRepo) GetAllHabits() ([]Habit, error) {
	var habits []Habit
	err := r.db.Find(&habits).Error
	return habits, err
}

func (r *HabitsRepo) CreateHabit(habit Habit) error {
	return r.db.Create(&habit).Error
}

func (r *HabitsRepo) DeleteHabit(id string) error {
	return r.db.Delete(&Habit{}, "id = ?", id).Error
}

func (r *HabitsRepo) UpdateHabit(Habit Habit) error {
	return r.db.Save(&Habit).Error
}

func (r *HabitsRepo) GetHabitById(id string) (Habit, error) {
	var habit Habit

	err := r.db.First(&Habit{}, "id = ?", id).Error

	return habit, err
}
