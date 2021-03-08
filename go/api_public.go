/*
 * Fast Counter
 *
 * A fast incrementing counter API
 *
 * API version: 1.0.0
 * Contact: daren@darenliang.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"github.com/darenliang/fast-counter/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func HitKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	key := mux.Vars(r)["key"]

	counters := db.Database.Collection("counters")
	response := counters.FindOneAndUpdate(
		db.Context,
		bson.M{"key": key},
		bson.M{"$inc": bson.M{"value": 1}},
		&options.FindOneAndUpdateOptions{
			Upsert:         &[]bool{true}[0],
			ReturnDocument: &[]options.ReturnDocument{options.After}[0],
		},
	)

	var parsedResponse HitResponse
	if err := response.Decode(&parsedResponse); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(parsedResponse)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
