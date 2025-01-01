package updates

type Contact struct {
    Phone_number string `json:"phone_number"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
    User_id int64 `json:"user_id"`
    Vcard string `json:"vcard"`
}