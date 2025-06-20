// import React from "react"
// import { ExternalLink, Clock, Star } from "lucide-react"
// import type { Match } from "../app"

// interface MatchResultsProps {
//   matches: Match[]
// }

// export function MatchResults({ matches }: MatchResultsProps) {
//   const formatTimestamp = (seconds: number) => {
//     const mins = Math.floor(seconds / 60)
//     const secs = seconds % 60
//     return `${mins}:${secs.toString().padStart(2, "0")}`
//   }

//   const getScoreColor = (score: number) => {
//     if (score >= 0.8) return "text-green-400"
//     if (score >= 0.6) return "text-yellow-400"
//     return "text-red-400"
//   }

//   const getScoreLabel = (score: number) => {
//     if (score >= 0.9) return "Excellent"
//     if (score >= 0.8) return "Very Good"
//     if (score >= 0.6) return "Good"
//     if (score >= 0.4) return "Fair"
//     return "Poor"
//   }

//   return (
//     <div className="space-y-6 mt-12">
//       {/* Heading */}
//       <div className="text-center">
//         <h2 className="text-3xl font-bold text-white mb-2">
//           {matches.length} Match{matches.length !== 1 ? "es" : ""} Found
//         </h2>
//         <p className="text-gray-400 text-sm">Top YouTube matches based on your audio</p>
//       </div>

//       {/* Carousel */}
//       <div className="relative">
//         <style jsx global>{`
//           .scrollbar-hide {
//             scrollbar-width: none;
//             -ms-overflow-style: none;
//           }
//           .scrollbar-hide::-webkit-scrollbar {
//             display: none;
//           }
//         `}</style>
//         <div className="overflow-x-auto scrollbar-hide pb-4">
//           <div className="flex gap-6 px-4 min-w-max">
//             {matches.map((match, index) => (
//               <div
//                 key={match.SongID}
//                 className="relative group w-80 flex-shrink-0 rounded-2xl overflow-hidden shadow-2xl border border-white/20 bg-gradient-to-br from-white/10 to-white/5 backdrop-blur-sm hover:shadow-cyan-500/20 transition-all duration-500 hover:scale-105 hover:border-cyan-400/40"
//               >
//                 {/* Thumbnail */}
//                 <div className="relative overflow-hidden">
//                   <img
//                     src={`https://img.youtube.com/vi/${match.YouTubeID}/maxresdefault.jpg`}
//                     alt={`${match.SongTitle} thumbnail`}
//                     className="w-full h-48 object-cover transition-all duration-700 group-hover:scale-110 group-hover:brightness-75"
//                   />
//                   <div className="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
                  
//                   {/* Rank Badge */}
//                   <div className="absolute top-4 left-4 z-10">
//                     <span className="bg-gradient-to-r from-cyan-500 to-blue-500 text-white font-bold px-3 py-1 rounded-full text-sm shadow-lg">
//                       #{index + 1}
//                     </span>
//                   </div>

//                   {/* Score Badge */}
//                   <div className="absolute top-4 right-4 z-10">
//                     <span className={`flex items-center gap-1 px-3 py-1 rounded-full text-sm font-semibold backdrop-blur-sm ${
//                       match.Score >= 0.8 ? 'bg-green-500/80 text-white' : 
//                       match.Score >= 0.6 ? 'bg-yellow-500/80 text-white' : 
//                       'bg-red-500/80 text-white'
//                     }`}>
//                       <Star className="w-3 h-3 fill-current" />
//                       {getScoreLabel(match.Score)}
//                     </span>
//                   </div>
//                 </div>

//                 {/* Content */}
//                 <div className="p-6 space-y-4">
//                   <div className="space-y-2">
//                     <h3 className="font-bold text-xl text-white line-clamp-2 group-hover:text-cyan-400 transition-colors duration-300">
//                       {match.SongTitle}
//                     </h3>
//                     <p className="text-gray-300 text-lg">by {match.SongArtist}</p>
//                   </div>

//                   <div className="flex items-center gap-2 text-gray-400">
//                     <Clock className="w-4 h-4" />
//                     <span className="text-sm">Matched at {formatTimestamp(match.Timestamp)}</span>
//                   </div>

//                   <a
//                     href={`https://www.youtube.com/watch?v=${match.YouTubeID}&t=${match.Timestamp}s`}
//                     target="_blank"
//                     rel="noopener noreferrer"
//                     className="flex items-center justify-center gap-2 w-full bg-gradient-to-r from-red-600 to-red-500 hover:from-red-500 hover:to-red-400 text-white font-semibold py-3 px-4 rounded-xl transition-all duration-300 transform hover:scale-105 hover:shadow-lg hover:shadow-red-500/25"
//                   >
//                     <ExternalLink className="w-4 h-4" />
//                     Watch on YouTube
//                   </a>
//                 </div>
//               </div>
//             ))}
//           </div>
//         </div>
//       </div>
//     </div>
//   )
// }




import React from "react"
import { ExternalLink, Clock, Star } from "lucide-react"
import type { Match } from "../app"

interface MatchResultsProps {
  matches: Match[]
}

export function MatchResults({ matches }: MatchResultsProps) {
  const formatTimestamp = (seconds: number) => {
    const mins = Math.floor(seconds / 60)
    const secs = seconds % 60
    return `${mins}:${secs.toString().padStart(2, "0")}`
  }

  // // Normalize the score to a 0–1 range
  // const maxScore = Math.max(...matches.map(m => m.Score))
  // const normalizeScore = (score: number) => score/ maxScore

  // const getScoreColor = (normalized: number) => {
  //   if (normalized >= 0.8) return "text-green-400"
  //   if (normalized >= 0.6) return "text-yellow-400"
  //   return "text-red-400"
  // }

  // const getScoreLabel = (normalized: number) => {
  //   if (normalized >= 0.9) return "Excellent"
  //   if (normalized >= 0.8) return "Very Good"
  //   if (normalized >= 0.6) return "Good"
  //   if (normalized >= 0.4) return "Fair"
  //   return "Poor"
  // }

  return (
    <div className="space-y-6 mt-12">
      <div className="text-center">
        <h2 className="text-3xl font-bold text-white mb-2">
          {matches.length} Match{matches.length !== 1 ? "es" : ""} Found
        </h2>
        <p className="text-gray-400 text-sm">Top YouTube matches based on your audio</p>
      </div>

      <div className="relative">
        <style jsx global>{`
          .scrollbar-hide {
            scrollbar-width: none;
            -ms-overflow-style: none;
          }
          .scrollbar-hide::-webkit-scrollbar {
            display: none;
          }
        `}</style>
        <div className="overflow-x-auto scrollbar-hide pb-4">
          <div className="flex gap-6 px-4 min-w-max">
            {matches.map((match, index) => {
              // const normalizedScore = normalizeScore(match.Score)

              return (
                <div
                  key={match.SongID}
                  className="relative group w-80 flex-shrink-0 rounded-2xl overflow-hidden shadow-2xl border border-white/20 bg-gradient-to-br from-white/10 to-white/5 backdrop-blur-sm hover:shadow-cyan-500/20 transition-all duration-500 hover:scale-105 hover:border-cyan-400/40"
                >
                  <div className="relative overflow-hidden">
                    <img
                      src={`https://img.youtube.com/vi/${match.YouTubeID}/maxresdefault.jpg`}
                      alt={`${match.SongTitle} thumbnail`}
                      className="w-full h-48 object-cover transition-all duration-700 group-hover:scale-110 group-hover:brightness-75"
                    />
                    <div className="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>

                    {/* Rank Badge */}
                    <div className="absolute top-4 left-4 z-10">
                      <span className="bg-gradient-to-r from-cyan-500 to-blue-500 text-white font-bold px-3 py-1 rounded-full text-sm shadow-lg">
                        #{index + 1}
                      </span>
                    </div>

                    {/* Score Badge */}
                    {/* <div className="absolute top-4 right-4 z-10">
                      <span className={`flex items-center gap-1 px-3 py-1 rounded-full text-sm font-semibold backdrop-blur-sm ${
                        normalizedScore >= 0.8
                          ? "bg-green-500/80 text-white"
                          : normalizedScore >= 0.6
                          ? "bg-yellow-500/80 text-white"
                          : "bg-red-500/80 text-white"
                      }`}>
                        <Star className="w-3 h-3 fill-current" />
                        {getScoreLabel(normalizedScore)}
                      </span>
                    </div> */}
                  </div>

                  {/* Content */}
                  <div className="p-6 space-y-4">
                    <div className="space-y-2">
                      <h3 className="font-bold text-xl text-white line-clamp-2 group-hover:text-cyan-400 transition-colors duration-300">
                        {match.SongTitle}
                      </h3>
                      <p className="text-gray-300 text-lg">by {match.SongArtist}</p>
                    </div>

                    <div className="flex items-center gap-2 text-gray-400">
                      <Clock className="w-4 h-4" />
                      <span className="text-sm">Matched at {formatTimestamp(match.Timestamp)}</span>
                    </div>

                    <a
                      href={`https://www.youtube.com/watch?v=${match.YouTubeID}&t=${match.Timestamp}s`}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center justify-center gap-2 w-full bg-gradient-to-r from-red-600 to-red-500 hover:from-red-500 hover:to-red-400 text-white font-semibold py-3 px-4 rounded-xl transition-all duration-300 transform hover:scale-105 hover:shadow-lg hover:shadow-red-500/25"
                    >
                      <ExternalLink className="w-4 h-4" />
                      Watch on YouTube
                    </a>
                  </div>
                </div>
              )
            })}
          </div>
        </div>
      </div>
    </div>
  )
}
