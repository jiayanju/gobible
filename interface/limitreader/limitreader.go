package limitreader

import (
	"io"
)

//LimitReader returns a reader
func LimitReader(r io.Reader, n int64) io.Reader  {
    return &LimitedReader{r, n}
}

//LimitedReader limit reader as with LimitReader in io package
type LimitedReader struct {
    R io.Reader
    N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error)  {
    if l.N <= 0 {
        return 0, io.EOF
    }
    
    if int64(len(p)) > l.N {
        p = p[0:l.N]
    }
    
    n, err = l.R.Read(p)
    l.N -= int64(n)
    return
}