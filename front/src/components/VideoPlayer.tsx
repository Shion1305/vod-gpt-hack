import React from "react";

const VideoPlayer = (
  props: Readonly<{
    src: string;
    currentTime: number;
    onTimeUpdate: (newTime: number) => void;
  }>,
) => {
  return (
    <div className="bg-gray-300 h-64 mb-4">
      <video className="w-full h-full" controls>
        <source src={props.src} type="video/mp4" />
        お使いのブラウザは動画タグをサポートしていません。
      </video>
    </div>
  );
};

export default VideoPlayer;
