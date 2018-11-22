package persistance

type DatabaseHandler interface {
	AddPhoto(Photo) ([]byte, error)
	FindPhoto([]byte) (Photo, error)
}
