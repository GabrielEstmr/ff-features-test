package ff_mongo

import (
	main_configs_ff_lib "ff-features-go/ff"
	ff_mongo_exceptions "ff-features-go/ff/mongo/errors"
	ff_mongo_repo "ff-features-go/ff/mongo/repo"
	ff_resources "ff-features-go/ff/resources"
)

type FeaturesMongoMethodsImpl struct {
	repo *ff_mongo_repo.FeaturesMongoRepo
}

func NewFeaturesMongoMethodsImpl(ffConfigData *main_configs_ff_lib.FfConfigData) *FeaturesMongoMethodsImpl {
	return &FeaturesMongoMethodsImpl{repo: ff_mongo_repo.NewFeaturesMongoRepo(ffConfigData)}
}

func (this *FeaturesMongoMethodsImpl) getFeature(key string) (
	ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}

	if byId.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	return byId.ToDomain(), nil
}

func (this *FeaturesMongoMethodsImpl) IsEnabled(key string,
) (bool, ff_mongo_exceptions.LibException) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsEnabled(), nil
}

func (this *FeaturesMongoMethodsImpl) IsDisabled(key string,
) (bool, ff_mongo_exceptions.LibException) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsDisabled(), nil
}

func (this *FeaturesMongoMethodsImpl) Enable(key string,
) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {
	featureDoc, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}
	if featureDoc.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	if featureDoc.IsDisabled() {
		featureDoc.DefaultValue = true
		savedFeatureDoc, err := this.repo.Update(*featureDoc)
		if err != nil {
			return *new(ff_resources.FeaturesData),
				ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
		}
		return savedFeatureDoc.ToDomain(), nil
	}
	return featureDoc.ToDomain(), nil
}

func (this *FeaturesMongoMethodsImpl) Disable(key string,
) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {
	featureDoc, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}
	if featureDoc.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	if featureDoc.IsEnabled() {
		featureDoc.DefaultValue = false
		savedFeatureDoc, err := this.repo.Update(*featureDoc)
		if err != nil {
			return *new(ff_resources.FeaturesData),
				ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
		}
		return savedFeatureDoc.ToDomain(), nil
	}
	return featureDoc.ToDomain(), nil
}
