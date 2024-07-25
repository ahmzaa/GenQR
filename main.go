package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/yeqown/go-qrcode/v2"
    "github.com/yeqown/go-qrcode/writer/standard"
    "github.com/yeqown/go-qrcode/writer/terminal"

    "github.com/charmbracelet/lipgloss"
)

// Style used for header
var headStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#585B70")).
    Background(lipgloss.Color("#A6E3A1")).
    PaddingTop(1).
    PaddingBottom(1).
    Width(80).
    Align(lipgloss.Center)

// Style used when user asked for input
var callToActStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#00FF00")).
    Blink(true)

// For the help bar
var helpStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#A6ADC8")).
    Background(lipgloss.Color("#45475A")).
    PaddingTop(1).
    PaddingBottom(1).
    Width(80).
    Align(lipgloss.Center)


// Generates the QRCode and outputs to stdout (Visual in the terminal)
func GenQRTerm(data string) {
    qrc, _ := qrcode.New(data)

    w := terminal.New()

    if err := qrc.Save(w); err != nil {
        panic(err)
    }
}

// Export the generated QRCode to an image file
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
    scanner := bufio.NewReader(os.Stdin)

    fmt.Println(headStyle.Render("GenQR - Your Terminal QR code Generator"))
    fmt.Println(helpStyle.Render("HELP\n\nQuit: Ctrl+C      Enter: Enter"))

    fmt.Println(callToActStyle.Render("Enter the data you would like to encode: "))
    data, err := scanner.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return
    }
    data = strings.TrimSpace(data)
    GenQRTerm(data)

    fmt.Println(callToActStyle.Render("Enter the path & filename you would like to save to: "))
    path, err := scanner.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return
    }
    path = strings.TrimSpace(path)
    GenQRFile(data, path)
}
