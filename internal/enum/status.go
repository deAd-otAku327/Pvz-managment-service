package enum

type ReceptionStatus int

const (
	StatusInProgress ReceptionStatus = iota
	StatusClose
)

var statusToName = map[ReceptionStatus]string{
	StatusInProgress: "in_progress",
	StatusClose:      "close",
}

var nameTostatus = map[string]ReceptionStatus{
	"in_progress": StatusInProgress,
	"close":       StatusClose,
}

func CheckStatus(city string) bool {
	_, ok := nameTostatus[city]
	return ok
}

func (s ReceptionStatus) String() string {
	return statusToName[s]
}
