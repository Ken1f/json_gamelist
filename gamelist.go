package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
)

func main() {
    dirname := ""
    playlist := ""
    path := ""

    fmt.Printf("neoKEN's JSON Gamelist Creator\nVersion 1.0\n")

    confirm := "n"
    for confirm != "y" {
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Enter ROM folder :")
      path, _ = reader.ReadString('\n')
      path = strings.Replace(path, "\n", "", -1)

      reader = bufio.NewReader(os.Stdin)
      fmt.Print("Enter PLAYLIST name :")
      playlist, _ = reader.ReadString('\n')
      playlist = strings.Replace(playlist, "\n", "", -1)

    fmt.Printf(`
        {
          "path": "/mnt/sdcard/%s/Cyber Troopers Virtual-On Oratorio Tangram.chd",
          "label": "Cyber Troopers Virtual-On Oratoram Tangram.chd",
          "core_path": "DETECT",
          "core_name": "DETECT",
          "crc32": "DETECT",
          "db_name": "%s.lpl"
        }
    Is this correct? y/n `, path, playlist)

      reader = bufio.NewReader(os.Stdin)
      confirm, _ = reader.ReadString('\n')
      confirm = strings.Replace(confirm, "\n", "", -1)
   }

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter path to scan (or enter '.' for currect path):")
    dirname, _ = reader.ReadString('\n')
    dirname = strings.Replace(dirname, "\n", "", -1)

    if dirname == "" {
        dirname = "."
    }

    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    filelist, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }

    file, err := os.Create(playlist+".lpl")
    if err != nil {
        return
    }
    defer file.Close()


    header := `{
    "version": "1.4",
    "default_core_path": "",
    "default_core_name": "",
    "label_display_mode": 0,
    "right_thumbnail_mode": 0,
    "left_thumbnail_mode": 0,
    "sort_mode": 0,
    "items": [`
    footer :=`
    ]
}`

    size := len(filelist)
    body := ""

    fmt.Printf(header)
    file.WriteString(header)

  for i, thisFile := range filelist {
    romname := thisFile.Name()
    body = fmt.Sprintf(`
        {
          "path": "/mnt/sdcard/%s/%s",
          "label": "%s",
          "core_path": "DETECT",
          "core_name": "DETECT",
          "crc32": "DETECT",
          "db_name": "%s.lpl"`, path, romname, romname, playlist)
    if i < size-1 {
        body = body + `
        },`
    } else {
        body = body + `
        }`
    }
    fmt.Printf(body)
    file.WriteString(body)
  }

    fmt.Printf(footer)
    file.WriteString(footer)
}
