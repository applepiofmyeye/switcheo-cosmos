package keeper

import (
	"context"

	"music/x/music/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPlaylist(goCtx context.Context, req *types.QueryListPlaylistRequest) (*types.QueryListPlaylistResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SongKey))

	var songs []*types.Song
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var song types.Song
		if err := k.cdc.Unmarshal(value, &song); err != nil {
			return err
		}

		// Check if the song's playlist matches the requested playlist ID
		if song.Playlist == req.Playlist {
			songs = append(songs, &song)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListPlaylistResponse{Song: songs, Pagination: pageRes}, nil
}
