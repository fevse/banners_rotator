package storage

type Banner struct {
	ID          int
	Description string
}

type Slot struct {
	ID          int
	Description string
}

type Group struct {
	ID          int
	Description string
}

type Storage struct {
	ID       int
	SlotID   int
	BannerID int
	GroupID  int
	View     int
	Click    int
	Rewards  float64
}
