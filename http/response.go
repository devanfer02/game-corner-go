package http 

type ResponseError struct {
	Status		int 	`json:"status"`
	Message		string 	`json:"message"`
}

type Response struct {
	Status		int 		`json:"status"`
	Message		string 		`json:"message"`
	Data		interface{}	`json:"data"`
}