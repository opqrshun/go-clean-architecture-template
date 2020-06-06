package tests

type Entity struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
type InvalidEntity struct {
	ID int `json:"id"`
}

// type InvalidEntity struct {
// 	ID int `json:"id"`
// }
