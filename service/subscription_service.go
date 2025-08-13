package service

import (
	"log/slog"

	"github.com/teryble09/subscription_service/api"
)

type SubscriptionService struct {
	Logger  *slog.Logger
	Storage SubscriptionStorage

	api.UnimplementedHandler
}

type SubscriptionStorage interface {
	SubscriptionLister
	SubscriptionCreator
	SubscriptionDeleter
	SubscriptionGetter
	SubscriptionUpdater
	CostCalculator
}
