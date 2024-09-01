"use client";

import { useRouter } from "next/navigation";
import React from "react";
import { Card, CardDescription, CardHeader, CardTitle } from "./ui/card";


const VideoCard = ({vid, s3}: {vid: string, s3: string}) => {
  const router = useRouter();

  const handleOnClick = () => {
    return router.push(`/watch/${vid}?s3=${s3}`);
  };

  return (
    <Card onClick={handleOnClick}>
      <CardHeader>
        <CardTitle>第3回 化学基礎A</CardTitle>
        <CardDescription>電気陰性度</CardDescription>
      </CardHeader>
      <CardTitle></CardTitle>
    </Card>
  );
};

export default VideoCard;
