package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"os"
	"time"

	"github.com/urfave/cli/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gopkg.in/gookit/color.v1"
)

/* 0722
	task manager command line program and learned the fundamentals of using the MongoDB Go driver 
*/

type Task struct{
	ID primitive.ObjectID 		 `bson:"_id"'`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

var collection *mongo.Collection
var ctx=context.TODO()

func init(){

	clientOptions:=options.Client().ApplyURI("mongodb://localhost:27017/")
	client,err:=mongo.Connect(ctx,clientOptions)
	if err!=nil {
		log.Fatal(err)
	}

	err=client.Ping(ctx,nil)
	if err!=nil {
		log.Fatal(err)
	}

	collection=client.Database("local").Collection("test1")
	insertMany()
}

func main(){

	app:=&cli.App{
		Name: "tasker",
		Usage:"A simple CLI program to manage your tasks",
		Action: func(c *cli.Context) error {
			tasks, err := getPending()
			if err != nil {
				if err == mongo.ErrNoDocuments {
					fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
					return nil
				}

				return err
			}

			printTasks(tasks)
			return nil
		},
		Commands:[]*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					str := c.Args().First()
					if str == "" {
						return errors.New("Cannot add an empty task")
					}

					isCompleted:= false
					if(c.Args().Len()>1){
						fmt.Print(c.Args().Get(1))
						if(c.Args().Get(1)=="true"){
							isCompleted=true
						}
					}

					task := &Task{
						ID:        primitive.NewObjectID(),
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Text:      str,
						Completed: isCompleted,
					}

					return createTask(task)
				},
			},
			//command all
			{
				Name:    "all",
				Aliases: []string{"l"},
				Usage:   "lst all tasks",
				Action: func(c *cli.Context) error {
					tasks,err:=getAll()
					if err!=nil{
						if err==mongo.ErrNoDocuments{
							fmt.Print("nothings to see here,no docs\nrun add to add task")
							return nil
						}
						return err
					}
					printTasks(tasks)
					return nil
				},
			},
			//command done
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					text:=c.Args().First()
					return completeTask(text)
				},
			},
			//command completed
			{
				Name:    "completed",
				Aliases: []string{"c"},
				Usage:   "lst all tasks completed",
				Action: func(c *cli.Context) error {
					tasks,err:=getCompleted()
					if err!=nil{
						if err==mongo.ErrNoDocuments{
							fmt.Print("nothings to see here,no docs\nrun add to add task")
							return nil
						}
						return err
					}
					printTasks(tasks)
					return nil
				},
			},
			//command del
			{
				Name:    "del",
				Aliases: []string{"del"},
				Usage:   "delete one task",
				Action: func(c *cli.Context) error {
					text:=c.Args().First()
					err:=deteleTask(text)
					if err!=nil{
						fmt.Println(err.Error())
						return err
					}
					return nil
				},
			},


		},
	}

	err:=app.Run(os.Args)
	if err!=nil {
		log.Fatal(err)
	}
}

//update
func completeTask(text string) error{
	//bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
	filter:=bson.D{primitive.E{Key:"text",Value: text}}//filter
	//filter1:=bson.D{bson.E{Key:"text",Value: text}}//filter

	//filter2:=bson.D{{"text",text}}

	update:=bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{//set
			primitive.E{Key: "completed",Value: true},
		}},
	}

	t:=&Task{}
	//update
	return collection.FindOneAndUpdate(ctx,filter,update).Decode(t)
}

func createTask(task *Task) error{
	_,err:=collection.InsertOne(ctx,task)
	return err
}

func deteleTask(text string) error{

	filter:=bson.D{primitive.E{Key: "text",Value: text}}

	res,err:=collection.DeleteOne(ctx,filter)
	if err!=nil{
		return err
	}

	if res.DeletedCount==0{
		return errors.New("No tasks were deleted")
	}

	return nil
}

func filterTasks(filter interface{})([]*Task,error){

	var tasks []*Task

	cur,err:=collection.Find(ctx,filter)
	if err!=nil {
		return tasks,err
	}

	for cur.Next(ctx) {
		var t Task
		err:=cur.Decode(&t)
		if err!=nil {
			return tasks,err
		}

		tasks=append(tasks, &t)
	}

	if err:=cur.Err();err!=nil{
		return tasks,err
	}

	cur.Close(ctx)
	if len(tasks)==0{
		return tasks,mongo.ErrNoDocuments
	}

	return tasks,nil
}

func getAll()([]*Task,error){
	filter:=bson.D{}
	return filterTasks(filter)
}
func getPending()([]*Task,error){
	filter:=bson.D{
		primitive.E{Key: "completed",Value: false},
	}
	return filterTasks(filter)
}
func getCompleted()([]*Task,error){
	filter:=bson.D{
		primitive.E{Key: "completed",Value: true},
	}
	return filterTasks(filter)
}

func printTasks(tasks []*Task){
	for i,v:=range tasks {
		if v.Completed{
			color.Green.Printf("%d:%s\n",i+1,v.Text)
		}else{
			color.Yellow.Printf("%d:%s\n",i+1,v.Text)
		}
		
	}
}


func insertMany(){
	docs := []interface{}{
		bson.D{{"type", "Masala"}, {"rating", 10},{"text", "Masala"}},
		bson.D{{"type", "Earl Grey"}, {"rating", 5},{"text", "Earl Grey"}},
	}
	
	collection.InsertMany(ctx,docs)

}

func insertOne(){
	
	doc := bson.D{{"type", "Masala"}, {"rating", 10},{"text", "Masala"}}
	collection.InsertOne(ctx,doc)

}




