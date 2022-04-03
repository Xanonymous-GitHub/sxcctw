package env

import "github.com/spf13/viper"

var (
	IsDebugMode             bool
	DBGrpcServerPort        string
	DBGrpcServerMetricsPort string
	DBGrpcServerHost        string
	DBUserName              string
	DBPassword              string
	DBHost                  string
	DBPort                  string
	DBName                  string
	ApiRestServerPort       string
	ApiRootPath             string
)

func init() {
	viper.AutomaticEnv()
	IsDebugMode = viper.GetBool("DEBUG")
	DBGrpcServerPort = viper.GetString("DB_GRPC_SEVER_PORT")
	DBGrpcServerMetricsPort = viper.GetString("DB_GRPC_SERVER_METRICS_PORT")
	DBGrpcServerHost = viper.GetString("DB_GRPC_SERVER_HOST")
	DBUserName = viper.GetString("DB_USERNAME")
	DBPassword = viper.GetString("DB_PASSWORD")
	DBHost = viper.GetString("DB_HOST")
	DBPort = viper.GetString("DB_PORT")
	DBName = viper.GetString("DB_NAME")
	ApiRestServerPort = viper.GetString("API_REST_SERVER_PORT")
	ApiRootPath = viper.GetString("API_ROOT_PATH")
}
