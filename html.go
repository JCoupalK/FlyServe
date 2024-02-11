package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func indexExists(path string) bool {
	indexFilePath := filepath.Join(path, "index.html")
	_, err := os.Stat(indexFilePath)
	return !os.IsNotExist(err)
}

func serveDirectoryWithCustomText(w http.ResponseWriter, requestedPath, dirPath string) {
	displayPath := requestedPath
	if displayPath == "." {
		displayPath = "/" // Ensure root is displayed as "/"
	}

	if !strings.HasSuffix(displayPath, "/") {
		displayPath += "/"
	}

	files, err := os.ReadDir(dirPath)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Set Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the HTML header with embedded styles
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <title>Index of %s</title>
	<link href="data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAAB8AAAAfCAYAAAAfrhY5AAAAAXNSR0IArs4c6QAAAQ1JREFUWEfl
l80NwjAMhRsxQcdgFA4swIAs0AOjMAYTIJArOXKCf5OKVkpPQJx8zyavcdIUfE7z+SNNeb+eKbKc
K5gCr/dHXn+5Xab6Ow56hKhwhCKAg0liIBYeTYQIB7CWJSxci+F+gxhJAAuvwV6QFscJ+IF7wRIo
IqCAc+AoxIqnFcjwKFiDeAUU8IhHe2Ix+xUuZd2anWceCDgunPMxLXfrOHo/tZbcKq01DgIGh/fY
pmfu4GWnR+cWVsI1NCuuu117w7X62IJnn4/9esXse2wTmVucaq3/u7WhuI1Me7pjdDJYtt16uKgA
T+uMp5ure5UEREB1vx/q26kA+Pz3Gwu1zi53Nc67W95Sv5g9lNOdT6o/AAAAAElFTkSuQmCC" rel="icon" type="image/x-icon">
    <style>
        body {
			font-family: -apple-system,BlinkMacSystemFont,Segoe UI,Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji";
            margin: 0;
            padding: 0;
            background: #1c1e21;
            color: #333;
        }
        .container {
            margin: 20px auto;
            width: 90%%;
            max-width: 800px;
            background: #202224;
            padding: 20px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
			border-radius: 25px;
        }
        h1 {
            color: #50b7e0;
        }
        a {
            color: #50b7e0;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
        ul {
            padding-right: 20px;
        }
        li {
            padding: 8px 0;
        }
		img {
			display: block;
			margin-left: auto;
			margin-right: auto;
			margin-top: 35px;
			image-rendering: pixelated;
		}
    </style>
	
</head>
<body>
<div class="container">
    <h1>Index of %s</h1>
    <ul>`, displayPath, displayPath)

	// Conditionally add a link to the parent directory if not at root
	if displayPath != "/" {
		// Create a parent directory path
		parentPath := filepath.Dir(filepath.Clean(displayPath))
		if parentPath == "." {
			parentPath = "/"
		}
		fmt.Fprintf(w, `<li><a href="%s">../</a></li>`, parentPath)
	}

	// Generate file links
	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			name += "/"
		}
		href := filepath.Join(displayPath, name) // Use web path for href
		fmt.Fprintf(w, `<li><a href="%s">%s</a></li>`, href, name)
	}

	// Close the HTML tags
	fmt.Fprintf(w, `</ul>
</div>
<a href="https://github.com/JCoupalK/FlyServe" target="_blank">
<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJsAAAAxCAYAAADN5tsEAAAAAXNSR0IArs4c6QAAAz1JREFUeF7t
nM1RAzEMhZOhAsqgFA40QIE0wIFSKIMKGJhNxoxjLD9JdhTH+3Ij/pH99K1ke2OOB36oQJACxyA7
NEMFDoSNEIQpQNjCpKYhwkYGwhQgbGFS0xBhIwNhChC2MKlpiLCRgTAFzLA9PD79SKP7/vo091fr
K9kY1V+YmjTUVEAFRw7Yy9vHX4fvr8+H8u9U2AMKYVuT2iZsyekJqBpcEnxb3e3jgY6w7Qy2zeGt
KLbJUcJX+26rYwWOsO0IthI0LVitehbgCNtOYNOCJoE1AjjCtgPYaqBZoUL1NRGOsC0OmxW0FlS9
wBG2HcAWNUUU3bSwtc788rkke2V9NI7UB2pnLUfnlNp5lcdMo9qVukjzs9o7HX1IUc0bvTTtWo4m
bPLBeQ1U6WFCwUP7EO4aNusTpRVVco41co12vhcab7ubRLbaOVo+AW85OntDkY2wXWI0Gu6rweZN
oShVovIWcFbYRq29rh3Z7nWc2jSK5nckbCjZnNe0tQ0H2kCgdl64ve280HjbleNcAjaEy+g1m9de
dDtvOvS2Q/M7wYYqXatcCrvWNIrGR9jOCnkjlLYd8gMjG1IoK0cPJooIyNTohb5kD41TW470YBrN
FEAL2lIsJC5yEmErfko04mgj9dE6Grmn3SiCzgpt1MYiGm6kQ/MNgvccDcE2+pwNTRI5FzmFsF0q
4F0D3+XrqlFHCggytIFJ7bX1tPZQui77Qc6X+rOmfWv9f2u2tEvJf5WrTaUoTbb67Hk3SthskWYq
2GrbYvQU9pSjtGc9+kD9oTSo3bVZ07E2QvVGRhTZeseN9NXO8+/Ci/QmoSd61SIbWq/l4GudgMQg
bGcFtBlBW0+CWHp4L25XWYGTQGx9rwEDRbaeqMq2t1Pg31W+GnAeqMqopolo5ZOiAfN20tGyVYHq
vVEtcJqrfNuALKBp0qh1kqw/hwLiJeUSOAtY5X1Ta4RiGp0DjtGj4I340YqyP1EB/q8PwhGmgAq2
fDStMxVrupRmyTQa5v9QQ2bYQkdHY0spQNiWcufckyFsc/tnqdERtqXcOfdkCNvc/llqdIRtKXfO
PRnCNrd/lhodYVvKnXNPhrDN7Z+lRvcLGFhjPks+TvsAAAAASUVORK5CYII=" alt="FlyServe Logo" width="155" height="49">
</a>
</body>
</html>`)
}
