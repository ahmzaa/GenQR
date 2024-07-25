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

// Styles for different parts of the application
var (
    headStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#585B70")).
        Background(lipgloss.Color("#A6E3A1")).
        PaddingTop(1).
        PaddingBottom(1).
        Width(80).
        Align(lipgloss.Center)

    helpStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#A6ADC8")).
        Background(lipgloss.Color("#45475A")).
        PaddingTop(1).
        PaddingBottom(1).
        Width(80).
        Align(lipgloss.Center)

    callToActStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#00FF00")).
        Blink(true)
)

// GenQRTerm generates the QRCode and outputs to stdout (visual in the terminal)
func GenQRTerm(data string) error {
    qrc, err := qrcode.New(data)
    if err != nil {
        return fmt.Errorf("could not generate QRCode: %v", err)
    }

    w := terminal.New()
    if err := qrc.Save(w); err != nil {
        return fmt.Errorf("could not output QRCode to terminal: %v", err)
    }
    return nil
}

// GenQRFile exports the generated QRCode to an image file
func GenQRFile(data, path string) error {
    qrc, err := qrcode.New(data)
    if err != nil {
        return fmt.Errorf("could not generate QRCode: %v", err)
    }

    w, err := standard.New(path)
    if err != nil {
        return fmt.Errorf("failed to create the writer: %v", err)
    }

    if err = qrc.Save(w); err != nil {
        return fmt.Errorf("could not save image: %v", err)
    }
    return nil
}

// promptUser prompts the user for input and return the input string
func promptUser(promptStyle lipgloss.Style, promptText string) (string, error) {
    scanner := bufio.NewReader(os.Stdin)
    fmt.Println(promptStyle.Render(promptText))
    input, err := scanner.ReadString('\n')
    if err != nil {
        return "", fmt.Errorf("error reading input: %v", err)
    }
    return strings.TrimSpace(input), nil
}

func main() {
    fmt.Println(headStyle.Render("GenQR - Your Terminal QR code Generator"))
    fmt.Println(helpStyle.Render("HELP\n\nQuit: Ctrl+C      Enter: Enter"))

    data, err := promptUser(callToActStyle, "Enter data to encode:")
    if err != nil {
        fmt.Println(err)
        return
    }

    if err := GenQRTerm(data); err != nil {
        fmt.Println(err)
        return
    }

    path, err := promptUser(callToActStyle, "Enter the path & filename to save:")
    if err != nil {
        fmt.Println(err)
        return
    }

    if err := GenQRFile(data, path); err != nil {
        fmt.Println(err)
    }
}
