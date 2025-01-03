package updates

type ForceReply struct {
    Force_reply bool `json:"force_reply"`
    Input_field_placeholder string `json:"input_field_placeholder"`
    Selective bool `json:"selective"`
}