package gobotapi

import (
	"bytes"
	"context"
	"fmt"
	rawTypes "gobotapi/types/raw"
	"gobotapi/utils/proxy_reader"
	"io"
	"mime/multipart"
	"net/http"
)

func (ctx *Client) executeRequest(url, method string, data map[string]string, files map[string]rawTypes.InputFile, callable rawTypes.ProgressCallable) ([]byte, error) {
	var body io.Reader = &bytes.Buffer{}
	var multiPartWriter *multipart.Writer
	if len(data) > 0 || len(files) > 0 {
		reader := &bytes.Buffer{}
		multiPartWriter = multipart.NewWriter(reader)
		for k, v := range data {
			_ = multiPartWriter.WriteField(k, v)
		}
		for k, v := range files {
			file, err := multiPartWriter.CreateFormFile(k, v.FileName())
			if err != nil {
				return nil, err
			}
			_, err = file.Write(v.Content())
			if err != nil {
				return nil, err
			}
		}
		_ = multiPartWriter.Close()
		body = proxy_reader.NewProxyReader(reader, ctx.DownloadRefreshRate, int64(reader.Len()), callable)
		defer body.(*proxy_reader.ProxyReader).Close()
	}
	ctxConn, cancelable := context.WithCancel(context.Background())
	cancelableContext := rawTypes.CancelableContext{
		Cancel: cancelable,
	}
	cancelableContext.GenerateId()
	defer func() {
		for i, x := range ctx.requestsContext {
			if x.Id == cancelableContext.Id {
				ctx.requestsContext = append(ctx.requestsContext[:i], ctx.requestsContext[i+1:]...)
				break
			}
		}
	}()
	req, err := http.NewRequestWithContext(ctxConn, method, url, body)
	ctx.requestsContext = append(ctx.requestsContext, cancelableContext)
	if err != nil {
		return nil, err
	}
	if multiPartWriter != nil {
		req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	}
	do, err := ctx.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(do.Body)
	var buf bytes.Buffer
	var proxyBody io.Reader = do.Body
	if method == "GET" {
		proxyBody = proxy_reader.NewProxyReader(do.Body, ctx.DownloadRefreshRate, do.ContentLength, callable)
		defer proxyBody.(*proxy_reader.ProxyReader).Close()
	}
	_, err = io.Copy(&buf, proxyBody)
	if err != nil {
		return nil, err
	}
	if do.StatusCode != http.StatusOK && do.StatusCode != http.StatusCreated && do.StatusCode != http.StatusNoContent {
		return buf.Bytes(), fmt.Errorf("error code %d", do.StatusCode)
	}
	return buf.Bytes(), nil
}
