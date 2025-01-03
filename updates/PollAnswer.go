package updates

type PollAnswer struct {
    Poll_id string `json:"poll_id"`
    Voter_chat *Chat `json:"voter_chat"`
    User *User `json:"user"`
    Option_ids []int64 `json:"option_ids"`
}