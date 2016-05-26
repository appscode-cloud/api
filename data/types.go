package data

import (
	"encoding/json"
	"math/big"
	"time"
)

type Product struct {
	Sku              string          `json:"sku"`
	DisplayName      string          `json:"display_name"`
	PricingModel     string          `json:"pricing_model"`
	SubscriptionType string          `json:"subscription_type"`
	PricingMetric    string          `json:"pricing_metric"`
	UnitPriceUSD     Money           `json:"unit_price_usd"`
	DisplayPriceUSD  json.RawMessage `json:"display_price_usd"`
	Metadata         json.RawMessage `json:"metadata"`
	DateStarted      time.Time       `json:"date_started"`
	DateEnded        time.Time       `json:"date_ended"`
}

type BuildAgent struct {
	Sku              string `json:"sku"`
	DisplayName      string `json:"display_name"`
	PricingModel     string `json:"pricing_model"`
	SubscriptionType string `json:"subscription_type"`
	PricingMetric    string `json:"pricing_metric"`
	UnitPriceUSD     Money  `json:"unit_price_usd"`
	DisplayPriceUSD  struct {
		PerHour  Money `json:"per_hour"`
		PerMonth Money `json:"per_month"`
	} `json:"display_price_usd"`
	Metadata struct {
		Provider             string `json:"provider"`
		Os                   string `json:"os"`
		CPU                  int    `json:"cpu"`
		RAM                  int    `json:"ram"`
		Disk                 int    `json:"disk"`
		RecommendedExecutors int    `json:"recommended_executors"`
		ExternalID           string `json:"external_id"`
	} `json:"metadata"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

type CIProduct struct {
	BuildAgent []*BuildAgent `json:"build_agent"`
}

type KubeAgent struct {
	Sku              string `json:"sku"`
	DisplayName      string `json:"display_name"`
	PricingModel     string `json:"pricing_model"`
	SubscriptionType string `json:"subscription_type"`
	PricingMetric    string `json:"pricing_metric"`
	UnitPriceUSD     Money  `json:"unit_price_usd"`
	DisplayPriceUSD  struct {
		PerHour  Money `json:"per_hour"`
		PerMonth Money `json:"per_month"`
	} `json:"display_price_usd"`
	Metadata struct {
		CPU int `json:"cpu"`
	} `json:"metadata"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

type ClusterProduct struct {
	KubeAgent []*KubeAgent `json:"kube_agent"`
}

type DB struct {
	Sku              string `json:"sku"`
	DisplayName      string `json:"display_name"`
	PricingModel     string `json:"pricing_model"`
	SubscriptionType string `json:"subscription_type"`
	PricingMetric    string `json:"pricing_metric"`
	UnitPriceUSD     Money  `json:"unit_price_usd"`
	DisplayPriceUSD  struct {
		PerHour  Money `json:"per_hour"`
		PerMonth Money `json:"per_month"`
	} `json:"display_price_usd"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

type DBInfo struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
	DB       []*DB    `json:"db"`
}

type DBProduct map[string]*DBInfo

type Package struct {
	Sku              string `json:"sku"`
	DisplayName      string `json:"display_name"`
	PricingModel     string `json:"pricing_model"`
	SubscriptionType string `json:"subscription_type"`
	PricingMetric    string `json:"pricing_metric"`
	UnitPriceUSD     Money  `json:"unit_price_usd"`
	DisplayPriceUSD  struct {
		PerMonthM2M Money `json:"per_month_m2m"`
		PerMonthAPM Money `json:"per_month_apm"`
	} `json:"display_price_usd"`
	Metadata struct {
		User        int `json:"user"`
		Phabricator struct {
			DiskGB int `json:"disk_gb"`
		} `json:"phabricator"`
		Artifact struct {
			DiskGB int `json:"disk_gb"`
		} `json:"artifact"`
		Ci struct {
			BuildAgent int `json:"build_agent"`
		} `json:"ci"`
		Cluster struct {
			KubeAgent int `json:"kube_agent"`
		} `json:"cluster"`
		Database struct {
			Postgres      int `json:"postgres"`
			ElasticSearch int `json:"elasticsearch"`
			Influx        int `json:"influx"`
		} `json:"database"`
	} `json:"metadata"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

type PackageProduct struct {
	Package []*Package `json:"package"`
}

type Money string

const moneyPrecision = 40

func (m Money) Float() (*big.Float, bool) {
	return new(big.Float).SetPrec(moneyPrecision).SetString(string(m))
}

func (m Money) String() string {
	return string(m)
}
