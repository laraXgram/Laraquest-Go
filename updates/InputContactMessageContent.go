package updates

type InputContactMessageContent struct {
    Phone_number string `json:"phone_number"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
    Vcard string `json:"vcard"`
}