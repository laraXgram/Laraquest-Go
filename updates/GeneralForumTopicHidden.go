package updates

type GeneralForumTopicHidden struct {
    User_id int64 `json:"user_id"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
    Username string `json:"username"`
    Photo *[]PhotoSize `json:"photo"`
}