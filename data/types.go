package data

import "time"

type SubscriptionProduct struct {
	Sku     string `json:"sku"`
	Price   int    `json:"price"`
	Details struct {
		User        int `json:"user"`
		Phabricator struct {
			DiskSize int `json:"disk_size"`
		} `json:"phabricator"`
		Artifact struct {
			DiskSize int `json:"disk_size"`
		} `json:"artifact"`
		Kube struct {
			Cluster int `json:"cluster"`
			Nodes   int `json:"nodes"`
		} `json:"kube"`
		Ci struct {
			Agent int `json:"agent"`
		} `json:"ci"`
		Db struct {
			Postgres      int `json:"postgres"`
			Elasticsearch int `json:"elasticsearch"`
			Influxdb      int `json:"influxdb"`
		} `json:"db"`
	} `json:"details"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

type SubscriptionProductList struct {
	Primary []SubscriptionProduct `json:"primary"`
	Trial   []SubscriptionProduct `json:"trial"`
}

type CIProduct struct {
	BuildAgents []struct {
		Sku      string `json:"sku"`
		Provider string `json:"provider"`
		Os       string `json:"os"`
		Price    int    `json:"price"`
		Details  struct {
		CPU int `json:"cpu"`
		RAM int `json:"ram"`
		Disk int `json:"disk"`
		RecommendedExecutors int `json:"recommended_executors"`

		} `json:"details"`
		DateStarted time.Time `json:"date_started"`
		DateEnded   time.Time `json:"date_ended"`
	} `json:"build_agents"`
}

type DBProduct struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
	DbTypes  []struct {
		Sku         string    `json:"sku"`
		Price       int       `json:"price"`
		Name        string    `json:"name"`
		DateStarted time.Time `json:"date_started"`
		DateEnded   time.Time `json:"date_ended"`
	} `json:"db_types"`
}

type KubeProduct struct {
	KubeAgents struct {
		Sku         string    `json:"sku"`
		Price       float64   `json:"price"`
		CPU         int       `json:"cpu"`
		DateStarted time.Time `json:"date_started"`
		DateEnded   time.Time `json:"date_ended"`
	} `json:"kube_agents"`
}
