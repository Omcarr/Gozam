"use client"

import { useState, useRef } from "react"
import { Mic, Square, Play, Pause } from "lucide-react"
import { Button } from "@/components/ui/button"

interface AudioRecorderProps {
  onAudioReady: (file: File) => void
  disabled?: boolean
}

export function AudioRecorder({ onAudioReady, disabled }: AudioRecorderProps) {
  const [isRecording, setIsRecording] = useState(false)
  const [recordedAudio, setRecordedAudio] = useState<string | null>(null)
  const [isPlaying, setIsPlaying] = useState(false)
  const [duration, setDuration] = useState(0)

  const mediaRecorderRef = useRef<MediaRecorder | null>(null)
  const audioRef = useRef<HTMLAudioElement | null>(null)
  const chunksRef = useRef<Blob[]>([])
  const intervalRef = useRef<NodeJS.Timeout | null>(null)

  const startRecording = async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
      const mediaRecorder = new MediaRecorder(stream)
      mediaRecorderRef.current = mediaRecorder
      chunksRef.current = []

      mediaRecorder.ondataavailable = (event) => {
        if (event.data.size > 0) {
          chunksRef.current.push(event.data)
        }
      }

      mediaRecorder.onstop = () => {
        const blob = new Blob(chunksRef.current, { type: "audio/wav" })
        const audioUrl = URL.createObjectURL(blob)
        setRecordedAudio(audioUrl)

        // Stop all tracks to release microphone
        stream.getTracks().forEach((track) => track.stop())
      }

      mediaRecorder.start()
      setIsRecording(true)
      setDuration(0)

      // Start duration counter
      intervalRef.current = setInterval(() => {
        setDuration((prev) => prev + 1)
      }, 1000)
    } catch (error) {
      console.error("Error accessing microphone:", error)
    }
  }

  const stopRecording = () => {
    if (mediaRecorderRef.current && isRecording) {
      mediaRecorderRef.current.stop()
      setIsRecording(false)

      if (intervalRef.current) {
        clearInterval(intervalRef.current)
      }
    }
  }

  const togglePlayback = () => {
    if (audioRef.current) {
      if (isPlaying) {
        audioRef.current.pause()
      } else {
        audioRef.current.play()
      }
      setIsPlaying(!isPlaying)
    }
  }

  const handleSubmit = () => {
    if (recordedAudio && chunksRef.current.length > 0) {
      const blob = new Blob(chunksRef.current, { type: "audio/wav" })
      const file = new File([blob], "recording.wav", { type: "audio/wav" })
      onAudioReady(file)
    }
  }

  const formatDuration = (seconds: number) => {
    const mins = Math.floor(seconds / 60)
    const secs = seconds % 60
    return `${mins}:${secs.toString().padStart(2, "0")}`
  }

  return (
    <div className="text-center">
      <h3 className="text-xl font-semibold text-white mb-6">Record Audio</h3>

      <div className="space-y-6">
        {/* Recording Button */}
        <div className="relative">
          <Button
            onClick={isRecording ? stopRecording : startRecording}
            disabled={disabled}
            className={`w-24 h-24 rounded-full text-white transition-all duration-300 ${
              isRecording ? "bg-red-500 hover:bg-red-600 animate-pulse" : "bg-cyan-500 hover:bg-cyan-600"
            }`}
          >
            {isRecording ? <Square className="w-8 h-8" /> : <Mic className="w-8 h-8" />}
          </Button>

          {isRecording && <div className="absolute -inset-2 border-2 border-red-400 rounded-full animate-ping" />}
        </div>

        {/* Duration Display */}
        {(isRecording || recordedAudio) && (
          <div className="text-white text-lg font-mono">{formatDuration(duration)}</div>
        )}

        {/* Playback Controls */}
        {recordedAudio && (
          <div className="space-y-4">
            <audio ref={audioRef} src={recordedAudio} onEnded={() => setIsPlaying(false)} className="hidden" />

            <Button
              onClick={togglePlayback}
              variant="outline"
              className="bg-white/10 border-white/20 text-white hover:bg-white/20"
            >
              {isPlaying ? <Pause className="w-4 h-4 mr-2" /> : <Play className="w-4 h-4 mr-2" />}
              {isPlaying ? "Pause" : "Play"} Recording
            </Button>

            <Button
              onClick={handleSubmit}
              disabled={disabled}
              className="w-full bg-gradient-to-r from-cyan-500 to-blue-500 hover:from-cyan-600 hover:to-blue-600 text-white"
            >
              Find Matches
            </Button>
          </div>
        )}

        {!recordedAudio && !isRecording && (
          <p className="text-gray-300 text-sm">Click the microphone to start recording</p>
        )}
      </div>
    </div>
  )
}
