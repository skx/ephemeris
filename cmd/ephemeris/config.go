// config.go - Load our configuration struct.

package main

import (
	"encoding/json"
	"os"
)

// Config is the configuration object we use to guide our generation.
//
// This structure is populated from `ephemeris.json` when the application
// is launched.
type Config struct {
	// Posts holds the directory beneath which we can find blog-posts.
	Posts string

	// Prefix is the URL-prefix of the generated site, for example
	// https://blog.steve.fi/
	Prefix string

	// CommentAPI holds the (CGI) endpoint to be used for submitting
	// comments to, if that support is enabled.
	CommentAPI string

	// Comments points to a directory containing comment-files.
	Comments string

	// Output is the path to which we write our output files.
	OutputPath string

	// AddComments is used to determine whether there is an 'add comment'
	// form shown on the most recent entry.
	AddComments bool
}

// loadConfig loads the specified JSON file, and returns a
// configuration-object from the contents.
func loadConfig(path string) (Config, error) {

	//
	// The structure we'll populate
	//
	var config Config

	//
	// Read the config-file
	//
	configFile, err := os.Open(path)
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
