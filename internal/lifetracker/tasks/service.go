package tasks

type TasksService interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TasksRepository
}

func NewTasksService(r TasksRepository) TasksService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(task Task) (Task, error) {
	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTask(id string, task Task) (Task, error) {
	oldTask, err := s.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}

	oldTask.ID = task.ID
	oldTask.Title = task.Title
	oldTask.Done = task.Done
	oldTask.Timestamp = task.Timestamp

	if err := s.repo.UpdateTask(oldTask); err != nil {
		return Task{}, err
	}

	return oldTask, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
