package file_chunk

import (
	"io"
	"os"
)

const bufSize = 1024

// OffsetRange represents a content block of a file.
type OffsetRange struct {
	FileName string
	Start    int64
	End      int64
}

//SplitLineChunks splits file into chunks.
func SplitLineChunks(filename string, chunks int) ([]OffsetRange, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	if chunks <= 1 {
		return []OffsetRange{
			{
				FileName: filename,
				Start:    0,
				End:      info.Size(),
			},
		}, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var ranges []OffsetRange
	var offset int64
	// avoid the last chunk too few bytes
	preferSize := info.Size()/int64(chunks) + 1
	for {
		if offset+preferSize >= info.Size() {
			ranges = append(ranges, OffsetRange{
				filename,
				offset,
				info.Size(),
			})
			break
		}

		offsetRange, err := nextRange(file, offset, offset+preferSize)
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, offsetRange)
		if offsetRange.End < info.Size() {
			offset = offsetRange.End
		} else {
			break
		}
	}
	return ranges, nil
}

func nextRange(file *os.File, start, stop int64) (OffsetRange, error) {
	offset, err := skipPartialLine(file, stop)
	if err != nil {
		return OffsetRange{}, err
	}

	return OffsetRange{
		FileName: file.Name(),
		Start:    start,
		End:      offset,
	}, nil
}

func skipPartialLine(file *os.File, offset int64) (int64, error) {
	for {
		skipBuf := make([]byte, bufSize)
		n, err := file.ReadAt(skipBuf, offset)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if n == 0 {
			return 0, io.EOF
		}
		for i := 0; i < n; i++ {
			if skipBuf[i] != '\r' && skipBuf[i] != '\n' {
				offset++
			}
		}
		return offset, nil
	}
}
