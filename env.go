package dotenv

import "os"

func LoadFile(path string, env *map[string]string) (err error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return
	}

	r := newReader(fileContent)

	map_ := make(map[string]string)

	for r.bookmark < r.size {
		idx := string(r.readIdx())

		if !r.exceptionFounded {
			continue
		}

		val := string(r.readVal())

		map_[idx] = val
	}

	*env = map_

	return
}
