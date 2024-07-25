# GenQR

## Overview
GenQR is a QR code generator program written in Golang. This repository contains
the source code for the GenQR program, allowing users to generate QR codes
locally with ease.

## Features
- Generates QR codes locally.
- Output QR codes to terminal or as an image file.
- Encode any data into QR codes.

## Motivation
As a frequent traveler, I enjoy leaving stickers that indicate my presence at
various locations. While this may seem trivial to some, I find it adds a fun
touch to my travels.

I discovered a small, compact thermal label printer capable of printing on
waterproof paper. This allows me to print my own stickers on demand, without the
need for expensive custom sticker services.

## Benefits of Generating QR Codes locally

Many online QR code generators require account creation, restrict features, or
add tracking elements. For my basic use case, paying for a QR code service is
not sustainable, and there's uncertainty about the codes' functionality once the
subscription ends.

Generating QR codes locally ensures:
- No tracking or unwanted data collection.
- Full customization options.
- Independence from subscription services.
- Simple, reproducible encoded data, as intended for QR codes.

## Drawbacks
Local generation lacks the advanced features provided by some online services.
However, for those who value these additional features, creating a custom web
solution is always an option.

## Implementation
- Programming Language: Golang
- Libraries:
    - github.com/yeqown/go-qrcode/v2
    - github.com/charmbracelet/lipgloss

## To-Do List
- [x] Basic functionality
- [x] Make it pretty
- [ ] Implement more features provided by yeqown/go-qrcode

## Notes
GenQR is essentially a wrapper around yeqown/go-qrcode, with most of the heavy
lifting done by the go-qrcode library. This project serves as a convenient
interface for generating QR codes using Golang.
