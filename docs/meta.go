// Package docs meta.
//
// Documentation of your project API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//	- multipart/form-data
//
//	Produces:
//	- application/json
//
// swagger:meta
package docs

//go:generate swagger generate spec -i ../static/swagger.json -o ../static/swagger.json --scan-models
