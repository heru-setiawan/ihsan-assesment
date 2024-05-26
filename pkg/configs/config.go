package configs

type Config struct {
	Log       log
	Database  database
	Telemetry telemetry
	Service   service
}

func (c *Config) LoadFromEnv() {
	c.Service.loadFromEnv(service{
		Port: 8000,
		Name: "assesment",
	})

	c.Log.loadFromEnv(log{
		Level:           4,
		FormatTimestamp: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		ForceColors:     true,
	})

	c.Database.loadFromEnv(database{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "",
		Database: "assesment",
		Schema:   "test",
	})

	c.Telemetry.loadFromEnv(telemetry{
		Host: "localhost",
		Port: 4318,
	})
}
