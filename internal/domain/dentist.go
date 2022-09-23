package domain

type Dentist struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	RegistrationId string `json:"registration_id"`
}
