package database

type Database interface {
	GetName() string
	SetConfig() string
}
