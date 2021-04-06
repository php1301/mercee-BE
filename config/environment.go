package config

import (
	"github.com/spf13/viper"
	"log"
)

// use viper package to read .env file
// return the value of the key

func Env(key string) string {
	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory

	viper.SetConfigFile(".env")

	// Find and read the config file

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	
	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error

	value, ok := viper.Get(key).(string)
	// Type assertion return a second boolean value to check whether this is string
	// Or not
	if !ok {
		log.Printf("Environment error: Invalid type assertion for key \"%s\"", key)
		return ""
	}
	return value

}
