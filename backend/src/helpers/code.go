package helper

import (
	"context"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateNewCode(ctx context.Context, classRoomCollection *mongo.Collection, baseCode *string) (*string, error) {
	code := *baseCode
	counter := 1

	for {
		count, err := classRoomCollection.CountDocuments(ctx, bson.M{"code": code})
		if err != nil {
			return nil, err
		}

		if count == 0 {
			return &code, nil
		}

		code = fmt.Sprintf("%s-%d", *baseCode, counter)
		counter++
	}
}
