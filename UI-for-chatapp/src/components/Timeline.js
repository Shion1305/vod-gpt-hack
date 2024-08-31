import React, { useState, useRef, useEffect } from 'react';

const Timeline = ({ currentTime, selectionStart, selectionEnd, onTimeChange, onSelectionChange }) => {
  const [isDragging, setIsDragging] = useState(null);
  const timelineRef = useRef(null);

  const handleMouseDown = (e, type) => {
    setIsDragging(type);
  };

  const handleMouseMove = (e) => {
    if (!isDragging) return;
    
    const rect = timelineRef.current.getBoundingClientRect();
    const position = (e.clientX - rect.left) / rect.width * 100;
    
    if (isDragging === 'current') {
      onTimeChange(Math.max(0, Math.min(100, position)));
    } else if (isDragging === 'start') {
      onSelectionChange(Math.max(0, Math.min(selectionEnd, position)), selectionEnd);
    } else if (isDragging === 'end') {
      onSelectionChange(selectionStart, Math.max(selectionStart, Math.min(100, position)));
    }
  };

  const handleMouseUp = () => {
    setIsDragging(null);
  };

  useEffect(() => {
    document.addEventListener('mousemove', handleMouseMove);
    document.addEventListener('mouseup', handleMouseUp);
    return () => {
      document.removeEventListener('mousemove', handleMouseMove);
      document.removeEventListener('mouseup', handleMouseUp);
    };
  }, [isDragging]);

  return (
    <div className="mb-4" ref={timelineRef}>
      <div className="h-2 bg-gray-300 relative">
        <div
          className="absolute h-4 w-4 bg-red-500 top-1/2 transform -translate-y-1/2 -translate-x-1/2 cursor-pointer"
          style={{left: `${currentTime}%`}}
          onMouseDown={(e) => handleMouseDown(e, 'current')}
        ></div>
      </div>
      <div className="h-2 bg-gray-300 relative mt-4">
        <div
          className="absolute h-4 w-4 bg-red-500 top-1/2 transform -translate-y-1/2 -translate-x-1/2 cursor-pointer"
          style={{left: `${selectionStart}%`}}
          onMouseDown={(e) => handleMouseDown(e, 'start')}
        ></div>
        <div
          className="absolute h-4 w-4 bg-red-500 top-1/2 transform -translate-y-1/2 -translate-x-1/2 cursor-pointer"
          style={{left: `${selectionEnd}%`}}
          onMouseDown={(e) => handleMouseDown(e, 'end')}
        ></div>
        <div
          className="absolute h-full bg-blue-200 opacity-50"
          style={{left: `${selectionStart}%`, width: `${selectionEnd - selectionStart}%`}}
        ></div>
      </div>
    </div>
  );
};

export default Timeline;