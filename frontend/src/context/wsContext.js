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

    return () => {
      socket.close();
    };
  }, [ isLoggedIn]);

  return <WebSocketContext.Provider value={ws}>{children}</WebSocketContext.Provider>;
}

export function useWebSocket() {
  return useContext(WebSocketContext);
}
