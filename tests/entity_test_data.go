package tests

type Entity struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}
type InvalidEntity struct {
	Id int `json:"id"`
}

// type InvalidEntity struct {
// 	Id int `json:"id"`
// }
