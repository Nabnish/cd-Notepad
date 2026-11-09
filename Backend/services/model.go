package services

type loadbalancer struct {
	ID int `json:"id" bson:"id"`
	RequestCount int `json:"request_Count" bson:"request_Count"`
	NoOfServNodes int `json:"no_of_servnodes" bson:"no_of_servnodes"`
	IncommingRequest []request `json:"incomming_request" bson:"incomming_request"`

}

type request struct{
	ID int `json:"id" bson:"id"`
	ResourcesNeeded int `json:"resources_needed" bson:"resources_needed"`
	


}