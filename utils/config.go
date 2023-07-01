package utils

type Tokens struct {
	TelegramToken string `json:"telegram_token"`
	MisskeyToken  string `json:"misskey_token"`
}

type Configure struct {
	Token      Tokens `json:"tokens"`
	InsURL     string `json:"instance_url"`
	DBName     string `json:"db_filename"`
	TargetChan int64  `json:"telegram_channelid"`
}
