# codete-webscraping

## How to use
1) `go build` - this will create executable file
2) `chmod +x ./codete-webscraping` - change file rights
3) `./codete-webscraping` - start server on default port 8080 (can be changed in ./config.yaml)
4) Navigate to localhost:8080 in browser

## Assumptions made
Regarding links...

I assumed that _internal_ links are links started with **#**. Other ones are _external_.

Inaccessible are those links which responded with status >= 400. Also I checked only link
in href attribute itself. So if it is some relative link to a file (href = blah.asp),
then it will be marked as inaccessible.
