package storage

import "Anastasia/skillfactory/databases/module30/pkg/storage/postgres"

type Interface interface {
	GetTasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
	UpdateTask(int, postgres.Task) error
	DeleteTask(int) error
}
