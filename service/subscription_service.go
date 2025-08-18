package service

import (
	"github.com/teryble09/subscription_service/api"
)

type SubscriptionService struct {
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
