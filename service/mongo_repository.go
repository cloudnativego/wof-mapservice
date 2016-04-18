package service

import "github.com/cloudnativego/cfmgo"

// MongoMapRepository anchor struct for mongo repository
type MongoMapRepository struct {
	Collection cfmgo.Collection
}

// NewMongoRepository creates a new mongo map repository
func NewMongoRepository(col cfmgo.Collection) (repo *MongoMapRepository) {
	repo = &MongoMapRepository{
		Collection: col,
	}
	return
}

// GetMap retrieves a map from the mongo database
func (repo *MongoMapRepository) GetMap(mapID string) (gameMap WofMap, err error) {
	return
}

// GetMapList retrieves all maps
func (repo *MongoMapRepository) GetMapList() (maps []WofMap, err error) {
	return
}

// UpdateMap updates an individual map
func (repo *MongoMapRepository) UpdateMap(mapID string, gameMap WofMap) (err error) {
	return
}
