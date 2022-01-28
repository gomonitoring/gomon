package tasks

import (
	"log"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gomonitoring/http-server/internal/database"
	"github.com/gomonitoring/http-server/internal/model"
)

func FindUrlsToCall() error {
	db, _ := database.NewDB()
	var urls []model.Url
	query := db.Find(&urls)
	groupSigs := make([]*tasks.Signature, query.RowsAffected)
	for i, url := range urls {
		sig := tasks.Signature{
			Name:       "call_url",
			RoutingKey: "monitoring",
			Args: []tasks.Arg{
				{
					Name:  "url",
					Type:  "string",
					Value: url.Url,
				}, {
					Name:  "id",
					Type:  "uint",
					Value: url.ID,
				}, {
					Name:  "threshhold",
					Type:  "int",
					Value: url.Threshold,
				}, {
					Name:  "resetTime",
					Type:  "int64",
					Value: url.ResetTime,
				},
			},
		}
		groupSigs[i] = &sig
	}
	collectorSig := tasks.Signature{
		Name:       "collect_results",
		RoutingKey: "local",
	}
	group, _ := tasks.NewGroup(groupSigs...)
	chord, _ := tasks.NewChord(group, &collectorSig)
	_, err := GetMachineryServer().SendChord(chord, 0)
	if err != nil {
		log.Fatal("Could not push tasks to queue.")
	}
	return nil
}
