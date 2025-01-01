package updates

type PassportData struct {
    Data *[]EncryptedPassportElement `json:"data"`
    Credentials *EncryptedCredentials `json:"credentials"`
}