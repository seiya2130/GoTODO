package firebase

import (
	"GoTODO/config"
	"GoTODO/model"
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var app *firebase.App
var ctx = context.Background()

func init() {
	opt := option.WithCredentialsFile(config.FirebaseConfigPath)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}
}

func CreateTask(task model.Task) (*model.Task, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ref, _, err := client.Collection("tasks").Add(ctx, task)
	if err != nil {
		return nil, err
	}

	task.ID = ref.ID
	return &task, nil
}
