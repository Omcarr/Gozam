export function LoadingSpinner() {
  return (
    <div className="flex items-center justify-center">
      <div className="relative">
        {/* Outer ring */}
        <div className="w-16 h-16 border-4 border-cyan-200 rounded-full animate-spin border-t-cyan-500"></div>

        {/* Inner ring */}
        <div className="absolute top-2 left-2 w-12 h-12 border-4 border-blue-200 rounded-full animate-spin border-t-blue-500 animate-reverse"></div>

        {/* Center dot */}
        <div className="absolute top-6 left-6 w-4 h-4 bg-gradient-to-r from-cyan-500 to-blue-500 rounded-full animate-pulse"></div>
      </div>
    </div>
  )
}
