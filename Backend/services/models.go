package services

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)
type LoadBalancer struct {
	ID               int       `json:"id,omitempty" bson:"_id,omitempty"`
	RequestCount     int       `json:"request_Count" bson:"request_Count"`
	IncommingRequest []Request `json:"incomming_request" bson:"incomming_request"`
	NoOfServNodes    int       `json:"no_of_servnodes" bson:"no_of_servnodes"`
	ServNodes        []Server  `json:"server_nodes" bson:"server_nodes"`
}

type Request struct {
	ID             int  `json:"id,omitempty" bson:"_id,omitempty"`
	ProcessingTime int  `json:"time_required" bson:"time_required"`
	CreatedAt      int  `json:"created_at bson:created_at"`
	Exprired       bool `json:"expired" bson:"expired"`
}

type Server struct {
	ID               int       `json:"id,omitempty" bson:"_id,omitempty"`
	RequestCount     int       `json:"request_Count" bson:"request_Count"`
	IncommingRequest []Request `json:"incomming_request" bson:"incomming_request"`
	MaxRequests      int       `json:"max_request" bson:"max_request"`
}

var client *mongo.Client

func New(mongo *mongo.Client) LoadBalancer{
	client = mongo
	return LoadBalancer{ }
}

func returnCollectionPointer(collectionName string) *mongo.Collection{
	collection := client.Database("notepad").Collection(collectionName)
	return collection

}

func(r *Request) InsertRequest(Req Request) error {
	collection := returnCollectionPointer("Request")

	_,err := collection.InsertOne(context.Background(),Request{
		ProcessingTime : Req.ProcessingTime,
		CreatedAt :  Req.CreatedAt, 
		Exprired : Req.Exprired,
	})
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}