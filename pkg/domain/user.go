package domain

type Users struct{
	Name 		string `json:"id"`
	Email 		string `json:"email" gorm:"primarykey"`
	Phone 		string `json:"phone"`
	Password 	string `json:"password"`
}