// Auto-generated by avdl-compiler v1.3.19 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/git.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type EncryptedGitMetadata struct {
	V   int                  `codec:"v" json:"v"`
	E   []byte               `codec:"e" json:"e"`
	N   BoxNonce             `codec:"n" json:"n"`
	Gen PerTeamKeyGeneration `codec:"gen" json:"gen"`
}

func (o EncryptedGitMetadata) DeepCopy() EncryptedGitMetadata {
	return EncryptedGitMetadata{
		V: o.V,
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.E),
		N:   o.N.DeepCopy(),
		Gen: o.Gen.DeepCopy(),
	}
}

type RepoID string

func (o RepoID) DeepCopy() RepoID {
	return o
}

type GitLocalMetadata struct {
	RepoName GitRepoName `codec:"repoName" json:"repoName"`
}

func (o GitLocalMetadata) DeepCopy() GitLocalMetadata {
	return GitLocalMetadata{
		RepoName: o.RepoName.DeepCopy(),
	}
}

type GitServerMetadata struct {
	Ctime                 Time     `codec:"ctime" json:"ctime"`
	Mtime                 Time     `codec:"mtime" json:"mtime"`
	LastModifyingUsername string   `codec:"lastModifyingUsername" json:"lastModifyingUsername"`
	LastModifyingDeviceID DeviceID `codec:"lastModifyingDeviceID" json:"lastModifyingDeviceID"`
}

func (o GitServerMetadata) DeepCopy() GitServerMetadata {
	return GitServerMetadata{
		Ctime: o.Ctime.DeepCopy(),
		Mtime: o.Mtime.DeepCopy(),
		LastModifyingUsername: o.LastModifyingUsername,
		LastModifyingDeviceID: o.LastModifyingDeviceID.DeepCopy(),
	}
}

type GitRepoResult struct {
	Folder         Folder            `codec:"folder" json:"folder"`
	RepoID         RepoID            `codec:"repoID" json:"repoID"`
	LocalMetadata  GitLocalMetadata  `codec:"localMetadata" json:"localMetadata"`
	ServerMetadata GitServerMetadata `codec:"serverMetadata" json:"serverMetadata"`
}

func (o GitRepoResult) DeepCopy() GitRepoResult {
	return GitRepoResult{
		Folder:         o.Folder.DeepCopy(),
		RepoID:         o.RepoID.DeepCopy(),
		LocalMetadata:  o.LocalMetadata.DeepCopy(),
		ServerMetadata: o.ServerMetadata.DeepCopy(),
	}
}

type PutGitMetadataArg struct {
	Folder   Folder           `codec:"folder" json:"folder"`
	RepoID   RepoID           `codec:"repoID" json:"repoID"`
	Metadata GitLocalMetadata `codec:"metadata" json:"metadata"`
}

func (o PutGitMetadataArg) DeepCopy() PutGitMetadataArg {
	return PutGitMetadataArg{
		Folder:   o.Folder.DeepCopy(),
		RepoID:   o.RepoID.DeepCopy(),
		Metadata: o.Metadata.DeepCopy(),
	}
}

type GetGitMetadataArg struct {
	Folder Folder `codec:"folder" json:"folder"`
}

func (o GetGitMetadataArg) DeepCopy() GetGitMetadataArg {
	return GetGitMetadataArg{
		Folder: o.Folder.DeepCopy(),
	}
}

type GetAllGitMetadataArg struct {
}

func (o GetAllGitMetadataArg) DeepCopy() GetAllGitMetadataArg {
	return GetAllGitMetadataArg{}
}

type GitInterface interface {
	PutGitMetadata(context.Context, PutGitMetadataArg) error
	GetGitMetadata(context.Context, Folder) ([]GitRepoResult, error)
	GetAllGitMetadata(context.Context) ([]GitRepoResult, error)
}

func GitProtocol(i GitInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.git",
		Methods: map[string]rpc.ServeHandlerDescription{
			"putGitMetadata": {
				MakeArg: func() interface{} {
					ret := make([]PutGitMetadataArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PutGitMetadataArg)
					if !ok {
						err = rpc.NewTypeError((*[]PutGitMetadataArg)(nil), args)
						return
					}
					err = i.PutGitMetadata(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getGitMetadata": {
				MakeArg: func() interface{} {
					ret := make([]GetGitMetadataArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetGitMetadataArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetGitMetadataArg)(nil), args)
						return
					}
					ret, err = i.GetGitMetadata(ctx, (*typedArgs)[0].Folder)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getAllGitMetadata": {
				MakeArg: func() interface{} {
					ret := make([]GetAllGitMetadataArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetAllGitMetadata(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type GitClient struct {
	Cli rpc.GenericClient
}

func (c GitClient) PutGitMetadata(ctx context.Context, __arg PutGitMetadataArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.git.putGitMetadata", []interface{}{__arg}, nil)
	return
}

func (c GitClient) GetGitMetadata(ctx context.Context, folder Folder) (res []GitRepoResult, err error) {
	__arg := GetGitMetadataArg{Folder: folder}
	err = c.Cli.Call(ctx, "keybase.1.git.getGitMetadata", []interface{}{__arg}, &res)
	return
}

func (c GitClient) GetAllGitMetadata(ctx context.Context) (res []GitRepoResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.git.getAllGitMetadata", []interface{}{GetAllGitMetadataArg{}}, &res)
	return
}
