package task

type Repository interface {
    Create(task *Task) error
    FindByID(id int) (*Task, error)
    Update(task *Task) error
    Delete(id int) error
    FindAll() ([]*Task, error)
}