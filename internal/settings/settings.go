package settings

import "os"

var MachineryBroker = os.Getenv("MACHINERY_BROKER")
var MachineryResultBackend = os.Getenv("MACHINERY_RESULTS_BACKEND")
