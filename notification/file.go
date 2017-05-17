package notification

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/blockninja/leeroyci/database"
)

// FileInfo contains the struct to generate the TOML from
type FileInfo struct {
	Channel string
	Message string
}

func sendFile(job *database.Job, event string) {
	notification, _ := database.GetNotificationForRepoAndType(
		&job.Repository,
		database.NotificationServiceFile,
	)
	channel, err := notification.GetConfigValue("channel")
	if err != nil {
		log.Println(err)
		return
	}

	payload := &FileInfo{
		Channel: channel,
		Message: buildMessage(job),
	}
	var b bytes.Buffer

	err = toml.NewEncoder(&b).Encode(payload)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(b.Bytes()))

	err = ioutil.WriteFile("/home/leeroyci/go/bin/postbox/leeroy", b.Bytes(), 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func buildMessage(job *database.Job) string {
	if job.Passed() {
		return passMessage(job.Branch, job.Name)
	}
	return failMessage(job.Branch, job.Name)
}

func failMessage(branch, name string) string {
	return fmt.Sprintf(":poop: Branch %s build pooped out by %s", branch, name)
}

func passMessage(branch, name string) string {
	return fmt.Sprintf(":heart: Branch %s build made with love by %s", branch, name)
}
