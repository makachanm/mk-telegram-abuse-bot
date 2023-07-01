package botservices

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type MkGetAbuse struct {
	MkData Misskey
}

type MisskeyAbusePayload struct {
	Token string       `json:"i"`
	Limit int          `json:"limit"`
	State MkAbuseState `json:"state"`
}

func NewMkGetAbuse(iMkdata Misskey) MkGetAbuse {
	return MkGetAbuse{MkData: iMkdata}
}

func (m *MkGetAbuse) GetAbuse() ([]MisskeyAbuse, error) {
	request_dest_url := m.MkData.InstanceURL + "/api/admin/abuse-user-reports"
	get_payload := MisskeyAbusePayload{
		Token: m.MkData.MisskeyToken,
		Limit: 100,
		State: UNRESOLVED,
	}

	pdata, perr := json.Marshal(get_payload)
	if perr != nil {
		println(perr)
		return []MisskeyAbuse{}, perr
	}

	resp, reperr := http.Post(request_dest_url, "application/json", bytes.NewBuffer(pdata))
	if reperr != nil {
		println(reperr)
		return []MisskeyAbuse{}, reperr
	}

	var ab_data []MisskeyAbuse
	var reqb_byte []byte
	reqb_byte, ierr := io.ReadAll(resp.Body)

	if ierr != nil {
		println(ierr)
		return []MisskeyAbuse{}, ierr
	}

	jerr := json.Unmarshal(reqb_byte, &ab_data)
	if jerr != nil {
		println(jerr)
		return []MisskeyAbuse{}, jerr
	}

	return ab_data, nil
}
