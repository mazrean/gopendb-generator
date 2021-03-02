package config

import "testing"

func TestConfig(t *testing.T) {
	t.Parallel()

	t.Run("ReadYAML", readYAMLTest)
}
