package notification

import (
	"log"

	"github.com/blockninja/leeroyci/database"
	"github.com/davecgh/go-spew/spew"
)

// Notify sends all relevant notifications for a job that are configured for
// the jobs repository.
func Notify(job *database.Job, event string) {
	repo, err := database.GetRepositoryByID(job.RepositoryID)
	if err != nil {
		log.Println(err)
		return
	}

	sendWebsocket(job, event)
	spew.Dump(repo)
	for _, notification := range repo.Notifications {
		switch notification.Service {
		case database.NotificationServiceFile:
			sendFile(job, event)
		case database.NotificationServiceEmail:
			sendEmail(job, event)
		case database.NotificationServiceSlack:
			sendSlack(job, event)
		case database.NotificationServiceHipchat:
			sendHipchat(job, event)
		case database.NotificationServiceCampfire:
			sendCampfire(job, event)
		default:
			continue
		}
	}
}
