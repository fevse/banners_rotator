package app

type App struct {
	Storage Storage
	Logger  Logger
	Rabbit  Rabbit
}

type Storage interface {
	AddBannerToSlot(int, int) error
	DeleteBannerFromSlot(int, int) error
	ClickBanner(int, int, int) error
	ChooseBannerToShow(int, int) (int, error)
}

type Logger interface {
	Info(string)
	Error(string)
}

type Rabbit interface {
	Publish(string) error
}

func New(storage Storage, logger Logger, rabbit Rabbit) *App {
	return &App{
		Storage: storage,
		Logger:  logger,
		Rabbit:  rabbit,
	}
}
