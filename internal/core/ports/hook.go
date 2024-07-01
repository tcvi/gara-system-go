package ports

type HookService interface {
	Send(mess string) error
}
