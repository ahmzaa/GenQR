package main

import (
    "fmt"

    "github.com/yeqown/go-qrcode/v2"
    "github.com/yeqown/go-qrcode/writer/standard"
    "github.com/yeqown/go-qrcode/writer/terminal"

    "github.com/charmbracelet/lipgloss"
)

var head = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingTop(1).
    PaddingRight(1).
    PaddingBottom(1).
    PaddingLeft(1).
    Width(80).
    Align(lipgloss.Center)

var callToAct = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#00FF00")).
    Blink(true)


func GenQRTerm(data string) {
    qrc, _ := qrcode.New(data)

    w := terminal.New()

    if err := qrc.Save(w); err != nil {
        panic(err)
    }
}

func GenQRFile(data string, path string) {
    qrc, err := qrcode.New(data)
    if err != nil {
        fmt.Printf("could not generate QRCode: %v", err)
        return
    }

    w, err := standard.New(path)
    if err != nil {
        fmt.Printf("standard.New failed: %v", err)
        return
    }

    // save file
    if err = qrc.Save(w); err != nil {
        fmt.Printf("could not save image: %v", err)
    }
}

func main() {
    fmt.Println(head.Render("GenQR - Your Terminal QR code Generator"))
    fmt.Println(callToAct.Render("Enter the data you would like to encode: "))
    var data string
    fmt.Scanln(&data)
    GenQRTerm(data)
    fmt.Println(callToAct.Render("Enter the path & filename you would like to save to: "))
    var path string
    fmt.Scanln(&path)
    GenQRFile(data, path)
}
