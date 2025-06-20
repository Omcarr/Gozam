"use client"

import type React from "react"

import { useRef, useState } from "react"
import { Upload, Music } from "lucide-react"
import { Button } from "@/components/ui/button"

interface FileUploadProps {
  onFileSelect: (file: File) => void
  disabled?: boolean
}

export function FileUpload({ onFileSelect, disabled }: FileUploadProps) {
  const [selectedFile, setSelectedFile] = useState<File | null>(null)
  const [dragActive, setDragActive] = useState(false)
  const fileInputRef = useRef<HTMLInputElement>(null)

  const handleFileSelect = (file: File) => {
    if (file && file.type.startsWith("audio/")) {
      setSelectedFile(file)
    }
  }

  const handleDrag = (e: React.DragEvent) => {
    e.preventDefault()
    e.stopPropagation()
    if (e.type === "dragenter" || e.type === "dragover") {
      setDragActive(true)
    } else if (e.type === "dragleave") {
      setDragActive(false)
    }
  }

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault()
    e.stopPropagation()
    setDragActive(false)

    if (e.dataTransfer.files && e.dataTransfer.files[0]) {
      handleFileSelect(e.dataTransfer.files[0])
    }
  }

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      handleFileSelect(e.target.files[0])
    }
  }

  const handleSubmit = () => {
    if (selectedFile) {
      onFileSelect(selectedFile)
    }
  }

  const formatFileSize = (bytes: number) => {
    if (bytes === 0) return "0 Bytes"
    const k = 1024
    const sizes = ["Bytes", "KB", "MB", "GB"]
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    return Number.parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i]
  }

  return (
    <div className="text-center">
      <h3 className="text-xl font-semibold text-white mb-6">Upload Audio File</h3>

      <div className="space-y-6">
        {/* Drop Zone */}
        <div
          className={`border-2 border-dashed rounded-lg p-8 transition-all duration-300 cursor-pointer ${
            dragActive ? "border-cyan-400 bg-cyan-400/10" : "border-gray-400 hover:border-cyan-400 hover:bg-cyan-400/5"
          }`}
          onDragEnter={handleDrag}
          onDragLeave={handleDrag}
          onDragOver={handleDrag}
          onDrop={handleDrop}
          onClick={() => fileInputRef.current?.click()}
        >
          <input
            ref={fileInputRef}
            type="file"
            accept="audio/*"
            onChange={handleInputChange}
            className="hidden"
            disabled={disabled}
          />

          <div className="space-y-4">
            <Upload className="w-12 h-12 text-gray-400 mx-auto" />
            <div className="text-white">
              <p className="text-lg font-medium">Drop your audio file here</p>
              <p className="text-gray-300 text-sm">or click to browse</p>
            </div>
            <p className="text-gray-400 text-xs">Supports MP3, WAV, M4A and other audio formats</p>
          </div>
        </div>

        {/* Selected File Info */}
        {selectedFile && (
          <div className="bg-white/10 rounded-lg p-4 space-y-4">
            <div className="flex items-center justify-center gap-3">
              <Music className="w-5 h-5 text-cyan-400" />
              <div className="text-white text-left">
                <p className="font-medium truncate max-w-48">{selectedFile.name}</p>
                <p className="text-gray-300 text-sm">{formatFileSize(selectedFile.size)}</p>
              </div>
            </div>

            <Button
              onClick={handleSubmit}
              disabled={disabled}
              className="w-full bg-gradient-to-r from-cyan-500 to-blue-500 hover:from-cyan-600 hover:to-blue-600 text-white"
            >
              Find Matches
            </Button>
          </div>
        )}
      </div>
    </div>
  )
}
