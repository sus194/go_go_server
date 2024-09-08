package schedulars

import (
    "sync"
)

var Mu sync.Mutex

func Schedule() {
    Mu.Lock()
    defer Mu.Unlock()
    
}