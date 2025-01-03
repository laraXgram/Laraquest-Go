package updates

type Chat struct {
    Id int64 `json:"id"`
    Type string `json:"type"`
    Title string `json:"title"`
    Username string `json:"username"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
    Is_forum bool `json:"is_forum"`
}