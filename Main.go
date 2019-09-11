package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/git-lfs/gitobj"
)

func main() {
	nodes, edges := visitAll()
	output, _ := json.MarshalIndent(Data{Nodes: nodes, Edges: edges}, "", "   ")
	fmt.Print(string(output))
}

func visitAll() (nodes []Node, edges []Edge) {
	nodes, edges = visitObjects(nodes, edges)
	nodes, edges = visitBranches(nodes, edges)
	nodes, edges = visitTags(nodes, edges)

	return
}

type Data struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func addEdgesFromTree(id string, tree *gitobj.Tree) (edges []Edge) {
	for _, entry := range tree.Entries {
		edges = append(edges, Edge{From: id, To: hex.EncodeToString(entry.Oid)})
	}

	return
}

func addEdgesFromCommit(id string, commit *gitobj.Commit) (edges []Edge) {
	edges = append(edges, Edge{From: id, To: hex.EncodeToString(commit.TreeID)})
	for _, parent := range commit.ParentIDs {
		edges = append(edges, Edge{From: id, To: hex.EncodeToString(parent)})
	}

	return
}

func walkObjects() (nodes []string) {
	var dir = ""
	filepath.Walk(".git/objects", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			nodes = append(nodes, fmt.Sprintf("%s%s", dir, info.Name()))
		} else {
			switch info.Name() {
			case "info", "pack":
				return filepath.SkipDir
			}

			dir = info.Name()
		}
		return nil
	})

	return
}

func visitBranches(nodes []Node, edges []Edge) ([]Node, []Edge) {
	return visitRef(".git/refs/heads", nodes, edges)
}

func visitTags(nodes []Node, edges []Edge) ([]Node, []Edge) {
	return visitRef(".git/refs/tags", nodes, edges)
}

func visitRef(dir string, nodes []Node, edges []Edge) ([]Node, []Edge) {
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		nodes = append(nodes, Node{Id: file.Name(), Type: "branch"})

		data := readFirstLine(path.Join(dir, file.Name()))
		edges = append(edges, Edge{From: file.Name(), To: string(data)})
	}

	return nodes, edges
}

func readFirstLine(path string) string {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Text()
}

type Node struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

func visitObjects(nodes []Node, edges []Edge) ([]Node, []Edge) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	objects := walkObjects()
	for _, id := range objects {
		node, e, _ := visitObject(repo, id)
		nodes = append(nodes, node)
		edges = append(edges, e...)
	}

	return nodes, edges
}

func visitObject(repo *gitobj.ObjectDatabase, id string) (node Node, edges []Edge, err error) {
	sha, _ := hex.DecodeString(id)

	object, err := repo.Object(sha)
	if err != nil {
		return Node{}, edges, err
	}

	switch v := object.(type) {
	case *gitobj.Tree:
		edges = append(addEdgesFromTree(id, v))
		return Node{Type: "tree", Id: id}, edges, nil
	case *gitobj.Commit:
		edges = append(addEdgesFromCommit(id, v))
		return Node{Type: "commit", Id: id}, edges, nil
	case *gitobj.Blob:
		return Node{Type: "blob", Id: id}, edges, nil
	}

	return Node{}, edges, fmt.Errorf("Unkown object type for sha-ish %s", id)
}
