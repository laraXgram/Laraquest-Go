package updates

type EncryptedPassportElement struct {
    Type string `json:"type"`
    Data string `json:"data"`
    Phone_number string `json:"phone_number"`
    Email string `json:"email"`
    Files *[]PassportFile `json:"files"`
    Front_side *PassportFile `json:"front_side"`
    Reverse_side *PassportFile `json:"reverse_side"`
    Selfie *PassportFile `json:"selfie"`
    Translation *[]PassportFile `json:"translation"`
    Hash string `json:"hash"`
}