package domain

type Dentist struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"lastName"`
	RegistrationId string `json:"registrationId"`
}
