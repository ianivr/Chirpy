package server

import "sync/atomic"

type apiConfig struct {
	fileserverHits atomic.Int32
}
