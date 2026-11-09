package services

type LoadBalancer struct {
	ID int `json:"id" bson:"id"`
	RequestCount int `json:"request_Count" bson:"request_Count"`
	IncommingRequest []Request `json:"incomming_request" bson:"incomming_request"`
	NoOfServNodes int `json:"no_of_servnodes" bson:"no_of_servnodes"`
	ServNodes []Server `json:"server_nodes" bson:"server_nodes"`

}

type Request struct{
	ID int `json:"id" bson:"id"`
	ProcessingTime int `json:"time_required" bson:"time_required"`
	CreatedAt int `json:"created_at bson:created_at"`
	Exprired bool `json:"expired" bson:"expired"`

}

type Server struct{
	ID int `json:"id" bason:"id"`
	RequestCount int `json:"request_Count" bson:"request_Count"`
	IncommingRequest []Request `json:"incomming_request" bson:"incomming_request"`
	MaxRequests int `json:"max_request" bson:"max_request"`

}




