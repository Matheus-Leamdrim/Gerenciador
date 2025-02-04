package task

type Service struct {
    repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) CreateTask(task *Task) error {
    return s.repo.Create(task)
}

func (s *Service) GetTaskByID(id int) (*Task, error) {
    return s.repo.FindByID(id)
}

func (s *Service) UpdateTask(task *Task) error {
    return s.repo.Update(task)
}

func (s *Service) DeleteTask(id int) error {
    return s.repo.Delete(id)
}

func (s *Service) GetAllTasks() ([]*Task, error) {
    return s.repo.FindAll()
}