package main

import (
	"context"
	"fmt"
	"time"

	// "time"

	"github.com/go-redis/redis/v8"
)

type Node struct {
	val         int
	left, right *Node
}

var Ctx = context.Background()

var rdb1 *redis.Client
var rdb2 *redis.Client



func main() {
	rdb1 = redis.NewClient(&redis.Options{
		Addr:     "121.5.231.228:6379",
		Password: "tomaChen513?",
		DB:       0, //  选择将点赞视频id信息存入 DB0.
	})
	
	rdb2 = redis.NewClient(&redis.Options{
		Addr:     "121.5.231.228:6379",
		Password: "tomaChen513?",
		DB:       1, //  选择将点赞视频id信息存入 DB0.
	})

	rdb1.Set(Ctx, "at_123", "1", 20*time.Second)
	rdb2.Set(Ctx, "at_123", "rt_123", 40*time.Second)
	ans:=rdb1.Get(Ctx,"at_123")
	fmt.Println(ans)
	fmt.Println("pass")
	time.Sleep(1)
}

func request(at string){
	valid:=rdb1.Get(Ctx,"at_123")
	if valid==nil{
		// search db2

	}
	
}

// func TestSetKey(t *testing.T) {
// 	var Ctx = context.Background()

// 	RdbLikeUserId = redis.NewClient(&redis.Options{
// 		Addr:     "121.5.231.228:6379",
// 		Password: "wintercamp",
// 		DB:       0, //  选择将点赞视频id信息存入 DB0.
// 	})

// 	RdbLikeUserId.Set(Ctx,"toma","cxy1",20*time.Second)

// }

// 最近公共祖先
func lowestAncestor(root, p, q *Node) *Node {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestAncestor(root.left, p, q)
	right := lowestAncestor(root.right, p, q)

	if left != nil && right != nil {
		return root
	} else if left == nil && right != nil {
		return right
	} else if left != nil && right == nil {
		return left
	} else {
		return nil
	}

}

func lowestAncestor2(root, p, q *Node) *Node {
	if root.val > p.val && root.val > q.val {
		return lowestAncestor2(root.left, p, q)
	} else if root.val < p.val && root.val < q.val {
		return lowestAncestor2(root.right, p, q)
	} else {
		return root
	}
}
