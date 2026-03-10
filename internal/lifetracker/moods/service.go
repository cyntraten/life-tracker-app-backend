package moods

type MoodsService interface {
	CreateMood(Mood Mood) (Mood, error)
	GetAllMoods() ([]Mood, error)
	GetMoodByID(id string) (Mood, error)
	UpdateMood(id string, Mood Mood) (Mood, error)
	DeleteMood(id string) error
}

type MoodService struct {
	repo MoodsRepository
}

func NewMoodsService(r MoodsRepository) MoodsService {
	return &MoodService{repo: r}
}

func (s *MoodService) CreateMood(mood Mood) (Mood, error) {
	if err := s.repo.CreateMood(mood); err != nil {
		return Mood{}, err
	}

	return mood, nil
}

func (s *MoodService) GetAllMoods() ([]Mood, error) {
	return s.repo.GetAllMoods()
}

func (s *MoodService) GetMoodByID(id string) (Mood, error) {
	return s.repo.GetMoodById(id)
}

func (s *MoodService) UpdateMood(id string, mood Mood) (Mood, error) {
	oldMood, err := s.repo.GetMoodById(id)
	if err != nil {
		return Mood{}, err
	}

	oldMood.ID = mood.ID
	oldMood.Note = mood.Note
	oldMood.Mood = mood.Mood
	oldMood.Timestamp = mood.Timestamp

	if err := s.repo.UpdateMood(oldMood); err != nil {
		return Mood{}, err
	}

	return oldMood, nil
}

func (s *MoodService) DeleteMood(id string) error {
	return s.repo.DeleteMood(id)
}
