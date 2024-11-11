package infrastructure

type IFileNumberRepository interface {
	Get() ([]int, error)
}
