package logger

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	Debug string = "DEBUG"
)

// AccessLogger is the logger handle for logging service access events.
var AccessLogger *logrus.Logger

// timeFormat is the format of timestamp to log on all events.
var timeFormat = "2023-10-23 13:04:05"

// Logger is the logger handle for logging all service events.
var Logger *logrus.Logger

// log file used to save the logs of the running code
var date = time.Now().Format("02-01-2006")
var Accesslogfile = "access-" + date + ".log"

// init initializes all the logger handlers during service startup.This will be executed only once at the start of service.
func init() {

	// initAccessLogger initializes the logger handle for logging service access events.
	initAccessLogger()

	// initLogger initializes the logger handle for logging all service events.
	initLogger()

}

// initAccessLogger initializes the logger handle for logging service access events.
func initAccessLogger() {

	AccessLogger = logrus.New()
	AccessLogger.Level = getLogLevel(Debug)
	fmt.Println(AccessLogger.Level)
	AccessLogger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("./log/"+Accesslogfile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		AccessLogger.SetOutput(os.Stdout)
	} else {
		AccessLogger.SetOutput(file)
	}
}

// initLogger initializes the logger handle for logging all service events.
func initLogger() {
	Logger = logrus.New()
	Logger.Level = getLogLevel(Debug)
	Logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("./log/"+Accesslogfile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		Logger.SetOutput(os.Stdout)
	} else {
		logwriter := io.MultiWriter(os.Stdout, file)
		Logger.SetOutput(logwriter)
	}
}

func getLogLevel(loglevel string) logrus.Level {
	switch loglevel {
	case "DEBUG":
		return logrus.DebugLevel
	default:
		return logrus.ErrorLevel
	}
}

// getSystemIPAddress returns the system's IP address.
func getSystemIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil
		}
	}

	return "", fmt.Errorf("unable to determine the system's IP address")
}

// LogDebug ..
func LogDebug(obj interface{}, fields map[string]interface{}) {
	ipAddress, err := getSystemIPAddress()
	if err == nil {
		fields["ip_address"] = ipAddress
	}
	Logger.WithFields(fields).Debug(obj)
}

// LogError ..
func LogError(obj interface{}, fields map[string]interface{}) {
	ipAddress, err := getSystemIPAddress()
	if err == nil {
		fields["ip_address"] = ipAddress
	}
	Logger.WithFields(fields).Error(obj)
}
