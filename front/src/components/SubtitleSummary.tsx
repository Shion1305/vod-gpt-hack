import React from "react";

const SubtitleSummary: React.FC<{
  subtitle: string;
  summary: string;
  onSummarize: () => void;
}> = ({ subtitle, summary, onSummarize }) => {
  return (
    <div className="mt-4">
      <textarea
        className="w-full h-20 mb-4 p-2 border rounded"
        placeholder="字幕がここに表示されます"
        value={subtitle}
        readOnly
      />
      <textarea
        className="w-full h-20 p-2 border rounded"
        placeholder="要約がここに表示されます"
        value={summary}
        readOnly
      />
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