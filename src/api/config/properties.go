/*
Package config provides a manager to work with fury's config server.
For more information about the service: https://docs.furycloud.io/?lang=esp&section=configuration
*/

package config

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/magiconair/properties"
)

const (
	_defaultConfigPath      = "/configs/latest/application.properties"
	_propertyConfigFileName = "configFileName"
	_checksumEnabled        = "checksumEnabled"
)

// Config provides all configurations loaded from the fury's configuration.
type Config struct {
	prop     *properties.Properties
	filename string
}

// Load loads the configurations.
func Load() (*Config, error) {
	if c := os.Getenv(_propertyConfigFileName); c != "" {
		return load(c)
	}

	return load(_defaultConfigPath)
}

func load(filename string) (*Config, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading configuration: %v", err)
	}

	if err = verify(b, filename); err != nil {
		return nil, fmt.Errorf("verifying configuration: %v", err)
	}

	prop, err := properties.Load(b, properties.UTF8)
	if err != nil {
		return nil, fmt.Errorf("loading configuration: %v", err)
	}

	return &Config{
		prop:     prop,
		filename: filename,
	}, nil
}

func verify(b []byte, filename string) error {
	if c := os.Getenv(_checksumEnabled); c == "false" {
		return nil
	}

	if len(b) == 0 {
		return fmt.Errorf("the file %s is empty", filename)
	}

	filenameMD5 := filename + ".md5"

	expectedMD5, err := os.ReadFile(filenameMD5)
	if err != nil {
		return err
	}

	if len(expectedMD5) == 0 {
		return fmt.Errorf("the file %s is empty", filenameMD5)
	}

	currentMD5, err := md5FromBytes(b)
	if err != nil {
		return err
	}

	if !bytes.Equal(currentMD5, expectedMD5) {
		return fmt.Errorf("different md5 contents")
	}

	return nil
}

func md5FromBytes(b []byte) ([]byte, error) {
	hash := md5.New()
	if _, err := hash.Write(b); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err := hex.NewEncoder(&buf).Write(hash.Sum(nil)); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
