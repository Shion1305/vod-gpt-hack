"use client";

import React from "react";

// SubtitleSummaryコンポーネントの型定義
const SubtitleSummary: React.FC<{
  subtitle: string;      // 字幕テキスト
  summary: string;       // 要約テキスト
  onSummarize: () => void; // 要約ボタンがクリックされたときのコールバック関数
}> = ({ subtitle, summary, onSummarize }) => {
  return (
    <div className="mt-4">
      {/* 字幕表示エリア */}
      <textarea
        className="w-full h-20 mb-4 p-2 border rounded"
        placeholder="字幕がここに表示されます"
        value={subtitle}
        readOnly  // ユーザーによる編集を防ぐ
      />
      {/* 要約表示エリア */}
      <textarea
        className="w-full h-20 p-2 border rounded"
        placeholder="要約がここに表示されます"
        value={summary}
        readOnly  // ユーザーによる編集を防ぐ
      />
      {/* 要約生成ボタン */}
      <button
        className="mt-4 bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition-colors"
        onClick={onSummarize}
      >
        要約を出す
      </button>
    </div>
  );
};

export default SubtitleSummary;