/*
 * sample
 *
 * No description provided (generated by Openapi Generator https://github.com/httptools/http-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://http-generator.tech)
 */

package main

import (
	"log"

	sw "gobackend/internal/api/http"
)

func main() {

	router := sw.NewRouter()
	log.Print("Server started")
	log.Fatal(router.Run(":8080"))
}
