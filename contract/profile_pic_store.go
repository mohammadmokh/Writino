package contract

type ProfilePicStore interface {
	SaveImage([]byte, string) error
	DeleteImage(string) error
}
