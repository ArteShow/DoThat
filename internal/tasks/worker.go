package tasks

import (
	"context"
	"time"

	"github.com/ArteShow/DoThat/internal/database/repositories/tasks"
)

func (r *TaskRegistry) sync() error {
	dbTasks, err := r.TaskRepository.GetAll()
	if err != nil {
		return err
	}

	for _, task := range r.Tasks {
		var dbTask *tasks.Task

		for i := range dbTasks {
			if dbTasks[i].ID == task.ID {
				dbTask = &dbTasks[i]
				break
			}
		}

		if dbTask == nil {
			if err = r.TaskRepository.CreateTask(task); err != nil {
				return err
			}

			continue
		}

		if task.Deadline != dbTask.Deadline {
			if err = r.TaskRepository.UpdateDeadline(task.Deadline, task.ID); err != nil {
				return err
			}
		}

		if task.Status != dbTask.Status {
			if err = r.TaskRepository.UpdateStatus(task.ID, task.Status); err != nil {
				return err
			}
		}
	}

	for _, dbTask := range dbTasks {
		if _, exists := r.Tasks[dbTask.ID]; !exists {
			if err = r.TaskRepository.DeleteTask(dbTask.ID); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *TaskRegistry) StartSync(ctx context.Context) error {
	ticker := time.NewTicker(time.Minute * 2)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := r.sync(); err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}
