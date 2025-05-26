"use client";
import { createContext, useContext, useEffect, useState } from "react";
import { useUser } from "@/context/userContext";

const WebSocketContext = createContext(null);



export function WebSocketProvider({ children }) {
  const { isLoggedIn } = useUser();
  const [ws, setWs] = useState(null);


  useEffect(() => {
    if (!isLoggedIn) return;

    const socket = new WebSocket("ws://localhost:8080/api/ws");
    setWs(socket);
    const bc = new BroadcastChannel("ws");
    socket.onmessage = (event) => {
      bc.postMessage(event.data);
    };
    socket.onerror = (error) => {
      bc.postMessage(error.message);
    };
    socket.onclose = (err) => {
      console.error("WebSocket connection closed:", err);
    };
    return () => {
      socket.close();
      bc.close();
    };
  }, [ isLoggedIn]);

  return <WebSocketContext.Provider value={ws}>{children}</WebSocketContext.Provider>;
}

export function useWebSocket() {
  return useContext(WebSocketContext);
}
