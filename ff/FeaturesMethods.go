package ff

import (
	ff_mongo_exceptions "ff-features-go/ff/mongo/errors"
	ff_resources "ff-features-go/ff/resources"
)

type FeaturesMethods interface {
	IsEnabled(key string) (bool, ff_mongo_exceptions.LibException)
	IsDisabled(key string) (bool, ff_mongo_exceptions.LibException)
	Enable(key string) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException)
	Disable(key string) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException)
}
