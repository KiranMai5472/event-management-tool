package models

// User used to get the user information
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
} //name @User

// UpdateUser is model used to update the user data
type UserSignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
} //name @UserSignUp

// LoginSuccess is used to give response when user login successfully
type LoginSuccess struct {
	Status   string `json:"status"`
	Message  string `json:"Message"`
	Username string `json:"username"`
	//FullName string    `json:"fullname"`
	Token string `json:"token"`
	//Expiry   time.Time `json:"expiry"`
	//UserID   uint      `json:"userId"`
} //name @LoginSuccess

// Users used for getting all users data
type Users struct {
	UserID     string `gorm:"column:id" json:"userId"`
	UserName   string `gorm:"column:username" json:"username"`
	FirstName  string `gorm:"column:first_name" json:"firstName"`
	LastName   string `gorm:"column:last_name" json:"lastName"`
	Email      string `gorm:"column:email" json:"email"`
	Password   string `gorm:"column:password" json:"password"`
	Phone      string `gorm:"column:mobile_no" json:"phone"`
	EmployeeID string `gorm:"column:employee_id" json:"employeeId"`
	Department string `gorm:"column:department" json:"department"`
	Title      string `json:"title"`
} //name @Users

// GetUserById used to get the user information
type GetUserById struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FirstName  string `gorm:"column:first_name" json:"firstName"`
	LastName   string `gorm:"column:last_name" json:"lastName"`
	Email      string `json:"email"`
	RoleId     uint   `json:"role_id"`
	Active     uint   `json:"active"`
	Department string `json:"department"`
	EmployeeId uint   `gorm:"column:employee_id" json:"employeeId"`
	MobileNo   uint   `gorm:"column:mobile_no" json:"phone"`
} //name @GetUserById

type userLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
} //name @userLogin
