package updates

type User struct {
    Id int64 `json:"id"`
    Is_bot bool `json:"is_bot"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
    Username string `json:"username"`
    Language_code string `json:"language_code"`
    Is_premium bool `json:"is_premium"`
    Added_to_attachment_menu bool `json:"added_to_attachment_menu"`
    Can_join_groups bool `json:"can_join_groups"`
    Can_read_all_group_messages bool `json:"can_read_all_group_messages"`
    Supports_inline_queries bool `json:"supports_inline_queries"`
    Can_connect_to_business bool `json:"can_connect_to_business"`
    Has_main_web_app bool `json:"has_main_web_app"`
}