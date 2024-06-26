package main

import "time"

type ResultProjects struct {
	Projects   []Project `json:"projects"`
	TotalCount int       `json:"total_count"`
}

type ResultUsers struct {
	Users map[string]Owner `json:"users"`
}
type ResponseProjects struct {
	Status string         `json:"status"`
	Result ResultProjects `json:"result"`
}
type ResponseUsers struct {
	Status string      `json:"status"`
	Result ResultUsers `json:"result"`
}
type Currency struct {
	ID           int     `json:"id"`
	Code         string  `json:"code"`
	Sign         string  `json:"sign"`
	Name         string  `json:"name"`
	ExchangeRate float64 `json:"exchange_rate"`
	Country      string  `json:"country"`
	IsExternal   bool    `json:"is_external"`
}

type Budget struct {
	Minimum float64 `json:"minimum"`
	Maximum float64 `json:"maximum"`
}
type Location struct {
	Country struct {
		Name string `json:"name"`
	} `json:"country"`
}

type Upgrades struct {
	Featured          bool `json:"featured"`
	Sealed            bool `json:"sealed"`
	Nonpublic         bool `json:"nonpublic"`
	Fulltime          bool `json:"fulltime"`
	Urgent            bool `json:"urgent"`
	Qualified         bool `json:"qualified"`
	NDA               bool `json:"NDA"`
	IpContract        bool `json:"ip_contract"`
	NonCompete        bool `json:"non_compete"`
	ProjectManagement bool `json:"project_management"`
	PfOnly            bool `json:"pf_only"`
	Premium           bool `json:"premium"`
	Enterprise        bool `json:"enterprise"`
}

type Owner struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Profile   string `json:"profile"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatar_url"`
}

type Project struct {
	ID             int      `json:"id"`
	OwnerID        int      `json:"owner_id"`
	Title          string   `json:"title"`
	Status         string   `json:"status"`
	SeoURL         string   `json:"seo_url"`
	Currency       Currency `json:"currency"`
	Description    string   `json:"description"`
	SubmitDate     int64    `json:"submitdate"`
	PreviewDesc    string   `json:"preview_description"`
	Deleted        bool     `json:"deleted"`
	Nonpublic      bool     `json:"nonpublic"`
	HideBids       bool     `json:"hidebids"`
	Type           string   `json:"type"`
	BidPeriod      int      `json:"bidperiod"`
	Budget         Budget   `json:"budget"`
	Featured       bool     `json:"featured"`
	Urgent         bool     `json:"urgent"`
	TimeSubmitted  int64    `json:"time_submitted"`
	TimeUpdated    int64    `json:"time_updated"`
	Upgrades       Upgrades `json:"upgrades"`
	Language       string   `json:"language"`
	FrontendStatus string   `json:"frontend_project_status"`
	Location       Location `json:"location"`
	Local          bool     `json:"local"`

	UpdatedAt time.Time `json:"updated_at"`
}

// type Currency struct {
// 	ID           int     `json:"id"`
// 	Code         string  `json:"code"`
// 	Sign         string  `json:"sign"`
// 	Name         string  `json:"name"`
// 	ExchangeRate float64 `json:"exchange_rate"`
// 	Country      string  `json:"country"`
// }

// type Category struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// type Job struct {
// 	ID                 int      `json:"id"`
// 	Name               string   `json:"name"`
// 	Category           Category `json:"category"`
// 	ActiveProjectCount *int     `json:"active_project_count"`
// 	SeoURL             string   `json:"seo_url"`
// 	SeoInfo            *string  `json:"seo_info"`
// 	Local              bool     `json:"local"`
// }

// type Budget struct {
// 	Minimum     float64 `json:"minimum"`
// 	Maximum     float64 `json:"maximum"`
// 	Name        *string `json:"name"`
// 	ProjectType *string `json:"project_type"`
// 	CurrencyID  *int    `json:"currency_id"`
// }

// type Porject struct {
// 	ID                 int      `json:"id"`
// 	OwnerID            int      `json:"owner_id"`
// 	Title              string   `json:"title"`
// 	Status             string   `json:"status"`
// 	SeoURL             string   `json:"seo_url"`
// 	Currency           Currency `json:"currency"`
// 	Description        *string  `json:"description"`
// 	Jobs               []Job
// 	SubmitDate         int64    `json:"submitdate"`
// 	PreviewDescription string   `json:"preview_description"`
// 	Deleted            bool     `json:"deleted"`
// 	Nonpublic          bool     `json:"nonpublic"`
// 	HideBids           bool     `json:"hidebids"`
// 	Type               string   `json:"type"`
// 	BidPeriod          int      `json:"bidperiod"`
// 	Budget             Budget   `json:"budget"`
// 	BidStats           BidStats `json:"bid_stats"`
// }

// type BidStats struct {
// 	BidCount int     `json:"bid_count"`
// 	BidAvg   float64 `json:"bid_avg"`
// }
