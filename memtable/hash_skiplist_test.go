package memtable

import (
	"github.com/flowercorp/lotusdb/logfile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashSkipList_Put(t *testing.T) {
	var entry = &logfile.LogEntry{
		Key:   []byte("lotusdb"),
		Value: []byte("lotusdb"),
	}

	hashlist := NewHashSkipList()

	actEntry := hashlist.Put([]byte("lotusdb"), []byte("lotusdb"))
	require.Equal(t, entry, actEntry)
}

func TestHashSkipList_Get(t *testing.T) {
	var entry = &logfile.LogEntry{
		Key:   []byte("lotusdb"),
		Value: []byte("lotusdb"),
	}

	hashlist := NewHashSkipList()

	hashlist.Put([]byte("lotusdb"), []byte("lotusdb"))

	actEntry := hashlist.Get([]byte("lotusdb"))
	require.Equal(t, entry, actEntry)
}

func TestHashSkipList_Exist(t *testing.T) {
	var entry = &logfile.LogEntry{
		Key:   []byte("lotusdb"),
		Value: []byte("lotusdb"),
	}

	hashlist := NewHashSkipList()

	hashlist.Put([]byte("lotusdb"), []byte("lotusdb"))

	actEntry := hashlist.Get([]byte("lotusdb"))
	require.Equal(t, entry, actEntry)
	assert.True(t, hashlist.Exist([]byte("lotusdb")))

	hashlist.Remove([]byte("lotusdb"))
	assert.False(t, hashlist.Exist([]byte("lotusdb")))
}

func TestHashSkipList_Remove(t *testing.T) {
	var entry = &logfile.LogEntry{
		Key:   []byte("lotusdb"),
		Value: []byte("lotusdb"),
	}

	hashlist := NewHashSkipList()

	hashlist.Put([]byte("lotusdb"), []byte("lotusdb"))

	actEntry := hashlist.Get([]byte("lotusdb"))
	require.Equal(t, entry, actEntry)

	hashlist.Remove([]byte("lotusdb"))
	assert.Nil(t, hashlist.Get([]byte("lotusdb")))
}
