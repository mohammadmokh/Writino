package contract

type ImageStore interface {
	SaveImage([]byte, string) error
	DeleteImage(string) error
}
