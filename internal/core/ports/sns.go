package ports

type SNSService interface {
	SendSMS(phone string, code string) error
}
