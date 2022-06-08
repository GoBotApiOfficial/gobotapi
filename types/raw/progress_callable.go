package raw

type ProgressCallable func(downloadedBytes int64, totalBytes int64)
