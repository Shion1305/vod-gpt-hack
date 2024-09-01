"use client";

import {
  ChatContainer,
  Message,
  MessageInput,
  MessageList,
  TypingIndicator,
} from "@chatscope/chat-ui-kit-react";
import React, { useCallback, useEffect, useState } from "react";
import "@chatscope/chat-ui-kit-styles/dist/default/styles.min.css";
import { v4 as uuidv4 } from "uuid";
import SubtitleSummary from "@/components/SubtitleSummary";
import Timeline from "@/components/Timeline";
import VideoPlayer from "@/components/VideoPlayer";
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable";
import { Skeleton } from "@/components/ui/skeleton";
import { useSearchParams } from "next/navigation";
import { useParams } from "next/navigation";

const App = () => {
  // チャットメッセージの状態

  const host = "https://vod-gpt.gopher.jp";
  // const host = "http://localhost:8080";
  const [messages, setMessages] = useState<
    {
      message?: string;
      sentTime?: string;
      sender?: string;
      direction: "incoming" | "outgoing" | 0 | 1;
      position: "single" | "first" | "normal" | "last" | 0 | 1 | 2 | 3;
      type?: "html" | "text" | "image" | "custom";
    }[]
  >([
    {
      message: "Hello",
      direction: "incoming",
      position: "single",
    },
    {
      message: "Hello, gpt",
      direction: "outgoing",
      position: "single",
    },
    {
      message: "Bye",
      direction: "incoming",
      position: "single",
    },
  ]);

  // ビデオ関連の状態
  const [currentTime, setCurrentTime] = useState(0); // 現在の再生時間
  const [selectionStart, setSelectionStart] = useState(0); // 選択範囲の開始時間
  const [selectionEnd, setSelectionEnd] = useState(100); // 選択範囲の終了時間
  const [subtitle, setSubtitle] = useState(""); // 字幕テキスト
  const [summary, setSummary] = useState(""); // 要約テキスト
  const [videoSrc, setVideoSrc] = useState<string | null>(null); // ビデオソースのURL
  const [videoDuration, setVideoDuration] = useState(0); // ビデオの総再生時間
  const params = useParams();
  const videoId = params.id;
  const searchParams = useSearchParams();
  console.log(searchParams.get("s3"));

  // UIの状態
  const [isDarkMode, setIsDarkMode] = useState(false); // ダークモードの状態

  // ダークモードの切り替えを処理するエフェクト
  useEffect(() => {
    if (isDarkMode) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  }, [isDarkMode]);

  // ダークモードを切り替える関数
  const toggleDarkMode = () => {
    setIsDarkMode(!isDarkMode);
  };

  // ビデオの現在時間を更新する関数
  const handleTimeChange = (newTime: number) => {
    setCurrentTime(newTime);
  };

  // タイムライン上の選択範囲を更新する関数
  const handleSelectionChange = (start: number, end: number) => {
    setSelectionStart(start);
    setSelectionEnd(end);
  };

  // 選択範囲の要約を生成する非同期関数
  const handleSummarize = async () => {
    try {
      // APIエンドポイントに要約リクエストを送信
      const url = `${host}/api/v1/summary`;
      const postResponse = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          from: selectionStart, // 選択範囲の開始時間（秒）
          to: selectionEnd, // 選択範囲の終了時間（秒）
          vid: videoId, // ビデオの一意識別子
        }),
        redirect: "manual",
      });
      if (!postResponse.ok) {
        console.log("failed to send message");
        console.log(postResponse);
        return;
      }
      const id = await postResponse.json().then((data) => data.id);
      const eventSource = new EventSource(`${host}/api/v1/chat/${id}`);
      let aiResponse = "";
      eventSource.addEventListener("delta", (event) => {
        const data = event.data;
        aiResponse += data;
        setSummary(aiResponse);
      });
    } catch (error) {
      // エラーをコンソールに出力
      console.error("Error generating summary:", error);

      // エラーメッセージを要約状態にセット
      setSummary("申し訳ありませんが、要約の生成中にエラーが発生しました。");
    }
  };

  var aiResponse = "";

  // メッセージを送信し、APIと通信する関数
  const handleSend = useCallback(
    async (message: string) => {
      if (!message.trim()) return; // 空のメッセージを送信しない

      // ユーザーのメッセージをチャットに追加
      setMessages((prevMessages) => [
        ...prevMessages,
        {
          message,
          direction: "outgoing",
          sender: "user",
          position: "single",
          type: "text",
        },
      ]);

      try {
        // APIにリクエストを送信
        const postData = {
          question: message,
          from: selectionStart,
          to: selectionEnd,
          vid: videoId,
        };
        const url = `${host}/api/v1/chat`;
        const postResponse = await fetch(url, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(postData),
          redirect: "manual",
        });
        if (!postResponse.ok) {
          console.log("failed to send message");
          console.log(postResponse);
          return;
        }
        const id = await postResponse.json().then((data) => data.id);
        const prevMessages = messages;
        const eventSource = new EventSource(`${host}/api/v1/chat/${id}`);
        let aiResponse = "";

        eventSource.addEventListener("delta", (event) => {
          const data = event.data;
          aiResponse += data;
          setMessages((_) => [
            ...prevMessages,
            {
              message: aiResponse,
              direction: "incoming",
              sender: "AI",
              position: "single",
            },
          ]);
        });
      } catch (error) {
        console.error("Error sending message:", error);
        // エラーメッセージをチャットに追加
        setMessages((prevMessages) => [
          ...prevMessages,
          {
            message:
              "申し訳ありませんが、エラーが発生しました。もう一度お試しください。",
            direction: "incoming",
            sender: "AI",
            position: "single",
          },
        ]);
      }
    },
    [selectionStart, selectionEnd, videoId],
  );

  // ビデオの長さが変更されたときの処理
  const handleDurationChange = (duration: number) => {
    setVideoDuration(duration);
    setSelectionStart(0);
    setSelectionEnd(duration);
  };

  // キーボードイベントの処理（Enterキーでメッセージを送信）

  return (
    <div className="p-4 h-screen w-screen bg-blue-950">
      <ResizablePanelGroup direction="horizontal" className="rounded-xl">
        <ResizablePanel>
          <ChatContainer className="bg-gray-800 relative z-10">
            <MessageList
              typingIndicator={
                <TypingIndicator
                  content="AI is thinking"
                  className="text-blue-300 text-xl"
                />
              }
              className="bg-gray-800 text-xl"
            >
              {messages.map((m, i) => (
                <Message
                  key={i}
                  model={m}
                  className="bg-gray-700 border-gray-600 text-xl"
                >
                  <Message.Header>
                    {m.direction === "incoming" ? "GPT" : "You"}
                  </Message.Header>
                </Message>
              ))}
            </MessageList>
            <MessageInput
              placeholder="質問を入力してください..."
              onSend={handleSend}
              attachButton={false}
              className="text-gray-100 border-gray-600 text-lg"
            />
          </ChatContainer>
        </ResizablePanel>
        <ResizableHandle withHandle />
        <ResizablePanel>
          <div className="bg-gray-900 p-4 flex flex-col items-center h-full w-full">
            <div className="mb-8 flex items-center justify-center w-full">
              <Skeleton className="h-80 w-[80%] rounded-xl" />
              {videoSrc && (
                <VideoPlayer
                  src={videoSrc}
                  currentTime={currentTime}
                  onTimeUpdate={handleTimeChange}
                  onDurationChange={handleDurationChange}
                />
              )}
            </div>
            <div className="w-[80%]">
              <Timeline
                currentTime={(currentTime / videoDuration) * 100}
                selectionStart={(selectionStart / videoDuration) * 100}
                selectionEnd={(selectionEnd / videoDuration) * 100}
                onTimeChange={(newTime) =>
                  handleTimeChange((newTime / 100) * videoDuration)
                }
                onSelectionChange={(start, end) =>
                  handleSelectionChange(
                    (start / 100) * videoDuration,
                    (end / 100) * videoDuration,
                  )
                }
              />
              <SubtitleSummary
                subtitle={subtitle}
                summary={summary}
                onSummarize={handleSummarize}
              />
            </div>
          </div>
        </ResizablePanel>
      </ResizablePanelGroup>
    </div>
  );
};

export default App;
