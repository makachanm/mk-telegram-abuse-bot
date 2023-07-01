package botservices

import (
	dbmodels "abusebot/core/db/models"
	"context"
	"fmt"

	"github.com/reugn/go-quartz/quartz"
)

type BotService struct {
	desc        string
	GetMkAbuse  MkGetAbuse
	AbuseStor   AbuseIDStorage
	MkData      Misskey
	DBModel     *dbmodels.AbuseModel
	MessageChan chan<- string
}

func NewBotService(MkD Misskey, Dbm *dbmodels.AbuseModel, MesgChan chan string) BotService {
	AbStor := NewAbuseIDStorage()
	AbStor.InsertInital(Dbm.GetAbuse())

	return BotService{
		AbuseStor:   AbStor,
		MkData:      MkD,
		DBModel:     Dbm,
		MessageChan: MesgChan,
	}
}

func (ms *BotService) Description() string {
	return ms.desc
}

func (ms *BotService) Key() int {
	return quartz.HashCode(ms.Description())
}

func (ms *BotService) Execute(_ context.Context) {
	ms.GetMkAbuse = NewMkGetAbuse(ms.MkData)

	d, err := ms.GetMkAbuse.GetAbuse()
	if err != nil {
		fmt.Println("ERR", err)
	}

	var AbuseIDs []string
	for _, eAbu := range d {
		AbuseIDs = append(AbuseIDs, eAbu.AbuseID)
	}

	def_up, def_rm := ms.AbuseStor.UpdateDiffrence(d)

	if len(def_up) >= 1 {
		ms.DBModel.InsertAbuse(def_up)
		added_abs := ms.AbuseStor.FindAbuseFromID(d, def_up)
		fmt.Println("UPDATED: ", added_abs)

		for _, ctx := range added_abs {
			ms.MessageChan <- MessageBuilder(ctx)
		}
	} else if len(def_rm) >= 1 {
		ms.DBModel.DeleteAbuse(def_rm)
	}
}
