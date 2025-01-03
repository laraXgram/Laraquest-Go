package updates

type StarTransaction struct {
    Id string `json:"id"`
    Amount int64 `json:"amount"`
    Nanostar_amount int64 `json:"nanostar_amount"`
    Date int64 `json:"date"`
    Source *TransactionPartner `json:"source"`
    Receiver *TransactionPartner `json:"receiver"`
}