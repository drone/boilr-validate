// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package main

import (
	"net/http"

	"{{ GoModule }}/plugin"
	"github.com/drone/drone-go/plugin/validator"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// spec provides the plugin settings.
type spec struct {
	Bind   string `envconfig:"DRONE_BIND"`
	Debug  bool   `envconfig:"DRONE_DEBUG"`
	Secret string `envconfig:"DRONE_SECRET"`

	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	MinSize int `envconfig:"DRONE_MIN_SIZE"`
	MaxSize int `envconfig:"DRONE_MAX_SIZE"`
}

func main() {
	spec := new(spec)
	err := envconfig.Process("", spec)
	if err != nil {
		logrus.Fatal(err)
	}

	if spec.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if spec.Secret == "" {
		logrus.Fatalln("missing secret key")
	}
	if spec.Bind == "" {
		spec.Bind = ":3000"
	}

	handler := validator.Handler(
		spec.Secret,
		plugin.New(
			// TODO replace or remove these configuration
			// parameters. They are for demo purposes only.
			spec.MinSize,
			spec.MaxSize,
		),
		logrus.StandardLogger(),
	)

	logrus.Infof("server listening on address %s", spec.Bind)

	http.Handle("/", handler)
	logrus.Fatal(http.ListenAndServe(spec.Bind, nil))
}
