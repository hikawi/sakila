package utils

import (
	"log"
	"os"
)

// FailOnError fails the program and exits with a bad exit code if the error passed in is non-null.
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg)
	}
}

// FatalEnv retrieves an environment variable and fails the app if it doesn't exist.
func FatalEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("fatal: env variable %s not found\n", key)
	}
	return val
}
