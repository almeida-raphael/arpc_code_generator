package configs

// Max stores max configuration values for colfer size parameters
type Max struct {
	ArraySize      string `config:"array-size;default=10 * 1024 * 1024 * 1024"`
	SerializedSize string `config:"serialized-size;default=10 * 1024 * 1024 * 1024"`
}

// Config stores configuration for the generator application
type Config struct {
	Max              Max
	InputPath        string `config:"input-path;default=./arpc"`
	PackagesRootPath string `config:"packages-root-path;default=./"`
}