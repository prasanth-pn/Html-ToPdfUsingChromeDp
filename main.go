package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// Function to serve the HTML content
func writeHTML(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, strings.TrimSpace(content))
	})
}

func main() {
	// Start a temporary server to serve the HTML content
	ts := httptest.NewServer(writeHTML(htmlContent))
	defer ts.Close()

	// Create a new ChromeDP context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Variable to store the generated PDF
	var pdfBuffer []byte

	// Run ChromeDP tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),                      // Navigate to the temporary server's URL
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // Wait for the page to load
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuffer, _, err = page.PrintToPDF().
				WithPaperWidth(8.27).  // A4 width in inches
				WithPaperHeight(11.69). // A4 height in inches
				WithMarginTop(0.4).   // Adjust margins as needed
				WithMarginBottom(0.4).
				WithMarginLeft(0.4).
				WithMarginRight(0.4).
				WithPrintBackground(true). // Ensure background colors are printed
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		log.Fatal("Failed to generate PDF:", err)
	}

	// Save the PDF to a file
	err = os.WriteFile("invoice.pdf", pdfBuffer, 0644)
	if err != nil {
		log.Fatal("Failed to save PDF:", err)
	}
	log.Println("PDF generated successfully: invoice.pdf")
}

var htmlContent = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Invoice</title>
<style>
	body {
		font-family: Arial, sans-serif;
		margin: 0;
		padding: 20px;
		color: #333;
	}
	.invoice {
		max-width: 800px;
		margin: 0 auto;
		padding: 20px;
		border: 1px solid #ddd;
		border-radius: 5px;
		background-color: #f9f9f9;
	}
	.header {
		text-align: center;
		margin-bottom: 20px;
	}
	.header h1 {
		margin: 0;
		font-size: 24px;
		color: #333;
	}
	.company-info, .customer-info {
		margin-bottom: 20px;
	}
	.company-info h2, .customer-info h2 {
		margin: 0 0 10px 0;
		font-size: 18px;
		color: #555;
	}
	.invoice-details {
		margin-bottom: 20px;
	}
	.invoice-details table {
		width: 100%;
		border-collapse: collapse;
	}
	.invoice-details th, .invoice-details td {
		padding: 10px;
		border: 1px solid #ddd;
		text-align: left;
	}
	.invoice-details th {
		background-color: #f1f1f1;
	}
	.total {
		text-align: right;
		font-size: 18px;
		font-weight: bold;
		margin-top: 20px;
	}
	.footer {
		text-align: center;
		margin-top: 30px;
		font-size: 14px;
		color: #777;
	}
</style>
</head>
<body>
<div class="invoice">
	<!-- Header -->
	<div class="header">
		<h1>Invoice</h1>
	</div>

	<!-- Company Information -->
	<div class="company-info">
		<h2>ABC Corp</h2>
		<p>123 Main St, City, Country</p>
		<p>Email: info@abccorp.com | Phone: +123 456 7890</p>
	</div>

	<!-- Customer Information -->
	<div class="customer-info">
		<h2>Bill To:</h2>
		<p>John Doe</p>
		<p>456 Elm St, City, Country</p>
		<p>Email: john.doe@example.com</p>
	</div>

	<!-- Invoice Details -->
	<div class="invoice-details">
		<table>
			<thead>
				<tr>
					<th>Item</th>
					<th>Quantity</th>
					<th>Unit Price</th>
					<th>Total</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>Product A</td>
					<td>2</td>
					<td>$50.00</td>
					<td>$100.00</td>
				</tr>
				<tr>
					<td>Product B</td>
					<td>1</td>
					<td>$75.00</td>
					<td>$75.00</td>
				</tr>
				<tr>
					<td>Product C</td>
					<td>3</td>
					<td>$20.00</td>
					<td>$60.00</td>
				</tr>
			</tbody>
		</table>
	</div>

	<!-- Total -->
	<div class="total">
		<p>Total: $235.00</p>
	</div>

	<!-- Footer -->
	<div class="footer">
		<p>Thank you for your business!</p>
	</div>
</div>
</body>
</html>
`
