package main

import (
	"testing"

	"github.com/BestCharlemagne/StoreNavigator/repository"
)

//TestMain tests the main function
func TestMain(t *testing.T) {
	str := repository.Item{}
	print(str.Type)
}
