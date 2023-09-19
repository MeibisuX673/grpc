package directoryInfo

import (
	"context"
	infoDirectory "github.com/MeibisuX673/grpc/server/proto/infoDirectory"
	"io"
	"os"
	"sync"
)

var cache = make(map[string]*infoDirectory.DirInfoResponse)

type DirectoryInfo struct {
	mu sync.Mutex // protects routeNot
	infoDirectory.UnimplementedInfoDirectoryServer
}

func (d *DirectoryInfo) InfoDirStreamAll(stream infoDirectory.InfoDirectory_InfoDirStreamAllServer) error {

	var response *infoDirectory.DirInfoResponse

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		infoDir, err := getDirInfo(in)
		if err != nil {
			return nil
		}
		response = infoDir

		if err := stream.Send(response); err != nil {
			return err
		}
	}

}

func (d *DirectoryInfo) InfoDirStreamClient(stream infoDirectory.InfoDirectory_InfoDirStreamClientServer) error {

	var response []*infoDirectory.DirInfoResponse

	for {
		path, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&infoDirectory.DirInfoStreamClientResponse{
				DirectoriesInfo: response,
			})
		}
		if err != nil {
			return err
		}
		dirInfo, err := getDirInfo(path)
		if err != nil {
			return err
		}
		response = append(response, dirInfo)
	}

}

func (d *DirectoryInfo) InfoDir(ctx context.Context, request *infoDirectory.PathRequest) (*infoDirectory.DirInfoResponse, error) {

	response, err := getDirInfo(request)
	if err != nil {
		return nil, err
	}

	return response, err

}

func getDirInfo(pathReq *infoDirectory.PathRequest) (*infoDirectory.DirInfoResponse, error) {

	var response infoDirectory.DirInfoResponse
	var sumSize int64 = 0

	if len(cache) > 20 {
		cache = make(map[string]*infoDirectory.DirInfoResponse)
	}

	dirs, err := os.ReadDir(pathReq.Path)
	if err != nil {
		return nil, err
	}

	value, ok := cache[pathReq.Path]
	if ok {
		if len(dirs) == len(value.Files)+len(value.Directories) {
			return value, nil
		} else {
			cache[pathReq.Path] = nil
		}
	}

	var files []*infoDirectory.FileInfo
	var directories []*infoDirectory.DirInfo

	for _, value := range dirs {

		fileInfo, err := value.Info()
		if err != nil {
			return nil, err
		}

		sumSize += fileInfo.Size()
		fileName := fileInfo.Name()
		fileSize := fileInfo.Size()

		if fileInfo.IsDir() {
			directories = append(directories, &infoDirectory.DirInfo{
				Name: fileName,
				Size: fileSize,
			})
			continue
		}
		files = append(files, &infoDirectory.FileInfo{
			Name: fileName,
			Size: fileSize,
		})

	}

	response.Size = sumSize
	response.Directories = directories
	response.Files = files
	response.Path = pathReq.GetPath()

	cache[pathReq.Path] = &response

	return &response, err

}
