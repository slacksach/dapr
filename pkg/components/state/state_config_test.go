package state

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	SaveStateConfiguration("store1", map[string]string{strategyKey: strategyNone})
	SaveStateConfiguration("store2", map[string]string{strategyKey: strategyAppid})
	SaveStateConfiguration("store3", map[string]string{strategyKey: strategyDefault})
	SaveStateConfiguration("store4", map[string]string{strategyKey: strategyStoreName})
	SaveStateConfiguration("store5", map[string]string{strategyKey: "other-fixed-prefix"})
	// if strategyKey not set
	SaveStateConfiguration("store6", map[string]string{})
	os.Exit(m.Run())
}

func TestNonePrefix(t *testing.T) {
	var key = "state-key-1234567"

	modifiedStateKey := GetModifiedStateKey(key, "store1", "appid1")
	require.Equal(t, key, modifiedStateKey)

	originalStateKey := GetOriginalStateKey(modifiedStateKey)
	require.Equal(t, key, originalStateKey)
}

func TestAppidPrefix(t *testing.T) {
	var key = "state-key-1234567"

	modifiedStateKey := GetModifiedStateKey(key, "store2", "appid1")
	require.Equal(t, "appid1||state-key-1234567", modifiedStateKey)

	originalStateKey := GetOriginalStateKey(modifiedStateKey)
	require.Equal(t, key, originalStateKey)
}

func TestAppidPrefix_WithEnptyAppid(t *testing.T) {
	var key = "state-key-1234567"

	modifiedStateKey := GetModifiedStateKey(key, "store2", "")
	require.Equal(t, "state-key-1234567", modifiedStateKey)

	originalStateKey := GetOriginalStateKey(modifiedStateKey)
	require.Equal(t, key, originalStateKey)
}

func TestDefaultPrefix(t *testing.T) {
	var key = "state-key-1234567"

	modifiedStateKey := GetModifiedStateKey(key, "store3", "appid1")
	require.Equal(t, "appid1||state-key-1234567", modifiedStateKey)

	originalStateKey := GetOriginalStateKey(modifiedStateKey)
	require.Equal(t, key, originalStateKey)
}

func TestStoreNamePrefix(t *testing.T) {
	var key = "state-key-1234567"

	modifiedStateKey := GetModifiedStateKey(key, "store4", "appid1")
	require.Equal(t, "store4||state-key-1234567", modifiedStateKey)

	originalStateKey := GetOriginalStateKey(modifiedStateKey)
	require.Equal(t, key, originalStateKey)
}

func TestOtherFixedPrefix(t *testing.T) {
	var key = "state-key-1234567"

	modifiedStateKey := GetModifiedStateKey(key, "store5", "appid1")
	require.Equal(t, "other-fixed-prefix||state-key-1234567", modifiedStateKey)

	originalStateKey := GetOriginalStateKey(modifiedStateKey)
	require.Equal(t, key, originalStateKey)
}

func TestLegacyPrefix(t *testing.T) {
	var key = "state-key-1234567"

	modifiedStateKey := GetModifiedStateKey(key, "store6", "appid1")
	require.Equal(t, "appid1||state-key-1234567", modifiedStateKey)

	originalStateKey := GetOriginalStateKey(modifiedStateKey)
	require.Equal(t, key, originalStateKey)
}

func TestPrefix_StoreNotInitial(t *testing.T) {
	var key = "state-key-1234567"

	// no config for store999
	modifiedStateKey := GetModifiedStateKey(key, "store999", "appid99")
	require.Equal(t, "appid99||state-key-1234567", modifiedStateKey)

	originalStateKey := GetOriginalStateKey(modifiedStateKey)
	require.Equal(t, key, originalStateKey)
}
