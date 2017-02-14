package common

import (
	"log"
	"os"
	"strconv"
)

//Reads ENV variable and converts it to integer.
//Uses default value if error occurs.
func readEnvIntParam(val *int, def int, env string) {
	*val = def //default value
	if v := os.Getenv(env); v != "" {
		if iv, err := strconv.Atoi(v); err != nil {
			log.Printf(
				"readEnvIntParam: invalid %s value: %s: using default", env, err)
		} else {
			*val = iv
		}
	}
}

//Call log.Fatalln if err not equal to nil.
func FatalIf(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//Call panic if err not equal to nil.
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
