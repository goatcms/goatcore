package loop

import "github.com/goatcms/goat-core/filesystem"

type producerBody struct {
	fs          filesystem.Filespace
	filter      filesystem.LoopFilter
	chans       *chans
	ignoreFiles bool
}

func (j producerBody) Step() (bool, error) {
	select {
	case row := <-j.chans.baseChan:
		list, err := j.fs.ReadDir(row)
		if err != nil {
			return false, err
		}
		for _, childNode := range list {
			if childNode.Name() == "." || childNode.Name() == ".." {
				continue
			}
			childPath := row + childNode.Name()
			if childNode.IsDir() {
				if j.filter != nil && !j.filter(j.fs, childPath) {
					continue
				}
				exChildPath := childPath + "/"
				j.chans.baseChan <- exChildPath
				j.chans.dirChan <- exChildPath
			} else {
				if j.ignoreFiles == false {
					if j.filter != nil && !j.filter(j.fs, childPath) {
						continue
					}
					j.chans.fileChan <- childPath
				}
			}
		}
	default:
		return false, nil
	}
	return true, nil
}
