# musicbrainz-wrapper

This is a simple package in Go, that when inputted a song name, will output artist, genre and language info, by fetching the MusicBrainz API.  
Keep in mind that multiple artists can have a song with the same name, at that point you will have to step in (See the examples).  

Usage:  
```bash 
CGO_ENABLED=0 go run example/main.go  
```   
Check the /example folder for the example script and its output.  