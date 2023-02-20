package main

import (
	"github.com/go-kipi/currency/api/server/monolith"
	"github.com/go-kipi/currency/application/defaultApp"
)

func deps() []interface{} {
	return append(internalDeps(), externalDeps()...)
}

func internalDeps() []interface{} {
	return []interface{}{
		defaultApp.NewCurrencyApp,
		monolith.NewCurrencyInterface,
	}
}

func externalDeps() []interface{} {
	return []interface{}{}
}
