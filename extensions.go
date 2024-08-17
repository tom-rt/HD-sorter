package main

var pictureExt []string = []string{
	".jpg",
	".jpeg",
	".png",
	".gif",
	".JPG",
	".webp",
	".jpe",
	".svg",
}

var videoExt []string = []string{
	".mov",
	".wav",
	".mp4",
	".MP4",
	".MOV",
	"m4a",
}

var docExt []string = []string{
	".pdf",
	".gpx",
}

var audioExt []string = []string{
	".mp3",
	".flac",
}

func checkType(ext string, extChecked []string) bool {
	for _, i := range extChecked {
		if ext == i {
			return true
		}
	}
	return false
}

func toSort(name string, path string) bool {
	fullPath := path + name
	if name[0] == '.' ||
		fullPath == ".pictures" ||
		fullPath == ".videos" ||
		fullPath == ".documents" ||
		fullPath == ".misc" ||
		fullPath == ".audios" {
		return false
	}
	return true
}
