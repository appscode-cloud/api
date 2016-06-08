package data

import (
	"encoding/json"
	"math/big"
	"time"
)

type Product struct {
	Sku              string          `json:"sku"`
	Type             string          `json:"type"`
	SubType          string          `json:"sub_type"`
	QuotaRequired    string          `json:"quota_required"`
	DisplayName      string          `json:"display_name"`
	PricingModel     string          `json:"pricing_model"`
	SubscriptionType string          `json:"subscription_type"`
	TimeUnit         string          `json:"time_unit"`
	PricingFormula   string          `json:"pricing_formula"`
	DisplayPriceUSD  json.RawMessage `json:"display_price_usd"`
	Metadata         json.RawMessage `json:"metadata"`
	Quota            json.RawMessage `json:"quota"`
	DateStarted      time.Time       `json:"date_started"`
	DateEnded        time.Time       `json:"date_ended"`
}

type BuildAgent struct {
	Sku              string `json:"sku"`
	Type             string `json:"type"`
	DisplayName      string `json:"display_name"`
	PricingModel     string `json:"pricing_model"`
	SubscriptionType string `json:"subscription_type"`
	TimeUnit         string `json:"time_unit"`
	PricingFormula   string `json:"pricing_formula"`
	DisplayPriceUSD  struct {
		PerHour  Money `json:"per_hour"`
		PerMonth Money `json:"per_month"`
	} `json:"display_price_usd"`
	Metadata struct {
		Provider             string `json:"provider"`
		OS                   string `json:"os"`
		CPU                  int    `json:"cpu"`
		RAM                  int    `json:"ram"`
		Disk                 int    `json:"disk"`
		RecommendedExecutors int    `json:"recommended_executors"`
		ExternalSKU           string `json:"external_sku"`
	} `json:"metadata"`
	Quota       map[string]int `json:"quota"`
	DateStarted time.Time      `json:"date_started"`
	DateEnded   time.Time      `json:"date_ended"`
}

type CIProduct struct {
	BuildAgent []*BuildAgent `json:"build_agent"`
}

type KubeAgent struct {
	Sku              string `json:"sku"`
	Type             string `json:"type"`
	DisplayName      string `json:"display_name"`
	PricingModel     string `json:"pricing_model"`
	SubscriptionType string `json:"subscription_type"`
	TimeUnit         string `json:"time_unit"`
	PricingFormula   string `json:"pricing_formula"`
	DisplayPriceUSD  struct {
		PerHour  Money `json:"per_hour"`
		PerMonth Money `json:"per_month"`
	} `json:"display_price_usd"`
	Quota       map[string]int `json:"quota"`
	DateStarted time.Time      `json:"date_started"`
	DateEnded   time.Time      `json:"date_ended"`
}

type ClusterProduct struct {
	KubeAgent []*KubeAgent `json:"kube_agent"`
}

type DB struct {
	Sku              string `json:"sku"`
	Type             string `json:"type"`
	DisplayName      string `json:"display_name"`
	PricingModel     string `json:"pricing_model"`
	SubscriptionType string `json:"subscription_type"`
	TimeUnit         string `json:"time_unit"`
	PricingFormula   string `json:"pricing_formula"`
	DisplayPriceUSD  struct {
		PerHour  Money `json:"per_hour"`
		PerMonth Money `json:"per_month"`
	} `json:"display_price_usd"`
	Quota       map[string]int `json:"quota"`
	DateStarted time.Time      `json:"date_started"`
	DateEnded   time.Time      `json:"date_ended"`
}

type DBInfo struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
	DB       []*DB    `json:"db"`
}

type DBProduct map[string]*DBInfo

type Package struct {
	Sku              string `json:"sku"`
	Type             string `json:"type"`
	DisplayName      string `json:"display_name"`
	PricingModel     string `json:"pricing_model"`
	SubscriptionType string `json:"subscription_type"`
	TimeUnit         string `json:"time_unit"`
	PricingFormula   string `json:"pricing_formula"`
	DisplayPriceUSD  struct {
		PerMonthM2M Money `json:"per_month_m2m"`
		PerMonthAPM Money `json:"per_month_apm"`
	} `json:"display_price_usd"`
	Quota struct {
		PkgUser           int `json:"pkg.user"`
		PhabricatorDiskGB int `json:"phabricator.disk_gb"`
		ArtifactDiskGB    int `json:"artifact.disk_gb"`
		CIBuildAgent      int `json:"ci.build_agent"`
		ClusterKubeAgent  int `json:"cluster.kube_agent"`
		DBPostgres        int `json:"db.postgres"`
		DBElasticSearch   int `json:"db.elasticsearch"`
		DBInflux          int `json:"db.influx"`
	} `json:"metadata"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

type PackageProduct struct {
	Package []*Package `json:"package"`
}

type InstanceType struct {
	Category string      `json:"category"`
	Sku      string      `json:"sku"`
	CPU      int         `json:"cpu"`
	Memory   interface{} `json:"memory"`
}

type CloudInstance struct {
	Name    string `json:"name"`
	Regions []struct {
		Location string   `json:"location"`
		Region   string   `json:"region"`
		Zones    []string `json:"zones,omitempty"`
		Zone     []string `json:"zone,omitempty"`
	} `json:"regions"`
	InstanceTypes []*InstanceType `json:"instance_types"`
}
type ClusterProvider struct {
	Provider map[string]CloudInstance `json:"cloud_provider"`
}

type Money string

const moneyPrecision = 40

func (m Money) Float() (*big.Float, bool) {
	return new(big.Float).SetPrec(moneyPrecision).SetString(string(m))
}

func (m Money) String() string {
	return string(m)
}
