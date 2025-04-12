package enum

type ReceptionStatus int

const (
	InProgress ReceptionStatus = iota
	Close
)

var statusToName = map[ReceptionStatus]string{
	InProgress: "in_progress",
	Close:      "close",
}

var nameTostatus = map[string]ReceptionStatus{
	"in_progress": InProgress,
	"close":       Close,
}

func CheckStatus(city string) bool {
	_, ok := nameTostatus[city]
	return ok
}

func (s ReceptionStatus) String() string {
	return statusToName[s]
}
