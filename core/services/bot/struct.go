package botservices

type MkAbuseState string

const AB_LOCAL = "local"
const AB_REMOTE = "remote"

const UNRESOLVED MkAbuseState = "unresolved"
const RESOLVED MkAbuseState = "resolved"

type Misskey struct {
	MisskeyToken string
	InstanceURL  string
}

type MisskeyAbuse struct {
	AbuseID         string `json:"id"`
	AbuseComment    string `json:"comment"`
	IsAbuseSolved   bool   `json:"resolved"`
	AbuseReporterID string `json:"reporterId"`
	AbuseTargetID   string `json:"targetUserId"`
}
