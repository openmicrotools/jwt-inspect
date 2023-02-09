# README

to build this project, run ```go build``` from within the cli folder
to use this in command line, run ```cp testcli /usr/local/bin```

This is currently an internal POC, and is completely stubbed.

To run the POC after adding to bin, run ```jwt-inspect <jwt>```. You can optionally use a boolean flag --epoch ```jwt-inspect <jwt> --epoch=true```. This defaults to false if not provided. If only the flag is provided, true is assumed.