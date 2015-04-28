package file

import "os"

func FindFirstExistedFile(files []string) string {
	for _, conf := range confs {
		_, err := os.Stat(conf)
		if err == nil {
			return conf
		}
	}
	return ""
}
