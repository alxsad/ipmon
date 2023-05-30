package main

import (
	"bytes"
	"context"
	"log"
	"time"
)

func main() {
	var (
		state    []byte
		newState []byte
		ip       string
		country  string
		err      error
	)

	for {
		time.Sleep(time.Second * 3)

		var ctx, _ = context.WithTimeout(context.Background(), time.Second*3)

		if err, newState = getRoutes(ctx); err != nil {
			log.Printf("[ERROR] get routes: %s", err)
			continue
		}

		if bytes.Equal(state, newState) {
			continue
		}

		log.Println("[INFO] routes has been modified")

		if err, ip = getIpAddress(ctx); err != nil {
			log.Printf("[ERROR] resolve address: %s", err)
			continue
		}

		state = newState

		log.Printf("[INFO] ip address: %s", ip)

		if err, country = getCountry(ctx, ip); err != nil {
			log.Printf("[ERROR] resolve country: %s", err)
			continue
		}

		log.Printf("[INFO] country: %s", country)

		if err = sendNotification(ctx, country, ip); err != nil {
			log.Printf("[ERROR] show notification: %s", err)
			continue
		}
	}
}
