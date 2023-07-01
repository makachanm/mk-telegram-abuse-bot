package botservices

type AbuseIDStorage struct {
	AbuseIDs map[string]string
}

func NewAbuseIDStorage() AbuseIDStorage {
	return AbuseIDStorage{AbuseIDs: make(map[string]string)}
}

func IsExistInMap(src *map[string]string, target string) bool {
	v := *src
	_, exist := v[target]
	return exist
}

func IsExistInArray(src []MisskeyAbuse, target string) bool {
	for _, v := range src {
		if v.AbuseID == target {
			return true
		}
	}

	return false
}

func (as *AbuseIDStorage) InsertInital(ind []string) {
	for _, indx := range ind {
		as.AbuseIDs[indx] = indx
	}
}

func (as *AbuseIDStorage) FindAbuseFromID(abss []MisskeyAbuse, input []string) []MisskeyAbuse {
	var result []MisskeyAbuse = make([]MisskeyAbuse, 0)

	for _, val := range input {
		for _, aval := range abss {
			if aval.AbuseID == val {
				result = append(result, aval)
			}
		}
	}

	return result
}

func (as *AbuseIDStorage) UpdateDiffrence(abss []MisskeyAbuse) ([]string, []string) {
	var added_ab []string = make([]string, 0)
	var removed_ab []string = make([]string, 0)

	for _, val := range abss {
		if !IsExistInMap(&as.AbuseIDs, val.AbuseID) {
			as.AbuseIDs[val.AbuseID] = val.AbuseID
			added_ab = append(added_ab, val.AbuseID)
		}
	}

	for _, val := range as.AbuseIDs {
		if !IsExistInArray(abss, val) {
			delete(as.AbuseIDs, val)
			removed_ab = append(removed_ab, val)
		}
	}

	return added_ab, removed_ab
}
