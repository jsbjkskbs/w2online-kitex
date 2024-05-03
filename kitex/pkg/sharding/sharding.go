package sharding

import "gorm.io/sharding"

func NewSharding(shardingKey string, shardingNumber uint, tableName string) *sharding.Sharding {
	middleware := sharding.Register(sharding.Config{
		ShardingKey:         shardingKey,
		NumberOfShards:      shardingNumber,
		PrimaryKeyGenerator: sharding.PKMySQLSequence,
	}, tableName)
	return middleware
}
