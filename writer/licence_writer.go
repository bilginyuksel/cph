package writer

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

const LICENCE = `/*
    Copyright 2020. Huawei Technologies Co., Ltd. All rights reserved.

    Licensed under the Apache License, Version 2.0 (the "License")
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        https://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/
`

func Run(files []string) {
	for _, file := range files {
		WriteLicence(file)
	}
}

func WriteLicence(file string) {
	filename := filepath.Base(file)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(LICENCE)); err != nil {
		log.Fatal(err)
	}
	javaFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := javaFile.Close(); err != nil {
			panic(err)
		}
	}()
	buf := make([]byte, 1024)
	for {
		n, err := javaFile.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := f.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
