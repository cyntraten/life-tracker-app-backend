package habits

type HabitsService interface {
	CreateHabit(habit Habit) (Habit, error)
	GetAllHabits() ([]Habit, error)
	GetHabitByID(id string) (Habit, error)
	UpdateHabit(id string, habit Habit) (Habit, error)
	DeleteHabit(id string) error
}

type HabitService struct {
	repo HabitsRepository
}

func NewHabitsService(r HabitsRepository) HabitsService {
	return &HabitService{repo: r}
}

func (s *HabitService) CreateHabit(habit Habit) (Habit, error) {
	if err := s.repo.CreateHabit(habit); err != nil {
		return Habit{}, err
	}

	return habit, nil
}

func (s *HabitService) GetAllHabits() ([]Habit, error) {
	return s.repo.GetAllHabits()
}

func (s *HabitService) GetHabitByID(id string) (Habit, error) {
	return s.repo.GetHabitById(id)
}

func (s *HabitService) UpdateHabit(id string, habit Habit) (Habit, error) {
	oldHabit, err := s.repo.GetHabitById(id)
	if err != nil {
		return Habit{}, err
	}

	oldHabit.ID = habit.ID
	oldHabit.Name = habit.Name
	oldHabit.Streak = habit.Streak
	oldHabit.LastCompleted = habit.LastCompleted
	oldHabit.ProgressPercent = habit.ProgressPercent

	if err := s.repo.UpdateHabit(oldHabit); err != nil {
		return Habit{}, err
	}

	return oldHabit, nil
}

func (s *HabitService) DeleteHabit(id string) error {
	return s.repo.DeleteHabit(id)
}
