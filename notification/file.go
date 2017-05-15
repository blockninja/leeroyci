package notification

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/blockninja/leeroyci/database"
)

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
	txtMessage := message(job, database.NotificationServiceEmail, event, TypeText)

	payload := &FileInfo{
		Channel: channel,
		Message: txtMessage,
	}
	var b bytes.Buffer

	err = toml.NewEncoder(&b).Encode(payload)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(b.Bytes()))
	ioutil.WriteFile("/home/users/leeroyci/go/bin/postbox", b.Bytes(), 0644)
}
