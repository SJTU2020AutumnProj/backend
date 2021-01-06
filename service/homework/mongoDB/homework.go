package mongoDB

import (
	"context"
	// "log"
	"fmt"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Homework struct {
	HomeworkID   int32 `json:"homework_id,omitempty" bson:"homework_id,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
	Note string `json:"note,omitempty" bson:"note,omitempty"`
}

type Answer struct {
	AnswerID int32 `json:"answer_id,omitempty" bson:"answer_id,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
	Note string `json:"note,omitempty" bson:"note,omitempty"`
}

func (Homework) TableName() string {
	return "homework"
}

type HomeworkMongo interface {
	AddHomework(ctx context.Context, homework Homework) error
	DeleteHomework(ctx context.Context, homeworkID int32) error
	UpdateHomework(ctx context.Context, homework Homework) error
	SearchHomework(ctx context.Context, homeworkID int32) (Homework,error)
	// SearchHomeworkByTeacherID(ctx context.Context, homeworkID int32) ([]*Homework,error)
}

type HomeworkMongoImpl struct {
	CL *mongo.Collection
}

func (repo *HomeworkMongoImpl) AddHomework(ctx context.Context, homework Homework) error {
	id := homework.HomeworkID
	description := homework.Description
	content := homework.Content
	note := homework.Note
	h := Homework{id, description, content,note}

	insertResult, err := repo.CL.InsertOne(ctx, h)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

func (repo *HomeworkMongoImpl) DeleteHomework(ctx context.Context,homeworkID int32) error {	
	filter := bson.M{"homework_id": homeworkID}
	deleteResult1, err := repo.CL.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	return nil
}

func (repo *HomeworkMongoImpl) UpdateHomework(ctx context.Context,homework Homework) error {
	filter := bson.M{"homework_id": homework.HomeworkID}
	update := bson.M{"$set": bson.M{ "description": homework.Description, "content":homework.Content }}
	updateResult, err := repo.CL.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func (repo *HomeworkMongoImpl) SearchHomework(ctx context.Context, homeworkID int32) (Homework,error){
	filter := bson.M{"homework_id": homeworkID}
	var result Homework
	err := repo.CL.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result,err
	}
	return result,nil
}
