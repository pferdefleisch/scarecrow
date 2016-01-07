// Package types contains shareable types between sub-modules.
package types

////////////////////////////////////////////////////////////////////////////////
// Configuration for web.json
////////////////////////////////////////////////////////////////////////////////

type WebConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}