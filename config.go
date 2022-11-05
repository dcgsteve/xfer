package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getConfigDir() string {
	// set path and file names
	return fmt.Sprintf("%s/.config/", os.Getenv("HOME"))
}

func getConfigFile() string {
	return "xfer.conf"
}

func getConfigFullName() string {
	return fmt.Sprintf("%s%s", getConfigDir(), getConfigFile())
}

func saveConfig(c *xferConfig) error {

	// create dir if not there already
	if !fileExists(getConfigDir()) {
		e := os.Mkdir(getConfigDir(), 0755)
		if e != nil {
			return e
		}
	}

	// write out config
	e := os.WriteFile(getConfigFullName(), []byte(c.ServerEndpoint), 0600)
	if e != nil {
		return e
	}

	// done
	return nil
}

func loadConfig() (*xferConfig, error) {

	c := new(xferConfig)

	if fileNotExists(getConfigFullName()) {
		fmt.Println("Enter your transfer.sh server endpoint (e.g. https://transfer.sh): ")

		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Sorry, i couldn't understand what you typed: %s", err)
		}

		// remove trailing delimeters (in order)
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "/")
		input = strings.TrimSuffix(input, "\\")

		// create default config based on input
		c.ServerEndpoint = input
		err = saveConfig(c)
		if err != nil {
			log.Fatalf("Failed to save your configuration file: %s", err)
		}
	} else {
		// read in config
		b, err := os.ReadFile(getConfigFullName())
		if err != nil {
			return nil, err
		}
		c.ServerEndpoint = string(b)
	}

	// done
	return c, nil
}

// func resetConfig() (*xferConfig, error) {
// 	c := new(xferConfig)
// 	c.ServerEndpoint = "http://localhost"
// 	err := saveConfig(c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return c, nil
// }

// func encConfig(c *xferConfig) ([]byte, error) {

// 	var buf bytes.Buffer
// 	enc := gob.NewEncoder(&buf)
// 	e := enc.Encode(c)
// 	if e != nil {
// 		return nil, e
// 	}

// 	r, e := encBytes(buf.Bytes(), EncryptionKey)
// 	if e != nil {
// 		return nil, e
// 	}

// 	return r, nil

// }

// func decConfig(f []byte) (*xferConfig, error) {

// 	var c xferConfig

// 	d, e := decBytes(f, EncryptionKey)
// 	if e != nil {
// 		return nil, e
// 	}

// 	buf := bytes.NewBuffer(d)
// 	dec := gob.NewDecoder(buf)

// 	e = dec.Decode(&c)
// 	if e != nil {
// 		return nil, e
// 	}

// 	return &c, nil
// }
