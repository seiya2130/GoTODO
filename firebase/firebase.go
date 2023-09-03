package firebase

import (
	"GoTODO/config"
	"GoTODO/model"
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func GetTaskByID(id string) (*model.Task, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	doc, err := client.Collection("tasks").Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, fmt.Errorf("Document with ID %s not found", id)
		}
		return nil, err
	}

	task := model.Task{}
	doc.DataTo(&task)
	task.ID = doc.Ref.ID

	return &task, nil
}

func UpdateTask(id string, updatedTask model.Task) error {
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Collection("tasks").Doc(id).Set(ctx, updatedTask)
	return err
}

func DeleteTask(id string) error {
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Collection("tasks").Doc(id).Delete(ctx)
	return err
}
