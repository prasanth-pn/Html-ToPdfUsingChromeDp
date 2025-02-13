# PDF Generation Using Go and Chromedp

## Overview

This project provides a solution for generating PDFs from HTML files using `chromedp`, which leverages the Chrome engine for rendering. Unlike `wkhtmltopdf`, which is resource-intensive, this approach is more efficient and allows direct streaming of the PDF response without saving it on the server.

## Use Cases

- Generating invoice PDFs.
- Converting HTML-designed resumes into PDFs.
- Any other scenario where HTML needs to be rendered and exported as a PDF.

## How It Works

1. `chromedp` renders the given HTML file using the Chrome engine.
2. The rendered page is then converted into a PDF.
3. The PDF is streamed as a binary response, avoiding unnecessary file storage on the server.

## Prerequisites

- **Google Chrome** must be installed on the system.
- Go modules should be properly configured.

## Installation

1. Install Google Chrome if not already installed.
2. Install the required Go dependencies:

   ```sh
   go get -u github.com/chromedp/chromedp
