package config

import (
	"os"
	"strconv"
)

func GetEnvInt(envVar string, defaults int) int {
	value := os.Getenv(envVar)
	if value == "" {
		return defaults
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}

func GetEnvString(envVar string, defaults string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaults
	}
	return value
}

var (
	GRPCPort = GetEnvInt("GS_GRPC_PORT", 2303)
)
