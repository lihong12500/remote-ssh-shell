package command

type ICli interface {
	Run() (string, error)
}
