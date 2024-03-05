package services

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/cacheprovider"
	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/logger"
	"github.com/gin-gonic/gin"
)

// GetChacheData function is used for getting the required data save in cache
func GetChacheData(c *gin.Context, cacheKey string) interface{} {
	// Attempt to get data from the cache
	if data, found := cacheprovider.GetMemoryStore().Get(cacheKey); found {
		c.JSONP(http.StatusOK, data)
		logger.LogDebug("Data Featch From the Cache....", Constants.LogFields)
		return data
	}
	return nil
}

// SetChacheData function is used for setting the required data in cache
func SetChacheData(cacheKey string, data interface{}, cacheTime int) {
	cacheprovider.GetMemoryStore().Set(cacheKey, data, time.Duration(cacheTime)*time.Hour)
}

// GetCache function is used to get the cache variables from the env file
func GetCache() (database.Cache, error) {
	// load elements from the env file
	cache, err := database.LoadCache(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables 🚀", err)
	}
	return cache, err
}

// ValidateCountryCode function used for validating the country code
func ValidateCountryCode(countryCode string) error {

	// when country code length is 2
	if len(countryCode) == 2 {
		// country code is valid

		if !Constants.CountryCode(strings.ToLower(countryCode)).IsValidConutryCode() {
			return Constants.ErrInvalidCountryCodeFound
		}
		return nil
	} else if len(countryCode) > 2 {
		// country code is invalid
		return Constants.ErrInvalidCountryCodeFound
	}
	return Constants.ErrInvalidCountryCodeFound
}


