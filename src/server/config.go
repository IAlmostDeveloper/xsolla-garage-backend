package server

type Config struct {
	DbConnection  string `yml:"dbConnection"`
	RedisAddr     string `yml:"redisAddr"`
	RedisPassword string `yml:"redisPassword"`
}
