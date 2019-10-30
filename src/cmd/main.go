/* © Copyright IBM Corp. 2018
All Rights Reserved.
This software is the confidential and proprietary information
of the IBM Corporation. (‘Confidential Information’). Redistribution
of the source code or binary form is not permitted without prior authorization
from the IBM Corporation.
*/

package main

import (
	"flag"
	"github.com/AnnaKertesz/cp-music-lib/src/transport"
	"go.uber.org/zap"
	"net/http"
)

func main(){
	listenAddr := flag.String("listen", "0.0.0.0:3003", "Listen address")
	flag.Parse()
	logger, _ := zap.NewDevelopment()
	transp := transport.NewHTTP(logger)

	logger.Info("Started", zap.String("version", "dev"))
	if err := http.ListenAndServe(*listenAddr, transp.Routes()); err != nil {
		logger.Fatal("Could not start HTTP server", zap.Error(err))
	}
}