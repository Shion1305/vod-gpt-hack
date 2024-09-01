import React, { useState, useCallback } from 'react';

// コンポーネントのプロップスの型定義
interface ResizableLayoutProps {
  leftContent: React.ReactNode;  // 左側に表示するコンテンツ
  rightContent: React.ReactNode; // 右側に表示するコンテンツ
}

const ResizableLayout: React.FC<ResizableLayoutProps> = ({ leftContent, rightContent }) => {
  // 左側のコンテンツの幅を管理するstate（パーセンテージ）
  const [leftWidth, setLeftWidth] = useState(50);
  // ドラッグ中かどうかを管理するstate
  const [isDragging, setIsDragging] = useState(false);

  // マウスボタンが押された時の処理
  const handleMouseDown = useCallback(() => {
    setIsDragging(true);
  }, []);

  // マウスボタンが離された時の処理
  const handleMouseUp = useCallback(() => {
    setIsDragging(false);
  }, []);

  // マウスが動いた時の処理
  const handleMouseMove = useCallback(
    (e: React.MouseEvent) => {
      if (isDragging) {
        // マウスの位置から新しい左側の幅を計算
        const newLeftWidth = (e.clientX / window.innerWidth) * 100;
        // 左側の幅を20%から80%の間に制限
        setLeftWidth(Math.max(20, Math.min(80, newLeftWidth)));
      }
    },
    [isDragging]
  );

  return (
    <div
      className="flex h-screen"
      onMouseMove={handleMouseMove}
      onMouseUp={handleMouseUp}
      onMouseLeave={handleMouseUp} // マウスが要素外に出た場合もドラッグ終了
    >
      {/* 左側のコンテンツ */}
      <div style={{ width: `${leftWidth}%` }} className="overflow-auto">
        {leftContent}
      </div>
      {/* リサイズハンドル */}
      <div
        className="w-1 bg-gray-300 cursor-col-resize"
        onMouseDown={handleMouseDown}
      />
      {/* 右側のコンテンツ */}
      <div style={{ width: `${100 - leftWidth}%` }} className="overflow-auto">
        {rightContent}
      </div>
    </div>
  );
};

export default ResizableLayout;