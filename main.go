package main

const (
	darkMode mode = "dark"
	lightMode mode = "light"
)

type mode string

type AsciiArtConfig struct {
	scale float64
	
	mode mode
}

func main() {
	// change config as you like
	cfg := AsciiArtConfig{
		scale: 0.85,
		mode: darkMode,
	}
	repl(&cfg)
}