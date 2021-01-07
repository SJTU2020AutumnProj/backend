package mongoDB

import (
	"context"
	"log"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Check struct
type Check struct {
	CheckID   int32 `json:"check_id,omitempty" bson:"check_id,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Comment string `json:"comment,omitempty" bson:"comment,omitempty"`
}

// TableName configure table name
func (Check) TableName() string {
	return "check"
}

// CheckMongo interface
type CheckMongo interface {
	AddCheck(ctx context.Context, check Check) error
	DeleteCheck(ctx context.Context, checkID int32) error
	UpdateCheck(ctx context.Context, check Check) error
	SearchCheckByID(ctx context.Context, checkID int32) (Check,error)
}

// CheckMongoImpl implementation
type CheckMongoImpl struct {
	CL *mongo.Collection
}

// AddCheck add a check in MongoDB
func (repo *CheckMongoImpl) AddCheck(ctx context.Context, check Check) error {
	id := check.CheckID
	description := check.Description
	comment := check.Comment
	h := Check{id, description, comment}

	_, err := repo.CL.InsertOne(ctx, h)
	if err != nil {
		log.Println("CheckMongo AddCheck error ", err)
		return err
	}
	return nil
}

// DeleteCheck delete a check in MongoDB by its ID
func (repo *CheckMongoImpl) DeleteCheck(ctx context.Context,checkID int32) error {	
	filter := bson.M{"check_id": checkID}
	_, err := repo.CL.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("CheckMongo DeleteCheck error", err)
		return err
	}
	return nil
}

// UpdateCheck update a check in MongoDB
func (repo *CheckMongoImpl) UpdateCheck(ctx context.Context,check Check) error {
	filter := bson.M{"check_id": check.CheckID}
	update := bson.M{"$set": bson.M{ "description": check.Description, "comment": check.Comment}}
	_, err := repo.CL.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("CheckMongo UpdateCheck error", err)
		return err
	}
	return nil
}

// SearchCheckByID search a check by its ID in MongoDB
func (repo *CheckMongoImpl) SearchCheckByID(ctx context.Context, checkID int32) (Check,error){
	filter := bson.M{"check_id": checkID}
	var result Check
	err := repo.CL.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("CheckMongo SearchCheckByID error", err)
		return Check{},err
	}
	return result,nil
}
