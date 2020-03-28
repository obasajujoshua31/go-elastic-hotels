package services

type Hotel struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Stars   string `json:"stars"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	URI     string `json:"uri"`
}
