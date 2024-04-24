package ff_mongo_redis

import (
	"errors"
	ff "ff-features-go/ff"
	ff_redis_repo "ff-features-go/ff/redis/repo"
	ff_resources "ff-features-go/ff/resources"
)

type FeaturesRedisMethodsImpl struct {
	repo *ff_redis_repo.FeaturesRedisRepo
}

func NewFeaturesRedisMethodsImpl(ffConfigData *ff.FfConfigData) *FeaturesRedisMethodsImpl {
	return &FeaturesRedisMethodsImpl{repo: ff_redis_repo.NewFeaturesRedisRepo(ffConfigData)}
}

func (this *FeaturesRedisMethodsImpl) getFeature(key string) (ff_resources.FeaturesData, error) {

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData), err
	}

	if byId.IsEmpty() {
		return *new(ff_resources.FeaturesData), errors.New("feature not found")
	}

	return byId.ToDomain(), nil
}

func (this *FeaturesRedisMethodsImpl) IsEnabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsEnabled(), nil
}

func (this *FeaturesRedisMethodsImpl) IsDisabled(key string) (bool, error) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsDisabled(), nil
}
