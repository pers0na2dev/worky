package main

import (
	"github.com/kaazedev/worky"
)

func main() {
	pool := worky.New(5)

	for i := 0; i < 10; i++ {
		pool.Add(func() worky.Result {
			return worky.Result{
				Result: "Task done",
				Err:    nil,
				Meta: worky.Meta{
					"some_id": 1,
				},
			}
		})
	}

	go pool.Start()

	go func() {
		for result := range pool.Results() {
			if result.Err != nil {
				panic(result.Err)
			}
			println(result.Result.(string))
			println(result.Meta["some_id"].(int))
		}
	}()

	for {
	}
}
