package snsservice

func (s *SnsService) SendSMS(phone string, code string) error {
	//message := fmt.Sprintf("Your code is %s ", code)
	//message := code
	//
	//input := &sns.PublishInput{
	//	Message:     aws.String(message),
	//	PhoneNumber: aws.String(phone),
	//	MessageAttributes: map[string]types.MessageAttributeValue{
	//		"AWS.SNS.SMS.SMSType": {
	//			DataType:    aws.String("String"),
	//			StringValue: aws.String("Transactional"), // or "Promotional"
	//		},
	//	},
	//}
	//
	//_, err := s.client.Publish(context.TODO(), input)
	//if err != nil {
	//	return err
	//}
	return nil
}
