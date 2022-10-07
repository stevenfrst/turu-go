package repository

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ReadEnv(serviceTags string, tagsName string) {
	readFile, err := os.OpenFile(".env", os.O_RDWR, 0644)
	defer readFile.Close()

	if err != nil {
		log.Fatalln(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var regex, _ = regexp.Compile(tagsName + "=")
	//log.Println(tagsName + "=")
	var lines = []string{"\n", "\nHAPUS\n"}
	w := bufio.NewWriter(readFile)
	//defer w.Flush()
	var foundTag = false
	for fileScanner.Scan() {
		if regex.MatchString(fileScanner.Text()) {
			s := strings.Split(fileScanner.Text(), "=")
			//log.Println(s[1],"dapet")
			modifiedString := s[0] + "=" + serviceTags + "\n"
			lines = append(lines, modifiedString)
			log.Println(modifiedString)
			foundTag = true
			//panic(s[1])
			continue
		}
		lines = append(lines, fileScanner.Text()+"\n")
	}
	if !foundTag {
		log.Println("Tag not found")
		return
	}
	var writeString int
	for x := range lines {
		writeString, err = w.WriteString(lines[x])
		if err != nil {
			return
		}
	}
	w.Flush()
	//log.Println(lines)
	err = exec.Command("bash", "-c", "sed '1,/HAPUS/d' .env > .env.temp && mv .env.temp .env").Run()
	if err != nil {
		log.Println(err)
	}

	log.Println("Success Write", writeString)
}
