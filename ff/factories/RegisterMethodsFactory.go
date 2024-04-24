package ff_factories

import (
	"errors"
	main_configs_ff_lib "ff-features-go/ff"
	main_configs_ff_lib_mongo "ff-features-go/ff/mongo"
)

type RegisterMethodsFactory struct {
	ffConfigData *main_configs_ff_lib.FfConfigData
}

func NewRegisterMethodsFactory(ffConfigData *main_configs_ff_lib.FfConfigData) *RegisterMethodsFactory {
	return &RegisterMethodsFactory{ffConfigData}
}

func (this *RegisterMethodsFactory) Get() (main_configs_ff_lib.RegisterMethods, error) {
	if this.ffConfigData.GetClientType() == main_configs_ff_lib.MONGO {
		return main_configs_ff_lib_mongo.NewRegisterMethodsMongoImpl(this.ffConfigData), nil
	}
	return nil, errors.New("could not instantiate a valid FeaturesData")
}
