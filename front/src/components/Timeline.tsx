"use client";

import React, { useState, useRef, useEffect } from "react";

// Timelineコンポーネントの型定義
const Timeline = ({
  currentTime,         // 現在の再生時間（パーセンテージ）
  selectionStart,      // 選択範囲の開始位置（パーセンテージ）
  selectionEnd,        // 選択範囲の終了位置（パーセンテージ）
  onTimeChange,        // 現在時間が変更されたときのコールバック関数
  onSelectionChange,   // 選択範囲が変更されたときのコールバック関数
}: {
  currentTime: number;
  selectionStart: number;
  selectionEnd: number;
  onTimeChange: (newTime: number) => void;
  onSelectionChange: (start: number, end: number) => void;
}) => {
  // ドラッグ中の要素を管理するstate
  const [isDragging, setIsDragging] = useState<
    "current" | "start" | "end" | null
  >(null);
  
  // タイムライン要素への参照
  const timelineRef = useRef<HTMLDivElement>(null);

  // マウスダウンイベントのハンドラ
  const handleMouseDown = (
    e: React.MouseEvent,
    type: "current" | "start" | "end" | null,
  ) => {
    setIsDragging(type);
  };

  // マウス移動イベントのハンドラ
  const handleMouseMove = (e: MouseEvent) => {
    if (!isDragging || !timelineRef.current) return;

    const rect = timelineRef.current.getBoundingClientRect();
    const position = ((e.clientX - rect.left) / rect.width) * 100;

    if (isDragging === "current") {
      // 現在時間のマーカーをドラッグ中
      onTimeChange(Math.max(0, Math.min(100, position)));
    } else if (isDragging === "start") {
      // 選択範囲の開始マーカーをドラッグ中
      onSelectionChange(
        Math.max(0, Math.min(selectionEnd, position)),
        selectionEnd,
      );
    } else if (isDragging === "end") {
      // 選択範囲の終了マーカーをドラッグ中
      onSelectionChange(
        selectionStart,
        Math.max(selectionStart, Math.min(100, position)),
      );
    }
  };

  // マウスアップイベントのハンドラ
  const handleMouseUp = () => {
    setIsDragging(null);
  };

  // マウス移動とマウスアップのイベントリスナーを設定
  useEffect(() => {
    document.addEventListener("mousemove", handleMouseMove);
    document.addEventListener("mouseup", handleMouseUp);
    return () => {
      // コンポーネントのアンマウント時にイベントリスナーを削除
      document.removeEventListener("mousemove", handleMouseMove);
      document.removeEventListener("mouseup", handleMouseUp);
    };
  }, [isDragging]);

  return (
    <div className="mb-8" ref={timelineRef}>
      {/* 現在時間を示すタイムライン */}
      <div className="h-4 bg-gray-200 rounded-full relative">
        <div
          className="absolute h-8 w-8 bg-blue-500 top-1/2 transform -translate-y-1/2 -translate-x-1/2 rounded-full shadow-lg cursor-pointer transition-all duration-200 hover:scale-110"
          style={{ left: `${currentTime}%` }}
          onMouseDown={(e) => handleMouseDown(e, "current")}
        ></div>
      </div>
      {/* 選択範囲を示すタイムライン */}
      <div className="h-4 bg-gray-200 rounded-full relative mt-8">
        {/* 選択範囲の開始マーカー */}
        <div
          className="absolute h-8 w-8 bg-green-500 top-1/2 transform -translate-y-1/2 -translate-x-1/2 rounded-full shadow-lg cursor-pointer transition-all duration-200 hover:scale-110"
          style={{ left: `${selectionStart}%` }}
          onMouseDown={(e) => handleMouseDown(e, "start")}
        ></div>
        {/* 選択範囲の終了マーカー */}
        <div
          className="absolute h-8 w-8 bg-red-500 top-1/2 transform -translate-y-1/2 -translate-x-1/2 rounded-full shadow-lg cursor-pointer transition-all duration-200 hover:scale-110"
          style={{ left: `${selectionEnd}%` }}
          onMouseDown={(e) => handleMouseDown(e, "end")}
        ></div>
        {/* 選択範囲を示すハイライト */}
        <div
          className="absolute h-full bg-blue-200 opacity-50 rounded-full"
          style={{
            left: `${selectionStart}%`,
            width: `${selectionEnd - selectionStart}%`,
          }}
        ></div>
      </div>
    </div>
  );
};

export default Timeline;