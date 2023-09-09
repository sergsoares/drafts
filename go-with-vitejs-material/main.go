package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<html>
			<head>
				<!-- Include the CSS Material framework -->
				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@material/top-app-bar/dist/mdc.top-app-bar.min.css">
				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@material/button/dist/mdc.button.min.css">
				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@material/textfield/dist/mdc.textfield.min.css">

				<!-- Include ViteJS -->
				<script type="module" src="https://cdn.jsdelivr.net/npm/vite@2.0.0/dist/vite.esm.js"></script>
			</head>
			<body>
				<!-- Use the Material top app bar component -->
				<header class="mdc-top-app-bar">
					<div class="mdc-top-app-bar__row">
						<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
							<span class="mdc-top-app-bar__title">My App</span>

						</section>
					</div>
				</header>
				<main>
					<!-- Use the Material textfield component -->
					<div class="mdc-text-field" data-mdc-auto-init="MDCTextField">
						<input type="text" id="my-text-field" class="mdc-text-field__input">
						<label for="my-text-field" class="mdc-text-field__label">Enter text</label>
					</div>
					<!-- Use the Material button component -->
					<button class="mdc-button" data-mdc-auto-init="MDCRipple">
						Submit
					</button>
				</main>
				<!-- Initialize the Material components -->
				<script>mdc.autoInit();</script>
			</body>
		</html>
		`
		fmt.Fprintf(w, html)
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
