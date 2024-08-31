"use client";
import {
  MainContainer,
  ChatContainer,
  MessageList,
  Message,
  MessageInput,
  TypingIndicator,
  Avatar,
} from "@chatscope/chat-ui-kit-react";
import { MessagePayload } from "@chatscope/chat-ui-kit-react/src/components/Message/Message";
import { useState } from "react";
import "@chatscope/chat-ui-kit-styles/dist/default/styles.min.css";
import "../style.css";
import SubtitleSummary from "../components/SubtitleSummary";
import Timeline from "../components/Timeline";
import VideoPlayer from "../components/VideoPlayer";

const App = () => {
  const [messages, setMessages] = useState<
    {
      message?: string;
      sentTime?: string;
      sender?: string;
      direction: "incoming" | "outgoing" | 0 | 1;
      position: "single" | "first" | "normal" | "last" | 0 | 1 | 2 | 3;
      type?: "html" | "text" | "image" | "custom";
      payload?: MessagePayload;
    }[]
  >([]);
  const [currentTime, setCurrentTime] = useState(0);
  const [selectionStart, setSelectionStart] = useState(0);
  const [selectionEnd, setSelectionEnd] = useState(100);
  const [subtitle, setSubtitle] = useState("");
  const [summary, setSummary] = useState("");

  const handleTimeChange = (newTime: number) => {
    setCurrentTime(newTime);
  };

  const handleSelectionChange = (start: number, end: number) => {
    setSelectionStart(start);
    setSelectionEnd(end);
  };

  const handleSummarize = () => {
    setSummary(
      `${selectionStart}秒から${selectionEnd}秒までの要約がここに表示されます。`,
    );
  };

  const handleSend = (message: string) => {
    setMessages([
      ...messages,
      {
        message,
        direction: "outgoing",
        sender: "user",
        position: "single",
        type: "text",
      },
    ]);

    // AIの応答をシミュレート
    setTimeout(() => {
      setMessages((prevMessages) => [
        ...prevMessages,
        {
          message: "AIからの応答がここに表示されます。",
          direction: "incoming",
          sender: "AI",
          position: "single",
        },
      ]);
    }, 1000);
  };

  return (
    <div className="flex h-screen bg-gray-100">
      <div className="w-1/2 bg-white shadow-md">
        <MainContainer>
          <ChatContainer>
            <MessageList
              typingIndicator={<TypingIndicator content="AI is thinking" />}
            >
              {messages.map((m, i) => (
                <Message key={i} model={m}>
                  <Avatar
                    src={
                      m.direction === "incoming"
                        ? "/ai-avatar.png"
                        : "/user-avatar.png"
                    }
                    name={m.sender}
                  />
                  {m.message}
                </Message>
              ))}
            </MessageList>
            <MessageInput
              placeholder="質問を入力してください..."
              onSend={handleSend}
            />
          </ChatContainer>
        </MainContainer>
      </div>
      <div className="w-1/2 bg-pink-100 p-4 flex flex-col">
        <div className="flex-1 overflow-y-auto flex flex-col">
          <div className="flex-1 mb-4">
            <VideoPlayer
              src="https://example.com/sample-video.mp4"
              currentTime={currentTime}
              onTimeUpdate={handleTimeChange}
            />
          </div>
          <div className="flex-1">
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
      </div>
    </div>
  );
};

export default App;
