package config

import (
	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

type Config struct {
	Address               string
	AdService             string
	CartService           string
	CheckoutService       string
	CurrencyService       string
	ProductCatalogService string
	RecommendationService string
	ShippingService       string
}

var cfg *Config = &Config{
	Address:               ":80",
	AdService:             "adservice",
	CartService:           "cartservice",
	CheckoutService:       "checkoutservice",
	CurrencyService:       "currencyservice",
	ProductCatalogService: "productcatalogservice",
	RecommendationService: "recommendationservice",
	ShippingService:       "shippingservice",
}

func Get() Config {
	return *cfg
}

func Address() string {
	return cfg.Address
}

func Load() error {
	configor, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return errors.Wrap(err, "configor.New")
	}
	if err := configor.Load(); err != nil {
		return errors.Wrap(err, "configor.Load")
	}
	if err := configor.Scan(cfg); err != nil {
		return errors.Wrap(err, "configor.Scan")
	}
	return nil
}