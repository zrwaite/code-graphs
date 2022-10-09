package config

type Config struct {
	Directory            string
	SendGridAPIKey       string
	ContactEmail         string
	RedisPassword        string
	FromEmail            string
	MongoUrl             string
	WakatimeClientId     string
	WakatimeClientSecret string
	Dev                  bool
	RedirectURI          string
	AdminPassword        string
}
