package ff_mongo

import (
	main_configs_ff_lib "ff-features-go/ff"
	ff_mongo_documents "ff-features-go/ff/mongo/documents"
	ff_mongo_repo "ff-features-go/ff/mongo/repo"
	ff_resources "ff-features-go/ff/resources"
)

type RegisterMethodsMongoImpl struct {
	repo *ff_mongo_repo.FeaturesMongoRepo
}

func NewRegisterMethodsMongoImpl(ffConfigData *main_configs_ff_lib.FfConfigData) *RegisterMethodsMongoImpl {
	return &RegisterMethodsMongoImpl{repo: ff_mongo_repo.NewFeaturesMongoRepo(ffConfigData)}
}

func (this *RegisterMethodsMongoImpl) getFeature(key string) (ff_resources.FeaturesData, error) {

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData), err
	}

	if byId.IsEmpty() {
		return *new(ff_resources.FeaturesData), nil
	}

	return byId.ToDomain(), nil
}

func (this *RegisterMethodsMongoImpl) RegisterFeatures(features ff_resources.Features) error {
	for k, v := range features {
		feature, err := this.getFeature(k)
		if err != nil {
			return err
		}
		if feature.IsEmpty() {
			_, err2 := this.repo.Save(ff_mongo_documents.NewFeaturesDataDocument(v))
			if err2 != nil {
				return err2
			}
		}
	}
	return nil
}
