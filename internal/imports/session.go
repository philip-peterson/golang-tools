package imports

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/internal/lru"
)

func NewSession(cacheSizeMB uint) *Session {
	fset := token.NewFileSet()
	cache := lru.New(int(cacheSizeMB))
	return &Session{
		Fset:  fset,
		Cache: cache,
	}
}

type Session struct {
	Cache *lru.Cache
	Fset  *token.FileSet
}

type CacheEntry struct {
	astFile *ast.File
	err     error
}

func (s *Session) MemoizedParseFile(filePath string, rawSrc any, mode parser.Mode) (*ast.File, error) {
	src := []byte{}
	if rawSrc != nil {
		src = rawSrc.([]byte)
	}
	key := fmt.Sprintf("%d\x00%s\x00%s\x00", mode, filePath, src)
	cacheEntry := s.Cache.Get(key)
	if cacheEntry != nil {
		// Cache hit
		entry := cacheEntry.(CacheEntry)
		return entry.astFile, entry.err
	}

	// Cache miss
	sizeMB := len(key) / (1024 * 1024)
	if sizeMB < 1 {
		sizeMB = 1
	}
	f, err := parser.ParseFile(s.Fset, filePath, rawSrc, mode)
	if err != nil {
		s.Cache.Set(key, CacheEntry{
			astFile: nil,
			err:     err,
		}, sizeMB)
		return nil, err
	}

	s.Cache.Set(key, CacheEntry{
		astFile: f,
		err:     nil,
	}, sizeMB)
	return f, nil
}
