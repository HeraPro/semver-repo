package main

import (
	"fmt"
	_ "github.com/HeraPro/demo"
	_ "github.com/spf13/cobra"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "im not semver-repo yet")
}
