import React from "react";

export default function Watch({ params }: { params: { id: string } }) {
  return <div>Watch Video: {params.id}</div>;
}
