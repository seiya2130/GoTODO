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

func GetAllTasks() ([]model.Task, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// タスクのコレクションからドキュメントを取得
	iter := client.Collection("tasks").Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	for _, doc := range docs {
		task := model.Task{}
		doc.DataTo(&task)
		task.ID = doc.Ref.ID
		tasks = append(tasks, task)
	}
	return tasks, nil
}
