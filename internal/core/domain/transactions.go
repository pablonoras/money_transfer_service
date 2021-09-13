package domain

type Transaction struct{
	UserID string `json:"user_id"`
	ReceptorID string `json:"receptor_id"`
	TransactionID string `json:"transaction_id, omitempty"`
	Amount int `json:"amount"`
	SiteFrom  string `json:"site_from"`
	SiteTo string `json:"site_to"`
	Status string `json:"status"`
	CreationDate string `json:"creation_date"`
	LastModifiedDate string `json:"last_modified_date"`
}