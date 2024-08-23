package main

const (
	darkMode mode = "dark"
	lightMode mode = "light"
)

type mode string

type AsciiArtConfig struct {
	mode mode
}

func main() {
	cfg := AsciiArtConfig{
		mode: darkMode,
	}
	repl(&cfg)
}