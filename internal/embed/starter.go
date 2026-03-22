package embed

import "embed"

// Starter is the embedded filesystem containing the starter template for initializing a new site.
// go:embed ../../starter/* ../../starter/**/*
var Starter embed.FS
