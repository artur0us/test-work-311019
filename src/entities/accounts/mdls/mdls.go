package mdls

type BasicAuthAnswer struct {
	Token           string          `json:"token"`
	AccountTinyInfo AccountTinyInfo `json:"accounts_tiny_info"`
}

type AccountTinyInfo struct {
	AccountId      int64  `json:"account_id"`
	Username       string `json:"username"`
	CreatedAt      int64  `json:"created_at"`
	AccountGroupID int    `json:"acc_group_id"`
}
