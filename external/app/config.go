package app

import (
	"flag"
)

type Config struct {
	size          int
	nums          bool
	alphabetLower bool
	alphabetUpper bool
}

func getConfig() *Config {
	size := flag.Int("size", 8, "length of generated password")
	alphabetLower := flag.Bool("a", false, "use alphabet lowercase symbols")
	alphabetUpper := flag.Bool("A", false, "use alphabet uppercase symbols")
	nums := flag.Bool("n", false, "use numbers")

	flag.Parse()

	return &Config{
		size:          *size,
		nums:          *nums,
		alphabetLower: *alphabetLower,
		alphabetUpper: *alphabetUpper,
	}
}
