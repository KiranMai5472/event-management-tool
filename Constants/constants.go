package Constants

const (

	// status for the output
	Failed        string = "Failed"
	Success       string = "Success"
	Logout        string = "Logged Out Successfully"
	LogoutFailed  string = "Logout failed"
	Authorization string = "token"
	Status        string = "status"
	Message       string = "message"
	Code          string = "code"
	Data          string = "Data"
	Error         string = "error"
	Username      string = "username"
	Password      string = "password"
	ExpiryTime    string = "expiryTime"
	Token         string = "token"
	CacheEnabled  string = "ENABLED"
	EmptyString   string = ""

	//create user constants
	UnableToCreateUser    string = "Unable to Create User... Please try again"
	CreatedSuccessfully   string = "Created Successfully"
	MissingFields         string = "Missing Fields"
	UnableToUpdateContent string = "Unable to update content"

	//getUser constants
	InvalidExpiredToken     string = "Invalid or expired token"
	UnableTOFetchUser       string = "Unable to fetch users by id"
	UserFetchedSuccessfully string = "user Fetched successfully"
	RecordNotFound          string = "Record Not Found"
	UnableToUpdate          string = "Unable To Update"
	UserUpdatedSuccessfully string = "User Updated Successfully"
	UnableTOFetchData       string = "Unable to fetch users data"

	//getToken
	FailedToCreateToken string = "failed to create token String"
	InvalidToken        string = "Invalid Token failed to get the claims"
	MarshaingError      string = "Error while performing marshaing"
	UnmarshalError      string = "Error while performing unmarshaling"

	//login
	UserNameIncorrect       string = "Username is Incorrect or Missing"
	PasswordIncorrect       string = "Password is Incorrect or Missing"
	UserAuthorisationFailed string = "User Authorisation Failed"
	UserLoginSuccess        string = "User Logged in Successfully"
	UserIdNotProvided       string = "userId not provided"
	ServerError             string = "Internal Server Error"

	// element load from config
	CouldNotLoadFromEnv string = "Could not load the value from env. file"

	//Get Content
	UnableToGetContent string = "unable to get content"

	InvalidCountryCodeFound string = "InvalidCountryCodeFounds"

	CountryCodeNotProvided string = "country code is not provided "

	//create event constants
	UnableToCreateEvent string = "Unable to Create Event... Please try again"
)

var (
	LogFields = make(map[string]interface{})
)
