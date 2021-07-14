package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type Blockchain struct {
	Chain              []Block
	CurrentTransaction Transaction
	Nodes              []string
}

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

type Block struct {
	Index        int         `json:"index"`
	Timestamp    int64       `json:"timestamp"`
	Transactions Transaction `json:"transactions"`
	Proof        int         `json:"proof"`
	PreviousHash string      `json:"previous_hash"`
}

type FullChain struct {
	Chain  []Block `json:"chain"`
	Length int     ` json:"length"`
}

var blockchain Blockchain

func (Blockchain) Init() {
	//# ジェネシスブロックを作る
	blockchain = Blockchain{}
	blockchain.NewBlock("1", 100)
}

/*
   ブロックチェーンに新しいブロックを作る
   :param proof: <int> プルーフ・オブ・ワークアルゴリズムから得られるプルーフ
   :param previous_hash: (オプション) <str> 前のブロックのハッシュ
   :return: <dict> 新しいブロック
*/
func (Blockchain) NewBlock(previousHash string, proof int) Block {
	pg := ""
	if previousHash != "" {
		pg = previousHash
	} else {
		pg = blockchain.Hash(blockchain.Chain[len(blockchain.Chain)-1])
	}
	block := Block{
		Index:        len(blockchain.Chain) + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: blockchain.CurrentTransaction,
		Proof:        proof,
		PreviousHash: pg, // or blockchain.Hash(blockchain.Chain[-1])
	}
	blockchain.CurrentTransaction = Transaction{}
	blockchain.Chain = append(blockchain.Chain, block)
	return block
}

/*
   次に採掘されるブロックに加える新しいトランザクションを作る
   :param sender: <str> 送信者のアドレス
   :param recipient: <str> 受信者のアドレス
   :param amount: <int> 量
   :return: <int> このトランザクションを含むブロックのアドレス
*/
func (Blockchain) NewTransaction(sender string, recipient string, amount int) int {
	// 新しいトランザクションをリストに加える
	blockchain.CurrentTransaction = Transaction{Sender: sender, Recipient: recipient, Amount: amount}
	return blockchain.LastBlock().Index + 1
}

/*
   ブロックの　SHA-256　ハッシュを作る
   :param block: <dict> ブロック
   :return: <str>
*/
func (Blockchain) Hash(block Block) string {
	// ブロックをハッシュ化する
	blockJson, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}
	converted := sha256.Sum256([]byte(blockJson))
	return hex.EncodeToString(converted[:])
}

func (Blockchain) LastBlock() Block {
	// チェーンの最後のブロックをリターンする
	return blockchain.Chain[len(blockchain.Chain)-1]
}

/*
   シンプルなプルーフ・オブ・ワークのアルゴリズム:
   - hash(pp') の最初の4つが0となるような p' を探す
   - p は1つ前のブロックのプルーフ、 p' は新しいブロックのプルーフ
   :param last_proof: <int>
   :return: <int>
*/
func (Blockchain) ProofOfWork(lastProof int) int {
	proof := 0
	for blockchain.ValidProof(lastProof, proof) == false {
		proof += 1
	}

	return proof
}

/*
   プルーフが正しいかを確認する: hash(last_proof, proof)の最初の4つが0となっているか？
   :param last_proof: <int> 前のプルーフ
   :param proof: <int> 現在のプルーフ
   :return: <bool> 正しければ true 、そうでなれけば false
*/
func (Blockchain) ValidProof(lastProof int, proof int) bool {
	guess := []byte(strconv.Itoa(lastProof) + strconv.Itoa(proof))
	sha256s := sha256.Sum256(guess)
	guessHash := hex.EncodeToString(sha256s[:])
	return guessHash[:4] == "0000"
}

/*
   ブロックチェーンが正しいかを確認する
   :param chain: <list> ブロックチェーン
   :return: <bool> True であれば正しく、 False であればそうではない
*/
func (Blockchain) ValidChain(chain []Block) bool {
	lastBlock := chain[0]
	currentIndex := 1

	for currentIndex < len(chain) {
		block := chain[currentIndex]
		//# ブロックのハッシュが正しいかを確認
		if block.PreviousHash != blockchain.Hash(lastBlock) {
			return false
		}
		//# プルーフ・オブ・ワークが正しいかを確認
		if !blockchain.ValidProof(lastBlock.Proof, block.Proof) {
			return false
		}
		lastBlock = block
		currentIndex += 1
	}
	return true
}

/*
   これがコンセンサスアルゴリズムだ。ネットワーク上の最も長いチェーンで自らのチェーンを
   置き換えることでコンフリクトを解消する。
   :return: <bool> 自らのチェーンが置き換えられると True 、そうでなれけば False
*/
func (Blockchain) ResolveConflicts() bool {
	neighbours := blockchain.Nodes
	var newChain []Block

	//# 自らのチェーンより長いチェーンを探す必要がある
	maxLength := len(blockchain.Chain)

	for _, node := range neighbours {
		response, err := http.Get(node + "/chain")
		if err != nil {
			panic(err)
		}
		if response.StatusCode != 200 {
			panic(err)
		}
		var fullChain FullChain
		if err := json.NewDecoder(response.Body).Decode(&fullChain); err != nil {
			panic(err)
			continue
		}
		length := fullChain.Length
		chain := fullChain.Chain

		if length > maxLength && blockchain.ValidChain(chain) {
			maxLength = length
			newChain = chain
		}
	}
	if len(newChain) != 0 {
		blockchain.Chain = newChain
		return true
	}
	return false
}

/*
   ノードリストに新しいノードを加える
   :param address: <str> ノードのアドレス 例: 'http://192.168.0.5:5000'
   :return: None
*/
func (Blockchain) RegisterNode(address string) {
	blockchain.Nodes = append(blockchain.Nodes, address)

	fix := make(map[string]bool)
	one := []string{}
	for _, a := range blockchain.Nodes {
		if !fix[a] {
			fix[a] = true
			one = append(one, a)
		}
	}
	blockchain.Nodes = one
}

var NodeIdentifire string

func main() {
	// echo server
	e := echo.New()

	//# このノードのグローバルにユニークなアドレスを作る
	NodeIdentifire = strings.Replace(uuid.New().String(), "-", "", -1)

	blockchain.Init()

	e.GET("/mine", Mine)
	e.POST("/transactions/new", NewTransactionPost)
	e.GET("/chain", FullChainGET)

	e.POST("/nodes/register", RegisterNode)
	e.GET("/nodes/resolve", Consensus)

	go func(echoEcho *echo.Echo) {
		copyEcho := echoEcho
		copyEcho.Start(":5001")
	}(e)
	e.Start(":5000")
}

type Post2 struct {
	Nodes []string `json:"nodes"`
}

type Response2 struct {
	Message   string   `json:"message"`
	TotalNode []string `json:"total_node"`
}

func RegisterNode(e echo.Context) error {
	nodes := new(Post2)
	if err := e.Bind(nodes); err != nil {
		return e.JSON(http.StatusBadRequest, "Status Bad Request.")
	}

	for _, node := range nodes.Nodes {
		blockchain.RegisterNode(node)
	}

	var response2 Response2
	response2.Message = "新しいノードが追加されました"
	response2.TotalNode = blockchain.Nodes

	return e.JSON(http.StatusCreated, response2)
}

func Consensus(e echo.Context) error {
	replaced := blockchain.ResolveConflicts()
	if replaced {
		type Response struct {
			Message  string  `json:"message"`
			NewChain []Block `json:"new_chain"`
		}
		var response Response
		response.Message = "チェーンが置き換えられました"
		response.NewChain = blockchain.Chain
		return e.JSON(http.StatusOK, response)
	} else {
		type Response struct {
			Message string  `json:"message"`
			Chain   []Block `json:"chain"`
		}
		var response Response
		response.Message = "チェーンが確認されました"
		response.Chain = blockchain.Chain
		return e.JSON(http.StatusOK, response)
	}

}

type Post struct {
	Sender    string
	Recipient string
	Amount    int
}

//# メソッドはPOSTで/transactions/newエンドポイントを作る。メソッドはPOSTなのでデータを送信する
func NewTransactionPost(e echo.Context) error {
	post := new(Post)
	if err := e.Bind(post); err != nil {
		return e.JSON(http.StatusBadRequest, "Status Bad Request.")
	}

	// '新しいトランザクションを追加します'
	index := blockchain.NewTransaction(post.Sender, post.Recipient, post.Amount)
	return e.JSON(http.StatusCreated, "トランザクションはブロック"+strconv.Itoa(index)+"に追加されました")
}

//# メソッドはGETで/mineエンドポイントを作る
func Mine(e echo.Context) error {
	// '新しいブロックを採掘します'
	lastBlock := blockchain.LastBlock()
	lastProof := lastBlock.Proof
	proof := blockchain.ProofOfWork(lastProof)

	blockchain.NewTransaction("0", NodeIdentifire, 1)

	block := blockchain.NewBlock("", proof)

	response := struct {
		Message      string      `json:"message"`
		Index        int         `json:"index"`
		Transactions Transaction `json:"transactions"`
		Proof        int         `json:"proof"`
		PreviousHash string      `json:"previous_hash"`
	}{Message: "新しいブロックを採掘しました。",
		Index:        block.Index,
		Transactions: block.Transactions,
		Proof:        block.Proof,
		PreviousHash: block.PreviousHash}

	return e.JSON(http.StatusCreated, response)
}

//# メソッドはGETで、フルのブロックチェーンをリターンする/chainエンドポイントを作る
func FullChainGET(e echo.Context) error {
	var response FullChain
	response.Chain = blockchain.Chain
	response.Length = len(blockchain.Chain)

	return e.JSON(http.StatusOK, response)
}
