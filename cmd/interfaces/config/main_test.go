package config

import "testing"

func TestMain(t *testing.T) {
	t.Parallel()

	t.Run("ReadYAML", readYAMLTest)
	t.Run("Config", configTest)
	t.Run("Table", tableTest)
}
