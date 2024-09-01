"use client";

import React, { useEffect, useRef, useState } from "react";
import { PlayIcon, PauseIcon } from "@heroicons/react/24/solid";

// VideoPlayerコンポーネントの型定義
const VideoPlayer = ({
  src,
  currentTime,
  onTimeUpdate,
  onDurationChange,
}: {
  src: string;                                    // ビデオのソースURL
  currentTime: number;                            // 現在の再生時間
  onTimeUpdate: (newTime: number) => void;        // 再生時間が更新されたときのコールバック関数
  onDurationChange: (duration: number) => void;   // ビデオの長さが変更されたときのコールバック関数
}) => {
  // ビデオ要素への参照
  const videoRef = useRef<HTMLVideoElement>(null);
  // 再生状態を管理するstate
  const [isPlaying, setIsPlaying] = useState(false);
  // 音量を管理するstate
  const [volume, setVolume] = useState(1);
  // 現在の再生時間を管理するstate
  const [displayTime, setDisplayTime] = useState(0);

  // currentTimeが変更されたときに、ビデオの再生位置を更新
  useEffect(() => {
    if (videoRef.current && Math.abs(videoRef.current.currentTime - currentTime) > 0.1) {
      videoRef.current.currentTime = currentTime;
    }
  }, [currentTime]);

  // srcが変更されたときに、ビデオのソースを更新
  useEffect(() => {
    if (videoRef.current) {
      videoRef.current.src = src;
    }
  }, [src]);

  // 再生/一時停止を切り替える関数
  const togglePlay = () => {
    if (videoRef.current) {
      if (isPlaying) {
        videoRef.current.pause();
      } else {
        videoRef.current.play();
      }
      setIsPlaying(!isPlaying);
    }
  };

  // 音量を変更する関数
  const handleVolumeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newVolume = parseFloat(e.target.value);
    setVolume(newVolume);
    if (videoRef.current) {
      videoRef.current.volume = newVolume;
    }
  };

  // 現在の再生時間をフォーマットする関数
  const formatTime = (time: number) => {
    const minutes = Math.floor(time / 60);
    const seconds = Math.floor(time % 60);
    return `${minutes}:${seconds < 10 ? '0' : ''}${seconds}`;
  };

  return (
    <div className="w-full bg-black rounded-lg overflow-hidden shadow-xl">
      <video
        ref={videoRef}
        className="w-full"
        onTimeUpdate={(e) => {
          onTimeUpdate(e.currentTarget.currentTime);
          setDisplayTime(e.currentTarget.currentTime);
        }}
        onLoadedMetadata={(e) => onDurationChange(e.currentTarget.duration)}
      >
        お使いのブラウザは動画タグをサポートしていません。
      </video>
      <div className="bg-gray-800 p-4">
        <div className="flex items-center justify-between">
          {/* 再生/一時停止ボタン */}
          <button
            onClick={togglePlay}
            className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full transition duration-200"
          >
            {isPlaying ? (
              <PauseIcon className="h-6 w-6" />
            ) : (
              <PlayIcon className="h-6 w-6" />
            )}
          </button>
          {/* 現在の再生時間表示 */}
          <div className="text-white">
            {formatTime(displayTime)} / {videoRef.current ? formatTime(videoRef.current.duration) : '0:00'}
          </div>
          {/* 音量コントロール */}
          <div className="flex items-center">
            <input
              type="range"
              min="0"
              max="1"
              step="0.1"
              value={volume}
              onChange={handleVolumeChange}
              className="w-24 mr-2"
            />
            <span className="text-white">{Math.round(volume * 100)}%</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default VideoPlayer;