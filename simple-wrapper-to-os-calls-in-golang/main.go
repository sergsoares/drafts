package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main()  {
	fmt.Println("Wrapper to OS")

	cmd := exec.Command("")
	var out strings.Builder
	var stderr strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Println(err)
	}

	fmt.Printf("STDERR: %q\n", stderr.String())
	fmt.Printf("STDOUT: %q\n", out.String())
	
}

type Shell interface {
	Exec() stderr, error
}




// func (p *provider) ListNodes(cluster string) ([]nodes.Node, error) {
// 	cmd := exec.Command("docker",
// 		"ps",
// 		"-a", // show stopped nodes
// 		// filter for nodes with the cluster label
// 		"--filter", fmt.Sprintf("label=%s=%s", clusterLabelKey, cluster),
// 		// format to include the cluster name
// 		"--format", `{{.Names}}`,
// 	)
// 	lines, err := exec.OutputLines(cmd)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to list nodes")
// 	}
// 	// convert names to node handles
// 	ret := make([]nodes.Node, 0, len(lines))
// 	for _, name := range lines {
// 		ret = append(ret, p.node(name))
// 	}
// 	return ret, nil
// }