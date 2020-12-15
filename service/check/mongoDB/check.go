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

type Check struct {
	CheckID   int32 `json:"check_id,omitempty" bson:"check_id,omitempty"`
	CheckJson string `json:"check_json,omitempty" bson:"check_json,omitempty"`
}

func (Check) TableName() string {
	return "check"
}

type CheckMongo interface {
	AddCheck(ctx context.Context, check Check) error
	DeleteCheck(ctx context.Context, checkID int32) error
	UpdateCheck(ctx context.Context, check Check) error
	SearchCheck(ctx context.Context, checkID int32) (Check,error)
	// SearchCheckByTeacherID(ctx context.Context, checkID int32) ([]*Check,error)
}

type CheckMongoImpl struct {
	CL *mongo.Collection
}

func (repo *CheckMongoImpl) AddCheck(ctx context.Context, check Check) error {
	id := check.CheckID
	json := check.CheckJson
	h := Check{id, json}

	insertResult, err := repo.CL.InsertOne(ctx, h)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

func (repo *CheckMongoImpl) DeleteCheck(ctx context.Context,checkID int32) error {	
	filter := bson.M{"check_id": checkID}
	deleteResult1, err := repo.CL.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	return nil
}

func (repo *CheckMongoImpl) UpdateCheck(ctx context.Context,check Check) error {
	filter := bson.M{"check_id": check.CheckID}
	update := bson.M{"$set": bson.M{ "check_json": check.CheckJson, }}
	updateResult, err := repo.CL.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func (repo *CheckMongoImpl) SearchCheck(ctx context.Context, checkID int32) (Check,error){
	filter := bson.M{"check_id": checkID}
	var result Check
	err := repo.CL.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result,err
	}
	return result,nil
}
