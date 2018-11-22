package models

type Config struct {
	Model

	CommandSql string `json:"commandSql"`
	Type       string `json:"type"`
}
