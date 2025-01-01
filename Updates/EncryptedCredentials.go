package updates

type EncryptedCredentials struct {
    Data string `json:"data"`
    Hash string `json:"hash"`
    Secret string `json:"secret"`
}