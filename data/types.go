package data

import (
	"math/big"
	"time"
)

type GenericProduct struct {
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	PricingModel     string    `json:"pricing_model"`
	SubscriptionType string    `json:"subscription_type"`
	PricingMetric    string    `json:"pricing_metric"`
	UnitPriceUSD     big.Float `json:"unit_price_usd"`
	DisplayPriceUSD  []byte    `json:"display_price_usd"`
	Metadata         []byte    `json:"metadata"`
	DateStarted      time.Time `json:"date_started"`
	DateEnded        time.Time `json:"date_ended"`
}

type BuildAgent struct {
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	PricingModel     string    `json:"pricing_model"`
	SubscriptionType string    `json:"subscription_type"`
	PricingMetric    string    `json:"pricing_metric"`
	UnitPriceUSD     big.Float `json:"unit_price_usd"`
	DisplayPriceUSD  struct {
		PerHour  big.Float `json:"per_hour"`
		PerMonth big.Float `json:"per_month"`
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
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	PricingModel     string    `json:"pricing_model"`
	SubscriptionType string    `json:"subscription_type"`
	PricingMetric    string    `json:"pricing_metric"`
	UnitPriceUSD     big.Float `json:"unit_price_usd"`
	DisplayPriceUSD  struct {
		PerHour  big.Float `json:"per_hour"`
		PerMonth big.Float `json:"per_month"`
	} `json:"display_price_usd"`
	Metadata struct {
		CPU int `json:"cpu"`
	} `json:"metadata"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

type KubeProduct struct {
	KubeAgent []*KubeAgent `json:"kube_agent"`
}

type Database struct {
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	PricingModel     string    `json:"pricing_model"`
	SubscriptionType string    `json:"subscription_type"`
	PricingMetric    string    `json:"pricing_metric"`
	UnitPriceUSD     big.Float `json:"unit_price_usd"`
	DisplayPriceUSD  struct {
		PerHour  big.Float `json:"per_hour"`
		PerMonth big.Float `json:"per_month"`
	} `json:"display_price_usd"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

type DatabaseInfo struct {
	Name     string      `json:"name"`
	Versions []string    `json:"versions"`
	Database []*Database `json:"database"`
}

type DatabaseProduct map[string]*DatabaseInfo

type Subscription struct {
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	PricingModel     string    `json:"pricing_model"`
	SubscriptionType string    `json:"subscription_type"`
	PricingMetric    string    `json:"pricing_metric"`
	UnitPriceUSD     big.Float `json:"unit_price_usd"`
	DisplayPriceUSD  struct {
		PerMonthM2M big.Float `json:"per_month_m2m"`
		PerMonthAPM big.Float `json:"per_month_apm"`
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

type SubscriptionProduct struct {
	Subscription []*Subscription `json:"subscription"`
}
