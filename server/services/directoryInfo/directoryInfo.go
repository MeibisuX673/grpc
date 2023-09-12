package directoryInfo

import (
	"context"
	"github.com/MeibisuX673/grpc/proto/infoDirectory"
	"os"
	"sync"
)

var cache = make(map[string]*infoDirectory.DirInfoResponse)

type DirectoryInfo struct {
	mutex sync.Mutex
	infoDirectory.UnimplementedInfoDirectoryServer
}

func (d *DirectoryInfo) InfoDir(ctx context.Context, request *infoDirectory.PathRequest) (*infoDirectory.DirInfoResponse, error) {

	var response infoDirectory.DirInfoResponse
	var sumSize int64 = 0

	if len(cache) > 20 {
		cache = make(map[string]*infoDirectory.DirInfoResponse)
	}

	value, ok := cache[request.Path]
	if ok {
		return value, nil
	}

	dirs, err := os.ReadDir(request.Path)
	if err != nil {
		return nil, err
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

	cache[request.Path] = &response

	return &response, err
}
