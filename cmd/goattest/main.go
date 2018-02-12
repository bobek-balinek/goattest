package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		prefix := ""

		if strings.HasPrefix(trimmed, "--- FAIL") {
			prefix = `             / /
          (\/_//')
           /   '/
          0  0   \
         /        \
        /    __/   \
       /,  _/ \     \_
       '-./ )  |     ~^~^~^~^~^~^~^~\~.
           (   /                     \_}
              |               /      |
              ;     |         \      /
               \/ ,/           \    |
               / /~~|~|~~~~~~|~|\   |
              / /   | |      | | '\ \
             / /    | |      | |   \ \
            / (     | |      | |    \ \
           /,_)    /__)     /__)   /,_/
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`
		}

		fmt.Println(prefix + line)
	}
}
