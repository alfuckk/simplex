package task

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"

	"github.com/spf13/viper"
)

type GeoIPTask interface {
	DownloadToLocal(ctx context.Context) error
	downloadFile(URL, filePath string, resultChan chan<- DownloadResult) error
}

func NewGeoIPTask(
	task *Task,
	conf *viper.Viper,
	// userRepo repo.UserRepository,
) GeoIPTask {
	return &geoipTask{
		// userRepo: userRepo,
		conf: conf,
		Task: task,
	}
}

type geoipTask struct {
	conf *viper.Viper
	*Task
}

type DownloadResult struct {
	FileName string
	Success  bool
	Error    error
}

// TODO 按照日期来算文件夹，要入数据库记录。
func (g *geoipTask) DownloadToLocal(ctx context.Context) error {
	urls := g.conf.GetStringSlice("geoip.urls")
	results := make(chan DownloadResult, len(urls))
	var wg sync.WaitGroup

	for _, v := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			pathName, _ := getFileNameFromURL(v)

			path := g.conf.GetString("geoip.data_path") + "/" + pathName
			g.downloadFile(v, path, results)
		}(v)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.Success {
			g.logger.Sugar().Infof("✅ 文件 '%s' 下载成功\n", res.FileName)
		} else {
			g.logger.Sugar().Infof("❌ 文件 '%s' 下载失败: %v\n", res.FileName, res.Error)
		}
	}
	return nil
}

func getFileNameFromURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("解析 URL 失败: %w", err)
	}

	fileName := path.Base(parsedURL.Path)

	if fileName == "" || fileName == "/" {
		return "", fmt.Errorf("URL 中没有明确的文件名: %s", rawURL)
	}

	return fileName, nil
}

func (g *geoipTask) downloadFile(URL, filePath string, resultChan chan<- DownloadResult) error {
	resp, err := http.Get(URL)
	if err != nil {
		g.logger.Sugar().Errorf("发送请求失败%v", err)
		return fmt.Errorf("发送请求失败%w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		g.logger.Sugar().Errorf("下载失败，HTTP 状态码: %d %s", resp.StatusCode, resp.Status)
		return fmt.Errorf("下载失败，HTTP 状态码: %d %s", resp.StatusCode, resp.Status)
	}

	out, err := os.Create(filePath)
	if err != nil {
		g.logger.Sugar().Errorf("创建本地文件失败: %v", err)
		return fmt.Errorf("创建本地文件失败: %w", err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		g.logger.Sugar().Errorf("写入文件失败: %v", err)
		return fmt.Errorf("写入文件失败: %w", err)
	}
	g.logger.Sugar().Infof("文件 '%s' 下载成功！\n", filePath)
	resultChan <- DownloadResult{FileName: filePath, Success: true, Error: nil}
	return nil
}
