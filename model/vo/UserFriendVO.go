package vo

type UserFriendVO struct {
	FriendId      uint   `json:"friendId,omitempty"`
	FriendName    string `json:"friendName,omitempty"`
	FriendHeadUrl string `json:"friendHeadUrl,omitempty"`
}
