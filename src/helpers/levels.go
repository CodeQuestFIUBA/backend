package helper

import (
	"codequest/src/database"
	"codequest/src/models"
	"context"
	"time"

	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateScoresForUserAndRoom(userID string, room string) ([]models.Score, error) {
	var scoreCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "scores")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var scores []models.Score

	for _, level := range models.Levels {
		for _, subLevel := range level.SubLevels {
			score := models.Score{
				ID:       primitive.NewObjectID(),
				Room:     room,
				Level:    level.Key,
				SubLevel: subLevel.Key,
				UserID:   userID,
				Attempts: 0,
				Points:   0,
			}
			scores = append(scores, score)
		}
	}

	_, err := scoreCollection.InsertMany(ctx, scoresToInterface(scores))
	if err != nil {
		log.Println("Error inserting scores:", err)
		return nil, err
	}

	return scores, nil
}

func scoresToInterface(scores []models.Score) []interface{} {
	interfaces := make([]interface{}, len(scores))
	for i, score := range scores {
		interfaces[i] = score
	}
	return interfaces
}

func ConvertLevelsToResponseFormat() []models.LevelResponse {
	var responseLevels []models.LevelResponse

	for _, level := range models.Levels {
		var subLevels []models.SubLevelInfo
		for _, subLevel := range level.SubLevels {
			subLevels = append(subLevels, models.SubLevelInfo{
				Key:   subLevel.Key,
				Title: subLevel.Title,
			})
		}

		responseLevels = append(responseLevels, models.LevelResponse{
			Title:     level.Title,
			Key:       level.Key,
			SubLevels: subLevels,
		})
	}

	return responseLevels
}
