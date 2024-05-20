package app

type App struct {
	Storage Storage
	Logger Logger
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

func New(storage Storage, logger Logger) *App {
	return &App{
		Storage: storage,
		Logger: logger,
	}
}
