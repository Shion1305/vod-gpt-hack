import React, { useState } from 'react';
import { MainContainer, ChatContainer, MessageList, Message, MessageInput, TypingIndicator, Avatar } from '@chatscope/chat-ui-kit-react';
import VideoPlayer from './components/VideoPlayer';
import Timeline from './components/Timeline';
import SubtitleSummary from './components/SubtitleSummary';

const App = () => {
  const [messages, setMessages] = useState([]);
  const [currentTime, setCurrentTime] = useState(0);
  const [selectionStart, setSelectionStart] = useState(0);
  const [selectionEnd, setSelectionEnd] = useState(100);
  const [subtitle, setSubtitle] = useState('');
  const [summary, setSummary] = useState('');

  const handleTimeChange = (newTime) => {
    setCurrentTime(newTime);
  };

  const handleSelectionChange = (start, end) => {
    setSelectionStart(start);
    setSelectionEnd(end);
  };

  const handleSummarize = () => {
    setSummary(`${selectionStart}秒から${selectionEnd}秒までの要約がここに表示されます。`);
  };

  const handleSend = (message) => {
    setMessages([...messages, {
      message,
      direction: 'outgoing',
      sender: "user"
    }]);
    
    // AIの応答をシミュレート
    setTimeout(() => {
      setMessages((prevMessages) => [...prevMessages, {
        message: "AIからの応答がここに表示されます。",
        direction: 'incoming',
        sender: "AI"
      }]);
    }, 1000);
  };

  return (
    <div className="flex h-screen">
      <div className="w-1/2 bg-gray-900">
        <MainContainer>
          <ChatContainer>
            <MessageList
              typingIndicator={<TypingIndicator content="AI is thinking" />}
            >
              {messages.map((m, i) => <Message key={i} model={m}>
                <Avatar src={m.direction === 'incoming' ? '/ai-avatar.png' : '/user-avatar.png'} name={m.sender} />
              </Message>)}
            </MessageList>
            <MessageInput placeholder="質問を入力してください..." onSend={handleSend} />
          </ChatContainer>
        </MainContainer>
      </div>
      <div className="w-1/2 bg-pink-100 p-4">
        <VideoPlayer src="https://example.com/sample-video.mp4" currentTime={currentTime} onTimeUpdate={handleTimeChange} />
        <Timeline
          currentTime={currentTime}
          selectionStart={selectionStart}
          selectionEnd={selectionEnd}
          onTimeChange={handleTimeChange}
          onSelectionChange={handleSelectionChange}
        />
        <SubtitleSummary
          subtitle={subtitle}
          summary={summary}
          onSummarize={handleSummarize}
        />
      </div>
    </div>
  );
};

export default App;