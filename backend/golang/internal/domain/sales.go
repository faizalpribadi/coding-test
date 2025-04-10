package domain

// DummyData represents the structure of the dummyData.json file
type DummyData struct {
	SalesReps []SalesRep `json:"salesReps"`
}

// SalesRep represents a sales representative
type SalesRep struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Role          string   `json:"role"`
	Region        string   `json:"region"`
	Skills        []string `json:"skills"`
	Deals         []Deal   `json:"deals"`
	Clients       []Client `json:"clients"`
	FilteredDeals []Deal   `json:"filtered_deals,omitempty"`
}

type Deal struct {
	Client string `json:"client"`
	Value  int    `json:"value"`
	Status string `json:"status"`
}

type Client struct {
	Name     string `json:"name"`
	Industry string `json:"industry"`
	Contact  string `json:"contact"`
}
