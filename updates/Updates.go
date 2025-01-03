package updates

type Updates struct {
	Update_id                 int64                        `json:"update_id"`
	Message                   *Message                     `json:"message"`
	Edited_message            *Message                     `json:"edited_message"`
	Channel_post              *Message                     `json:"channel_post"`
	Edited_channel_post       *Message                     `json:"edited_channel_post"`
	Business_connection       *BusinessConnection          `json:"business_connection"`
	Edited_business_message   *Message                     `json:"edited_business_message"`
	Deleted_business_messages *BusinessMessagesDeleted     `json:"deleted_business_messages"`
	Message_reaction          *MessageReactionUpdated      `json:"message_reaction"`
	Message_reaction_count    *MessageReactionCountUpdated `json:"message_reaction_count"`
	Inline_query              *InlineQuery                 `json:"inline_query"`
	Chosen_inline_result      *ChosenInlineResult          `json:"chosen_inline_result"`
	Callback_query            *CallbackQuery               `json:"callback_query"`
	Shipping_query            *ShippingQuery               `json:"shipping_query"`
	Pre_checkout_query        *PreCheckoutQuery            `json:"pre_checkout_query"`
	Purchased_paid_media      *PaidMediaPurchased          `json:"purchased_paid_media"`
	Poll                      *Poll                        `json:"poll"`
	Poll_answer               *PollAnswer                  `json:"poll_answer"`
	My_chat_member            *ChatMemberUpdated           `json:"my_chat_member"`
	Chat_member               *ChatMemberUpdated           `json:"chat_member"`
	Chat_join_request         *ChatJoinRequest             `json:"chat_join_request"`
	Chat_boost                *ChatBoostUpdated            `json:"chat_boost"`
	Removed_chat_boost        *ChatBoostRemoved            `json:"removed_chat_boost"`
}
