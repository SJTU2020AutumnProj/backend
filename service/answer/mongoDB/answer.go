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

type Answer struct {
	AnswerID int32  `json:"answer_id,omitempty" bson:"answer_id,omitempty"`
	Content  string `json:"content,omitempty" bson:"content,omitempty"`
	Note     string `json:"note,omitempty" bson:"note,omitempty"`
}

func (Answer) TableName() string {
	return "answer"
}

type AnswerMongo interface {
	AddAnswer(ctx context.Context, answer Answer) error
	DeleteAnswer(ctx context.Context, answerID int32) error
	UpdateAnswer(ctx context.Context, answer Answer) error
	SearchAnswer(ctx context.Context, answerID int32) (Answer, error)
	// SearchAnswerByTeacherID(ctx context.Context, answerID int32) ([]*Answer,error)
}

type AnswerMongoImpl struct {
	CL *mongo.Collection
}

func (repo *AnswerMongoImpl) AddAnswer(ctx context.Context, answer Answer) error {
	id := answer.AnswerID
	content := answer.Content
	note := answer.Note
	h := Answer{id, content, note}

	insertResult, err := repo.CL.InsertOne(ctx, h)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

func (repo *AnswerMongoImpl) DeleteAnswer(ctx context.Context, answerID int32) error {
	filter := bson.M{"answer_id": answerID}
	deleteResult1, err := repo.CL.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	return nil
}

func (repo *AnswerMongoImpl) UpdateAnswer(ctx context.Context, answer Answer) error {
	filter := bson.M{"answer_id": answer.AnswerID}
	deleteResult1, err := repo.CL.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)

	id := answer.AnswerID
	content := answer.Content
	note := answer.Note
	h := Answer{id, content, note}

	insertResult, err := repo.CL.InsertOne(ctx, h)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil

	// filter := bson.M{"answer_id": answer.AnswerID}
	// update := bson.M{"$set": bson.M{ "content": answer.Content, "note":answer.Note}}
	// updateResult, err := repo.CL.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	// return nil
}

func (repo *AnswerMongoImpl) SearchAnswer(ctx context.Context, answerID int32) (Answer, error) {
	filter := bson.M{"answer_id": answerID}
	var result Answer
	err := repo.CL.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
