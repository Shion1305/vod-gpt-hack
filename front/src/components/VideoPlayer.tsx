import React from "react";

const VideoPlayer = ({
  src,
  currentTime,
  onTimeUpdate,
}: {
  src: string;
  currentTime: number;
  onTimeUpdate: (newTime: number) => void;
}) => {
  return (
    <div className="w-full h-full bg-black flex items-center justify-center">
      <video
        className="w-full h-full object-contain"
        controls
        src={src}
        onTimeUpdate={(e) => onTimeUpdate(e.currentTarget.currentTime)}
      >
        お使いのブラウザは動画タグをサポートしていません。
      </video>
    </div>
  );
};

export default VideoPlayer;