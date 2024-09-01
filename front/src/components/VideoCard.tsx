"use client";

import React from "react";
import { Card, CardDescription, CardHeader, CardTitle } from "./ui/card";
import { useRouter } from "next/navigation";

const VideoCard = () => {
  const router = useRouter();

  const handleOnClick = () => {
    return router.push("/watch/adsf");
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
