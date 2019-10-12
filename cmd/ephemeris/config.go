// config.go - Load our configuration struct.

package main

import (
	"encoding/json"
	"os"
)

// Config is the configuration object we use to guide our
// generation.
//
// It is loaded from `ephemeris.json` when the application
// loads. There are no sane defaults, so that file must be
// present.
//
type Config struct {
	// Posts holds the directory beneath which we can find blog-posts
	Posts string

	// Prefix is the URL-prefix of the generated site, for example
	// https://blog.steve.fi/
	Prefix string

	// Comments points to a directory containing comment-files.
	Comments string

	// Output is the path to which we write our output files.
	OutputPath string
}

// loadConfig loads the `ephemeris.json` file, and returns
// a configuration-object.
func loadConfig() (Config, error) {

	//
	// We look for the file "ephemeris.json" in the current
	// working directory.
	//
	var config Config
	configFile, err := os.Open("ephemeris.json")
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	//
	// Create a deserializer using the file.
	//
	jsonParser := json.NewDecoder(configFile)

	//
	// Deserialize into our structure.
	//
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}

	//
	// Return the populated structure.
	//
	return config, nil
}
