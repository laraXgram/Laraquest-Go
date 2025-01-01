package updates

type PollOption struct {
    Text string `json:"text"`
    Text_entities *[]MessageEntity `json:"text_entities"`
    Voter_count int64 `json:"voter_count"`
}