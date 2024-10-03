package music

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "music/api/music/music"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowSong",
					Use:            "show-song [id]",
					Short:          "Query show-song",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListPlaylist",
					Use:            "list-playlist [playlist]",
					Short:          "Query list-playlist",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "playlist"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateSong",
					Use:            "create-song [title]",
					Short:          "Send a create-song tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}},
				},
				{
					RpcMethod:      "UpdateSong",
					Use:            "update-song [title] [id]",
					Short:          "Send a update-song tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteSong",
					Use:            "delete-song [id]",
					Short:          "Send a delete-song tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreatePlaylist",
					Use:            "create-playlist [title] [id]",
					Short:          "Send a create-playlist tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
