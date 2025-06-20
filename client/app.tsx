// "use client"

// import { useState } from "react"
// import { AudioRecorder } from "./components/AudioRecorder"
// import { FileUpload } from "./components/FileUpload"
// import { MatchResults } from "./components/MatchResults"
// import { LoadingSpinner } from "./components/LoadingSpinner"
// import { WaveformIcon } from "./components/WaveformIcon"

// export interface Match {
//   SongID: number
//   SongTitle: string
//   SongArtist: string
//   YouTubeID: string
//   Timestamp: number
//   Score: number
// }

// export default function App() {
//   const [matches, setMatches] = useState<Match[]>([])
//   const [isLoading, setIsLoading] = useState(false)
//   const [error, setError] = useState<string | null>(null)

//   const handleAudioSubmit = async (audioFile: File) => {
//     setIsLoading(true)
//     setError(null)
//     setMatches([])

//     try {
//       const formData = new FormData()
//       formData.append("file", audioFile)

//       const response = await fetch("http://localhost:8000/find_matches", {
//         method: "POST",
//         body: formData,
//       })

//       if (!response.ok) {
//         throw new Error("Failed to find matches")
//       }

//       const data: Match[] = await response.json()
//       setMatches(data)

//       console.log(data)

//     } catch (err) {
//       setError(err instanceof Error ? err.message : "An error occurred")
//     } finally {
//       setIsLoading(false)
//     }
//   }

//   return (
//     <div className="min-h-screen bg-gradient-to-br from-purple-900 via-blue-900 to-indigo-900">
//       <div className="container mx-auto px-4 py-8">
//         {/* Header */}
//         <header className="text-center mb-12">
//           <div className="flex items-center justify-center gap-3 mb-4">
//             <WaveformIcon className="w-8 h-8 text-cyan-400" />
//             <h1 className="text-4xl md:text-6xl font-bold text-white tracking-tight">Gozam</h1>
//             <WaveformIcon className="w-8 h-8 text-cyan-400 scale-x-[-1]" />
//           </div>
//           <p className="text-xl text-gray-300 max-w-2xl mx-auto">
//             Discover any song in seconds. Record audio or upload a file to find your perfect match.
//           </p>
//         </header>

//         {/* Main Content */}
//         <div className="max-w-4xl mx-auto">
//           {/* Audio Input Section */}
//           <div className="bg-white/10 backdrop-blur-lg rounded-2xl p-8 mb-8 border border-white/20">
//             <div className="grid md:grid-cols-2 gap-8">
//               <AudioRecorder onAudioReady={handleAudioSubmit} disabled={isLoading} />
//               <FileUpload onFileSelect={handleAudioSubmit} disabled={isLoading} />
//             </div>
//           </div>

//           {/* Loading State */}
//           {isLoading && (
//             <div className="text-center py-12">
//               <LoadingSpinner />
//               <p className="text-white mt-4 text-lg">Analyzing your audio...</p>
//             </div>
//           )}

//           {/* Error State */}
//           {error && (
//             <div className="bg-red-500/20 border border-red-500/50 rounded-lg p-4 mb-8">
//               <p className="text-red-200 text-center">{error}</p>
//             </div>
//           )}

//           {/* Results */}
//           {matches.length > 0 && <MatchResults matches={matches} />}
//         </div>
//       </div>
//     </div>
//   )
// }



"use client"

import { useState } from "react"
import { AudioRecorder } from "./components/AudioRecorder"
import { FileUpload } from "./components/FileUpload"
import { MatchResults } from "./components/MatchResults"
import { LoadingSpinner } from "./components/LoadingSpinner"
import { WaveformIcon } from "./components/WaveformIcon"

export interface Match {
  SongID: number
  SongTitle: string
  SongArtist: string
  YouTubeID: string
  Timestamp: number
  Score: number
}

export default function App() {
  const [matches, setMatches] = useState<Match[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleAudioSubmit = async (audioFile: File) => {
    setIsLoading(true)
    setError(null)
    setMatches([])

    try {
      const formData = new FormData()
      formData.append("file", audioFile)

      const response = await fetch("http://localhost:8000/find_matches", {
        method: "POST",
        body: formData,
      })

      if (!response.ok) {
        throw new Error("Failed to find matches")
      }

      const data = await response.json()
      console.log(data)

      if (data.matches && Array.isArray(data.matches)) {
        setMatches(data.matches)
      } else {
        throw new Error("Invalid response format")
      }

    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred")
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-900 via-blue-900 to-indigo-900">
      <div className="container mx-auto px-4 py-8">
        {/* Header */}
        <header className="text-center mb-12">
          <div className="flex items-center justify-center gap-3 mb-4">
            <WaveformIcon className="w-8 h-8 text-cyan-400" />
            <h1 className="text-4xl md:text-6xl font-bold text-white tracking-tight">Gozam</h1>
            <WaveformIcon className="w-8 h-8 text-cyan-400 scale-x-[-1]" />
          </div>
          <p className="text-xl text-gray-300 max-w-2xl mx-auto">
            Discover any song in seconds. Record audio or upload a file to find your perfect match.
          </p>
        </header>

        {/* Main Content */}
        <div className="max-w-4xl mx-auto">
          {/* Audio Input Section */}
          <div className="bg-white/10 backdrop-blur-lg rounded-2xl p-8 mb-8 border border-white/20">
            <div className="grid md:grid-cols-2 gap-8">
              <AudioRecorder onAudioReady={handleAudioSubmit} disabled={isLoading} />
              <FileUpload onFileSelect={handleAudioSubmit} disabled={isLoading} />
            </div>
          </div>

          {/* Loading State */}
          {isLoading && (
            <div className="text-center py-12">
              <LoadingSpinner />
              <p className="text-white mt-4 text-lg">Analyzing your audio...</p>
            </div>
          )}

          {/* Error State */}
          {error && (
            <div className="bg-red-500/20 border border-red-500/50 rounded-lg p-4 mb-8">
              <p className="text-red-200 text-center">{error}</p>
            </div>
          )}

          {/* Results */}
          {matches.length > 0 && <MatchResults matches={matches} />}
        </div>
      </div>
    </div>
  )
}
