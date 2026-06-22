package structs

import (
	"sync"
	"time"
)


type TokenBucket struct{
	Qtd		int
	Max		int
	mu		sync.Mutex
}

func (b *TokenBucket) Consume() bool{
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Qtd <= 0{
		return false
	}
	b.Qtd--

	return  true
}

func (b *TokenBucket) InsertToken(){
	for{
		time.Sleep(1 * time.Second)
		b.mu.Lock()
		if b.Qtd < b.Max{
			b.Qtd ++
		}
		b.mu.Unlock()
	}
}