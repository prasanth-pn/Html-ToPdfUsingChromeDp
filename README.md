Invoice PDF Generation
Previously, I used the wkhtmltopdf package to generate invoice PDFs. However, it was not user-friendly and consumed a high amount of CPU resources.

In my current project, I generate PDFs from HTML files and directly return them as a binary stream response, eliminating the need to store them on the server.

Additionally, I have used this approach for various use cases, such as converting HTML-designed resumes into PDFs.

How It Works
chromedp operates using the Chrome engine.
Google Chrome must be installed for it to function properly.
The process involves rendering the HTML file in the Chrome engine and converting it into a PDF.
