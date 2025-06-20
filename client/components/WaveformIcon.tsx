interface WaveformIconProps {
  className?: string
}

export function WaveformIcon({ className = "w-6 h-6" }: WaveformIconProps) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <rect
        x="2"
        y="8"
        width="2"
        height="8"
        fill="currentColor"
        className="animate-pulse"
        style={{ animationDelay: "0ms" }}
      />
      <rect
        x="6"
        y="4"
        width="2"
        height="16"
        fill="currentColor"
        className="animate-pulse"
        style={{ animationDelay: "150ms" }}
      />
      <rect
        x="10"
        y="6"
        width="2"
        height="12"
        fill="currentColor"
        className="animate-pulse"
        style={{ animationDelay: "300ms" }}
      />
      <rect
        x="14"
        y="2"
        width="2"
        height="20"
        fill="currentColor"
        className="animate-pulse"
        style={{ animationDelay: "450ms" }}
      />
      <rect
        x="18"
        y="7"
        width="2"
        height="10"
        fill="currentColor"
        className="animate-pulse"
        style={{ animationDelay: "600ms" }}
      />
      <rect
        x="22"
        y="9"
        width="2"
        height="6"
        fill="currentColor"
        className="animate-pulse"
        style={{ animationDelay: "750ms" }}
      />
    </svg>
  )
}
