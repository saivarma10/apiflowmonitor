package processor

import (
	logger "apimonitor/pkg/logger"

	curl "github.com/andelf/go-curl"
)

var log = logger.GetLogger()

func CurlGet(url string) string {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	easy.Setopt(10002, url)
	if err := easy.Perform(); err != nil {
		return ""
		// log.Info().Msgf("ERROR: %v\n", err)
	}
	log.Info().Msg("Sucess curl get")
	return ""

}
