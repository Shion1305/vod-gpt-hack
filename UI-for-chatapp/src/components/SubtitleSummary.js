import React from 'react';
import { MainContainer, ChatContainer, MessageList, MessageInput } from '@chatscope/chat-ui-kit-react';

const App = () => {
  return (
    <div className="flex h-screen">
      <div className="w-1/2 bg-gray-900">
        <MainContainer>
          <ChatContainer>
            <MessageList>
              {/* メッセージはここに表示されます */}
            </MessageList>
            <MessageInput placeholder="質問を入力してください..." />
          </ChatContainer>
        </MainContainer>
      </div>
      <div className="w-1/2 bg-pink-100 p-4">
        <div className="bg-gray-300 h-64 mb-4">動画再生</div>
        <div className="mb-4">
          <div className="h-2 bg-gray-300 relative">
            <div className="absolute h-4 w-4 bg-red-500 top-1/2 transform -translate-y-1/2" style={{left: '20%'}}></div>
          </div>
        </div>
        <div className="mb-4">
          <div className="h-2 bg-gray-300 relative">
            <div className="absolute h-4 w-4 bg-red-500 top-1/2 transform -translate-y-1/2" style={{left: '30%'}}></div>
            <div className="absolute h-4 w-4 bg-red-500 top-1/2 transform -translate-y-1/2" style={{left: '70%'}}></div>
          </div>
        </div>
        <textarea className="w-full h-20 mb-4 p-2" placeholder="字幕がここに表示されます" readOnly></textarea>
        <textarea className="w-full h-20 p-2" placeholder="要約がここに表示されます" readOnly></textarea>
        <button className="mt-4 bg-blue-500 text-white px-4 py-2 rounded">要約を出す</button>
      </div>
    </div>
  );
};

export default App;