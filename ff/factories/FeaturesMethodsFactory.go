package ff_factories

import (
	"errors"
	main_configs_ff_lib "ff-features-go/ff"
	ff_mongo "ff-features-go/ff/mongo"
)

type FeaturesMethodsFactory struct {
	ffConfigData *main_configs_ff_lib.FfConfigData
}

func NewFeaturesMethodsFactory(ffConfigData *main_configs_ff_lib.FfConfigData) *FeaturesMethodsFactory {
	return &FeaturesMethodsFactory{ffConfigData}
}

func (this *FeaturesMethodsFactory) Get() (main_configs_ff_lib.FeaturesMethods, error) {
	if this.ffConfigData.GetClientType() == main_configs_ff_lib.MONGO && this.ffConfigData.GetHasCaching() == false {
		return ff_mongo.NewFeaturesMongoMethodsImpl(this.ffConfigData), nil
	}

	if this.ffConfigData.GetClientType() == main_configs_ff_lib.MONGO && this.ffConfigData.GetHasCaching() == true && this.ffConfigData.GetCacheClientType() == main_configs_ff_lib.REDIS {
		return ff_mongo.NewFeaturesCachedMongoMethodsImpl(this.ffConfigData), nil
	}

	return nil, errors.New("could not instantiate a valid FeaturesData")
}
