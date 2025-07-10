# Gozam

An audio recognition webapp built with Go, Gin, PostgreSQL, Redis, and FFmpeg that identifies songs using advanced digital signal processing techniques.

## Demo

<div align="center">
  <a href="https://www.youtube.com/watch?v=jl4M5GFASus">
    <img src="https://img.youtube.com/vi/jl4M5GFASus/maxresdefault.jpg" alt="Gozam Demo" width="600">
  </a>
</div>

## Features

- **Audio Fingerprinting Algorithm**: Converts spectral peaks into compact 32-bit hashes using frequency pairs and time deltas
- **Song Identification**: 100ms timing tolerance for accurate recognition
- **Hybrid Storage System**: Redis for real-time fingerprint lookups and PostgreSQL for song metadata
- **Performance Optimized**: Downsampling and frequency range restriction (20Hz–5kHz) for enhanced signal-to-noise ratio (SNR)
- **YouTube Integration**: Fetch song information and metadata from YouTube API

## Tech Stack

- **Backend**: Go with Gin framework
- **Databases**: PostgreSQL (metadata), Redis (fingerprint storage)
- **Audio Processing**: FFmpeg for handling audio files
- **APIs**: YouTube Data API v3

## Installation

### Prerequisites

- Go 1.19 or higher
- PostgreSQL
- Redis
- FFmpeg
- ytdlp
  
### Setup

1. Clone the repository
   ```bash
   git clone https://github.com/omcarr/Gozam.git
   cd gozam
   ```

2. Install Go dependencies
   ```bash
   cd server
   go mod download
   ```

3. Install client dependencies
   ```bash
   cd client
   npm install
   ```

4. Set up your environment variables by creating a `.env` file:
   ```env
   redisURL=""
   ytApiKey=""
   postgresURL=""
   ytBaseURL=""
   ```
   
## Usage

### Running the Application

1. Start the server:
   ```bash
   cd server
   go run main.go
   ```

2. Start the client (in a separate terminal):
   ```bash
   cd client
   npm run dev
   ```

3. Open your browser and navigate to the client URL (`http://localhost:3000`)

### How It Works

1. **Audio Processing**: Upload or record audio samples
2. **Fingerprinting**: The system extracts spectral peaks and converts them to 32-bit hashes
3. **Matching**: Compares fingerprints against the database using frequency pairs and time deltas
4. **Identification**: Returns matching songs with metadata from PostgreSQL

## Algorithm Overview

The audio fingerprinting algorithm is based on research in digital signal processing:

- **Spectral Peak Extraction**: Identifies prominent frequency peaks in audio spectrograms
- **Hash Generation**: Creates compact 32-bit fingerprints using frequency pairs and time offsets
- **Time Delta Matching**: Provides robust identification with 100ms timing tolerance
- **Frequency Range Optimization**: Focuses on 20Hz–5kHz range for optimal signal-to-noise ratio

### Research Papers

For a deeper understanding of the algorithm:
- [Prerequisite concepts](https://drive.google.com/file/d/1ahyCTXBAZiuni6RTzHzLoOwwfTRFaU-C/view)
- [Audio Fingerprinting Research Paper](https://hajim.rochester.edu/ece/sites/zduan/teaching/ece472/projects/2019/AudioFingerprinting.pdf)

## API Reference
### POST `/register_song`
**Description:** Register a new song or playlist in the database with its audio fingerprints. (only support for yt links for now. Will extend to spotify in future)

**Request:**
```json
{
    "url": "https://www.youtube.com/playlist?list=PLWmBFjd0O0lKObLmy6vdUBL5FVKxZCUIP"
}
```

**Response (Success):**
```json
{
    "message": "Song successfully registered and downloaded"
}
```

**Response (Error):**
```json
{
  "status": "error",
  "message": "Failed to process audio file",
  "error": "detailed_error_message"
}
```

### POST `/find_matches`
**Description:** Find matching songs from an audio sample

**Request:**
```json
{
  "audio_file": "base64_encoded_audio_data"
}
```

**Response (Match Found):**
```json
{
  "status": "success",
  "matches": [
    "matches": [
        {
            "SongID": 2274099950,
            "SongTitle": "Teri Ore",
            "SongArtist": "Pritam - Topic",
            "YouTubeID": "fa4l-vJ_5ic",
            "Timestamp": 23,
            "Score": 32004000
        },
]
```

**Response (Error):**
```json
{
  "status": "error",
  "message": "Failed to process audio",
  "error": "Audio file format not supported"
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
