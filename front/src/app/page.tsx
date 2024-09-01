"use client";

import Header from "@/components/Header";
import VideoCard from "@/components/VideoCard";
import { cache, useEffect, useState } from "react";

export default function Page() {
  const host = "http://localhost:8080";
  const [videos, setVideos] = useState([])

  useEffect(() => {(
    async() => {
      const storedUserId = localStorage.getItem('userId')
      const url = `${host}/api/v1/media/list`
      try{
        const res = await fetch(url, {
          method: 'POST',
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            'userId': storedUserId
          })
        })

        const data = await res.json();
        console.log(data)

        setVideos(data)

      } catch (error) {
        console.log(error)
      }
    }
  )})

  return (
    <div className="h-screen">
      <Header />
      <div className="p-4">
        <VideoCard />
      </div>
      <div className="p-4">
        <VideoCard />
      </div>
      <div className="p-4">
        <VideoCard />
      </div>
      <div className="p-4">
        <VideoCard />
      </div>
    </div>
  )
}
