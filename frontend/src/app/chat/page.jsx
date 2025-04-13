"use client"
// components/SocketComponent.js
import { useEffect } from 'react';


const SocketComponent = () => {
  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8080/api/ws');
    ws.onopen = () => {
      console.log('Connected to WebSocket');
      ws.send('Hello, WebSocket!');
    };
    ws.onmessage = (event) => {
      console.log('Message received:', event.data);
    };
    ws.onclose = () => {
      console.log('WebSocket connection closed');
    };
    return () => {
      ws.close();
    };
  }, []);
  return <div>WebSocket Example</div>;
};

export default SocketComponent;