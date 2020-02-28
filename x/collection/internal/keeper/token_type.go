package keeper

import (
	"github.com/line/link/x/collection/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TokenTypeKeeper interface {
	GetNextTokenType(ctx sdk.Context, contractID string) (tokenType string, err sdk.Error)
	GetTokenTypes(ctx sdk.Context, contractID string) (types.TokenTypes, sdk.Error)
	GetTokenType(ctx sdk.Context, contractID, tokenType string) (types.TokenType, sdk.Error)
	HasTokenType(ctx sdk.Context, contractID, tokenType string) bool
	SetTokenType(ctx sdk.Context, contractID string, token types.TokenType) sdk.Error
	UpdateTokenType(ctx sdk.Context, contractID string, token types.TokenType) sdk.Error
}

var _ TokenTypeKeeper = (*Keeper)(nil)

func (k Keeper) SetTokenType(ctx sdk.Context, contractID string, tokenType types.TokenType) sdk.Error {
	_, err := k.GetCollection(ctx, contractID)
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.TokenTypeKey(contractID, tokenType.GetTokenType())) {
		return types.ErrTokenTypeExist(types.DefaultCodespace, contractID, tokenType.GetTokenType())
	}
	store.Set(types.TokenTypeKey(contractID, tokenType.GetTokenType()), k.cdc.MustMarshalBinaryBare(tokenType))
	return nil
}

func (k Keeper) UpdateTokenType(ctx sdk.Context, contractID string, tokenType types.TokenType) sdk.Error {
	_, err := k.GetCollection(ctx, contractID)
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.TokenTypeKey(contractID, tokenType.GetTokenType())) {
		return types.ErrTokenTypeNotExist(types.DefaultCodespace, contractID, tokenType.GetTokenType())
	}
	store.Set(types.TokenTypeKey(contractID, tokenType.GetTokenType()), k.cdc.MustMarshalBinaryBare(tokenType))
	return nil
}

func (k Keeper) GetTokenType(ctx sdk.Context, contractID string, tokenTypeID string) (types.TokenType, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	tokenTypeKey := types.TokenTypeKey(contractID, tokenTypeID)
	bz := store.Get(tokenTypeKey)
	if bz == nil {
		return nil, types.ErrTokenTypeNotExist(types.DefaultCodespace, contractID, tokenTypeID)
	}
	tokenType := k.mustDecodeTokenType(bz)
	return tokenType, nil
}

func (k Keeper) GetTokenTypes(ctx sdk.Context, contractID string) (tokenTypes types.TokenTypes, err sdk.Error) {
	_, err = k.GetCollection(ctx, contractID)
	if err != nil {
		return nil, err
	}
	k.iterateTokenTypes(ctx, contractID, "", false, func(t types.TokenType) bool {
		tokenTypes = append(tokenTypes, t)
		return false
	})
	return tokenTypes, nil
}

func (k Keeper) HasTokenType(ctx sdk.Context, contractID, tokenType string) bool {
	store := ctx.KVStore(k.storeKey)
	tokenTypeKey := types.TokenTypeKey(contractID, tokenType)
	return store.Has(tokenTypeKey)
}

func (k Keeper) GetNextTokenType(ctx sdk.Context, contractID string) (tokenType string, err sdk.Error) {
	var lastTokenType types.TokenType
	k.iterateTokenTypes(ctx, contractID, "", true, func(t types.TokenType) bool {
		lastTokenType = t
		return true
	})

	if lastTokenType == nil {
		return types.SmallestNFTType, nil
	}
	tokenType = nextID(lastTokenType.GetTokenType(), "")
	if tokenType[0] == types.FungibleFlag[0] {
		return "", types.ErrTokenTypeFull(types.DefaultCodespace, contractID)
	}
	return tokenType, nil
}

func (k Keeper) iterateTokenTypes(ctx sdk.Context, contractID, prefix string, reverse bool, process func(types.TokenType) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	var iter sdk.Iterator
	if reverse {
		iter = sdk.KVStoreReversePrefixIterator(store, types.TokenTypeKey(contractID, prefix))
	} else {
		iter = sdk.KVStorePrefixIterator(store, types.TokenTypeKey(contractID, prefix))
	}
	defer iter.Close()
	for {
		if !iter.Valid() {
			return
		}
		val := iter.Value()
		tokenType := k.mustDecodeTokenType(val)
		if process(tokenType) {
			return
		}
		iter.Next()
	}
}

func (k Keeper) mustDecodeTokenType(bz []byte) (tokenType types.TokenType) {
	err := k.cdc.UnmarshalBinaryBare(bz, &tokenType)
	if err != nil {
		panic(err)
	}
	return tokenType
}
