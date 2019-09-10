package main

import (
	"testing"

	"github.com/git-lfs/gitobj"
	"github.com/stretchr/testify/assert"
)

func TestVisitObjectTree(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	node, edges, _ := visitObject(repo, "4e84516b47b89c12f2f9bf41f34725ef6ddce099")
	assert.Equal(t, "tree", node.Type)
	assert.Contains(t, edges, Edge{From: "4e84516b47b89c12f2f9bf41f34725ef6ddce099", To: "eea118847928ac06875446004228e11658bcb789"})
}

func TestVisitObjectCommit(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	node, edges, _ := visitObject(repo, "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5")
	assert.Equal(t, "commit", node.Type)
	assert.Contains(t, edges, Edge{From: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5", To: "c932e6ae3a32d3cdea9dd8043c165e52495ea4c9"})
	assert.Contains(t, edges, Edge{From: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5", To: "3dda9e9c4e40b7e1b743793075850eadf5817ab5"})
}

func TestVisitObjectBlob(t *testing.T) {
	repo, _ := gitobj.FromFilesystem(".git/objects", "")
	defer repo.Close()

	node, _, _ := visitObject(repo, "eea118847928ac06875446004228e11658bcb789")
	assert.Equal(t, "blob", node.Type)
}

func TestVisitBranches(t *testing.T) {
	branches, edges := visitBranches()
	assert.Contains(t, branches, Node{Id: "for-testing", Type: "branch"})
	assert.Contains(t, edges, Edge{From: "for-testing", To: "627c86822eaa47167417c2c7fc99ef42c599711a"})
}

func TestFilewalk(t *testing.T) {
	nodes := walkObjects()
	assert.Contains(t, nodes, "c568e498a51aa3a792956fc3e23d9b239631fbcd")
}

func Test(t *testing.T) {
	objects := [...]string{
		"0ec386ccd84a67cad2c62c49168ac1a28b348b55",
		"19e100b0c289bd2e66af31823c4873fc7476c770",
		"19e4c735747a5e959cf068400777c5171aaf2560",
		"26faf771e5f7c50d00a03b47e71b8df52ff8a7a7",
		"39c0b96dcf1b4270fb00064102c038be45732a5e",
		"3a424ede2743e029e5af344c955798ecd3483d2c",
		"3a888485e6f96629a183667cad18b0436d5d81e5",
		"3bb4c02e0e1863e765a2342a6d0fb123b8698151",
		"3dda9e9c4e40b7e1b743793075850eadf5817ab5",
		"417089cdaba1de59a9dcc2ebc95347732abc080f",
		"419ad108fda753eb1479515375998f8812657b86",
		"44b9d119eb1a71daff8c74c6d9548cc86c774003",
		"45c0b0e5a153b4a33811e0053332e7d2fce7156c",
		"46a9cc92e00a7ea3baf5a255591726eeca79d9b8",
		"496ee2ca6a2f08396a4076fe43dedf3dc0da8b6d",
		"4a3f1849c90a7fc5aac73c8fbbd41667a850eb46",
		"4e84516b47b89c12f2f9bf41f34725ef6ddce099",
		"57989ff2500eebdf604b01390d5b46977c2d1f1f",
		"57f83081c3b4f452de54e0971900dd55c41c00dd",
		"61a8dd61febeceff7d7db86c3e9c10beb2215a52",
		"6b9a52cfa1e72ffa9498fcc492ac24dd7df7bc3f",
		"6bd7ab3fdc448d06e1a9ffa98e0fce384fe422b7",
		"6d63fd572892576ce63393aeb2d0f957f162687d",
		"78da287a8d68eeedbd6bf19adce69e912ed9c18f",
		"78fc588c77086de4b5c3829d0c115928e0f09792",
		"7e3a67538ed6eeb676990897c5bd92fc6753aa73",
		"82219fb669939987715611e3e93db4dccf5dbda2",
		"8709fadf421ff79bea117e1195277253074d9bc8",
		"926e35d9421fa7caeb5da2bee7c75de807681d04",
		"968f281fa6c9ccd1b27a01bac75ca3ef99f8a85f",
		"9754373abed581d1f4d714a6094f025d8e6cab6f",
		"9a7300cb9b3f232e34b6b9be8552664f2aa1ca67",
		"a5bb3dc800006ce6c64e18863c31bb4bb5bf3e0d",
		"a94eca39c68a487e99ca3f3ab95bdfbd1f4b210f",
		"ae710a6ef6cd3145a5366d1ed2b2d918d529e88a",
		"af071a10c28390b5580a5050b3f26eb94772efaa",
		"b00a810b580df50bb3bbe958dbd9bb9894b85956",
		"b4d37418f02355a8915c00ca712000d69c89e888",
		"bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5",
		"bbcd4b292f036f0773150cd262afd6c9c513b040",
		"bd1640639746d5655366b11c0988bf64317f0069",
		"c3cd5dac6f60ff151fdb103efba53c1e9a7a4e03",
		"c568e498a51aa3a792956fc3e23d9b239631fbcd",
		"c932e6ae3a32d3cdea9dd8043c165e52495ea4c9",
		"cd71d6f8589f221ffd5d0ff540143d02da157a5d",
		"d042d5f86b4bb0c4eff1b60a89d6e3de729db3a6",
		"d84961d15296c4370eab81ec1d6b0000a53aafe9",
		"d9f111e7a85f0ce71fad355bd434e8f39ee87bf6",
		"e8dab7c09c5d7bd9dd48906b9406b5b6602f798c",
		"ea1737c8e1ea6779db6e317ae9ce4811e64e9ab1",
		"ec46eaa46143e97e638d1a963c9712183a02fd18",
		"eea118847928ac06875446004228e11658bcb789",
		"f4fc5076816e7e7d76341293d7ba3d15d64b44e8",
		"fd02d79f1d1d70ca41e5ed62c6d35d2a0f82de15"}

	nodes, edges := visitObjects(objects[:])

	assert.Contains(t, nodes, Node{Id: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5", Type: "commit"})
	assert.Contains(t, edges, Edge{From: "bb4840c0b5dc29bcb7d4e0e2e5d1b9e9dec721e5", To: "c932e6ae3a32d3cdea9dd8043c165e52495ea4c9"})
}
