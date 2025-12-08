package utils

import (
	"log"
	"os"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient snap.Client
var CoreClient coreapi.Client

func InitMidtrans() {
	// 1. Ambil Server Key dari env
	serverKey := strings.TrimSpace(os.Getenv("MIDTRANS_SERVER_KEY"))
	if serverKey == "" {
		log.Fatal("MIDTRANS_SERVER_KEY is not set or empty. Please set the environment variable or add it to .env")
	}

	env := strings.TrimSpace(os.Getenv("MIDTRANS_ENVIRONMENT"))

	var midtransEnv midtrans.EnvironmentType
	if env == "production" {
		midtransEnv = midtrans.Production
	} else {
		midtransEnv = midtrans.Sandbox
	}

	// 2. Inisialisasi Snap Client
	SnapClient.New(serverKey, midtransEnv)
	// Inisialisasi Core API Client (untuk CheckTransaction)
	CoreClient.New(serverKey, midtransEnv)

	// 3. Logging
	midtrans.DefaultLoggerLevel = &midtrans.LoggerImplementation{LogLevel: midtrans.LogInfo}
}