package main

var projectType string

type ConfigType interface {
	addUser(name string)
}

type Config struct {
}

// ConcretProduct
type ConfigCorporate struct {
	name   string
	region string
	path   string
}

// Concret Product
type ConfigHealth struct {
	name   string
	region string
	path   string
	health bool
}
