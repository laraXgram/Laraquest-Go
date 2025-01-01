package updates

type UsersShared struct {
    Request_id int64 `json:"request_id"`
    Users *[]SharedUser `json:"users"`
}