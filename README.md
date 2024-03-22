## Song Download Library Documentation

### Introduction

The Song Download Library provides functionalities to search for songs, pre-download them from YouTube, and download them using the APIDL service.

### Installation

To use the library in your Go project, you can install it using `go get`:

```bash
go get -u github.com/refaldyrk/song_download
```

### Usage

1. **Search for a Song**

   Use the `GetSong(keyword string) Song` function to search for a song based on the provided keyword.

    ```go
    song := song_download.GetSong("keyword")
    ```

2. **Pre-download a Song**

   Call the `PreDownload()` method on an `Item` to pre-download a song from YouTube.

    ```go
    extractor := item.PreDownload()
    ```

3. **Download a Song**

   Initialize a `Task` with the necessary parameters and call the `Download(jwt string) string` method to download the pre-downloaded song.

    ```go
   extractor.Download()
    ```

### Functions and Methods

- **GetSong(keyword string) Song**: Searches for a song based on the provided keyword and returns information about the song.

- **PreDownload() Extractor**: Pre-downloads a song from YouTube using its ID and returns an extractor for further processing.

- **Download(jwt string) string**: Downloads the pre-downloaded song using the provided JWT token and returns the download URL.

### Dependencies

This library relies on the following dependencies:

- **HTTP Client**: To make HTTP requests for searching, pre-downloading, and downloading songs.
- **JSON**: To encode and decode JSON data for communication with the APIDL service.

### Error Handling
The library uses panic for error handling. It's recommended to handle errors appropriately in your application code.