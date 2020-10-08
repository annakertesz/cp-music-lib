package updater

import (
	"fmt"
	"github.com/dhowden/tag"
	"os"
)

func getTags(file *os.File){
	//metadata, _ := tag.ReadFrom(file)
}

func a(){
	file, _ := os.Open("../sources/instr.mp3")
	metadata, _ := tag.ReadFrom(file)
	fmt.Println(metadata.Album())
	fmt.Println(metadata.Genre())
}