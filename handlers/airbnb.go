package handlers

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"airbnb-api/database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AirbnbData holds the information returned by the API
type AirbnbData struct {
	RoomID        string  `json:"room_id" bson:"room_id"`
	OccupancyRate float64 `json:"occupancy_rate_last_5_months" bson:"occupancy_rate_last_5_months"`
	AverageRate   float64 `json:"average_night_rate_next_30_days" bson:"average_night_rate_next_30_days"`
	HighestRate   float64 `json:"highest_night_rate_next_30_days" bson:"highest_night_rate_next_30_days"`
	LowestRate    float64 `json:"lowest_night_rate_next_30_days" bson:"lowest_night_rate_next_30_days"`
}

// GetAirbnbData handles the API call, fetching data from MongoDB and returning mock data if not found
func GetAirbnbData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomID := params["room_id"]

	// Try to fetch data from MongoDB
	data, err := fetchAirbnbDataFromDB(roomID)

	if err != nil && err != mongo.ErrNoDocuments {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// If data not found in MongoDB, generate mock data
	if err == mongo.ErrNoDocuments {
		data = generateMockData(roomID)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// fetchAirbnbDataFromDB fetches data from MongoDB
func fetchAirbnbDataFromDB(roomID string) (AirbnbData, error) {
	var data AirbnbData
	collection := database.GetCollection("airbnb_data")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"room_id": roomID}

	err := collection.FindOne(ctx, filter).Decode(&data)

	return data, err
}

// generateMockData creates random mock data
func generateMockData(roomID string) AirbnbData {
	rand.Seed(time.Now().UnixNano())

	occupancyRate := 60 + rand.Float64()*40   // Random occupancy rate between 60% and 100%
	averageRate := 50 + rand.Float64()*150    // Random average rate between 50 and 200
	highestRate := averageRate + rand.Float64()*50
	lowestRate := averageRate - rand.Float64()*30

	return AirbnbData{
		RoomID:        roomID,
		OccupancyRate: occupancyRate,
		AverageRate:   averageRate,
		HighestRate:   highestRate,
		LowestRate:    lowestRate,
	}
}
