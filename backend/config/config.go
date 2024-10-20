package config

import (
	"os"

	"github.com/google/uuid"
)

var SuperDeciderUUID = uuid.MustParse(os.Getenv("SUPERUSER_UUID"))
