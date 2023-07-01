package botservices

import (
	"fmt"
	"strings"
)

func MessageBuilder(mkb MisskeyAbuse) string {
	mka_id := mkb.AbuseID
	mka_ctx := mkb.AbuseComment

	msgbuilder := strings.Builder{}
	msgbuilder.WriteString("**New Abuse is Reported!** \n\n")
	msgbuilder.WriteString("*Abuse ID:* %v \n")
	msgbuilder.WriteString("*Comment:* \n")
	msgbuilder.WriteString("```\n%v\n```")

	return fmt.Sprintf(msgbuilder.String(), mka_id, mka_ctx)
}
