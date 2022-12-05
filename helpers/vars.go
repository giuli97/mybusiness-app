package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

var SecretKey = os.Getenv("secret")

const AllowedCORSDomain = "http://localhost"
