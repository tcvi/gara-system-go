package domain

type Notification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type PushNotificationReq struct {
	Tokens []string     `json:"tokens"`
	Data   Notification `json:"data"`
}
