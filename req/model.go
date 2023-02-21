package req

type ApplyApartment struct {
	StudentID   int64
	ApartmentID int64
	RoomID      int64
	BeginTime   string
	Duration    int64
}

type ApartmentInfoResp struct {
	ApartmentID int64
	RoomID      int64
	Roommate    []string
	AdminName   string
	Deadline    string
}
