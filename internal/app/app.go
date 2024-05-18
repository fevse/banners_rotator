package app

type App struct {
	Storage Storage
}

type Storage interface {
	Add(int, int) error
	Delete(int, int) error
	Click(int, int, int) error
	Choose(int, int) (int, error)
}

func New(storage Storage) *App {
	return &App{
		Storage: storage,
	}
}
