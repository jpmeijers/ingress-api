package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	_ "time"
	"ttnmapper-ingress-api/ttnV2"
	"ttnmapper-ingress-api/types"
)

func TtnRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/v2", PostTtnV2)
	router.Get("/v2", GetTtnV2)

	return router
}

/*
TTN V2 HTTP integration
Authorization header contains email address
There is an extra Experiment field added to the model
*/
func PostTtnV2(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	email := r.Header.Get("Authorization")
	log.Print("["+i+"] User: ", email)
	if err := validateEmail(email); err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print("[" + i + "] " + err.Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print("[" + i + "] " + err.Error())
		return
	}

	var packetIn ttnV2.UplinkMessage
	if err := json.Unmarshal(body, &packetIn); err != nil {
		response["success"] = false
		response["message"] = "Can not parse json body"
		log.Print("[" + i + "] " + err.Error())
		return
	}

	var packetOut types.TtnMapperUplinkMessage
	packetOut.NetworkType = types.NS_TTN_V2
	packetOut.NetworkAddress = r.RemoteAddr

	packetOut.UserAgent = r.Header.Get("USER-AGENT")
	packetOut.UserId = email
	packetOut.Experiment = packetIn.Experiment

	// Try to use the location from the metadata first. This is likely the location set on the console.
	if packetIn.Metadata.Latitude != 0 && packetIn.Metadata.Longitude != 0 {
		packetOut.AccuracySource = packetIn.Metadata.Source
		packetOut.Latitude = float64(packetIn.Metadata.Latitude)
		packetOut.Longitude = float64(packetIn.Metadata.Longitude)
		packetOut.Altitude = float64(packetIn.Metadata.Altitude)
	}

	// If payload fields are available, try getting coordinates from there
	if packetIn.PayloadFields != nil {
		if err := ParsePayloadFields(int64(packetIn.FPort), packetIn.PayloadFields, &packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print("[" + i + "] " + err.Error())
			return
		}
	}

	// Ignore data validity checks for experiments
	if packetOut.Experiment == "" {
		if err := CheckData(packetOut); err != nil {
			response["success"] = false
			response["message"] = err.Error()
			log.Print("["+i+"] Data invalid: ", err.Error())
			return
		}
		SanitizeData(&packetOut)
	}

	// Copy metadata fields
	CopyTtnV2Fields(packetIn, &packetOut)

	log.Print("["+i+"] Network: ", packetOut.NetworkType, "://", packetOut.NetworkAddress)
	log.Print("["+i+"] Device: ", packetOut.AppID, " - ", packetOut.DevID)

	publishChannel <- packetOut

	response["success"] = true
	response["message"] = "New packet accepted into queue"
}

func GetTtnV2(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["message"] = "GET test success"
	render.JSON(w, r, response)
}
