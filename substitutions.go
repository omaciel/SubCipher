package main

// Substitution cipher called Zenit Polar
func zenit_polar_substitutions() map[string]string {
	substitutions := make(map[string]string)
	substitutions["p"] = "z"
	substitutions["o"] = "e"
	substitutions["l"] = "n"
	substitutions["a"] = "i"
	substitutions["r"] = "t"
	substitutions["z"] = "p"
	substitutions["e"] = "o"
	substitutions["n"] = "l"
	substitutions["i"] = "a"
	substitutions["t"] = "r"
	return substitutions
}
