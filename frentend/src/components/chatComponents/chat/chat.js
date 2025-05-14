"use client";

import React, { useState, useRef, useEffect, useCallback } from "react";
import styles from "./chat.module.css";

import { throttle } from "@/utils/throttle";
import { postMessages } from "@/services/chat";

import Header from "./header/header";
import Footer from "./footer/footer";
import ChatBody from "./chatBody/chatBody";
import { getUserInfoClient } from "@/services/profileClient";

import { useWebSocket } from "@/context/wsContext";
import { useUser } from "@/context/userContext";

function Chat({ id }) {
  const { user } = useUser();
  const ws = useWebSocket();

  const chatBodyRef = useRef(); // Add ref for chat body

  const [message, setMessage] = useState(""); // State for input field
  const [messages, setMessages] = useState([]); // Empty initially

  const [scroll, setScroll] = useState(false); // Empty initially


  /*________________________fetch user info___________________________*/
  const [userInfo, setUserInfo] = useState(null);
  useEffect(() => {
    const fetchUser = async () => {
      const info = await getUserInfoClient(id);
      setUserInfo(info)
    };
    fetchUser();
  }, [id]);


  /*________________________fetch more data___________________*/
  const [isFetching, setIsFetching] = useState(false);
  const [hasMore, setHasMore] = useState(true);

  // Throttled fetch more function
  const fetchMore = useCallback(
    throttle(async () => {
      const chatBody = chatBodyRef.current; // Or a ref
      const previousScrollTop = chatBody.scrollTop;
      const previousScrollHeight = chatBody.scrollHeight;

      if (!hasMore || isFetching || messages.length < 15) return;

      setIsFetching(true);
      try {
        const cursor = {
          id: id,
          creation_time: messages[0].created_at,
          last_id: messages[0].from,
          is_group: false,
          limit: 15,
        };

        const data = await postMessages(cursor);
        if (data && data.length > 0) {
          setMessages((prev) => [...data.reverse(), ...prev]);

          requestAnimationFrame(() => {
            chatBody.scrollTop =
              chatBody.scrollHeight - previousScrollHeight + previousScrollTop;
          });
        } else {
          setHasMore(false);
        }
      } catch (err) {
        console.error("Failed to fetch more messages:", err);
      } finally {
        setIsFetching(false);
      }
    }, 300),
    [id, hasMore, isFetching, messages]
  );

  useEffect(() => {
    function handleScroll() {
      if (!chatBodyRef.current || isFetching || !hasMore) return;

      const { scrollTop } = chatBodyRef.current;

      if (scrollTop === 0) {
        fetchMore();
      }
    }

    const chatBody = chatBodyRef.current;
    chatBody?.addEventListener("scroll", handleScroll);

    return () => {
      chatBody?.removeEventListener("scroll", handleScroll);
    };
  }, [fetchMore, isFetching, hasMore]);

  /*________________________Initial fetch of data_______________________*/
  useEffect(() => {
    const cursor = {
      id: id,
      limit: 15,
      is_group: false,
    };
    const fetchMessages = async () => {
      const data = await postMessages(cursor); // or your cursor
      if (data) setMessages(data);
    };
    fetchMessages();
    setScroll(!scroll)
  }, [id]);

  /*________________________ web socket _______________________*/
  useEffect(() => {
    if (!ws) return;

    const handleMessage = (event) => {
      try {
        const message = JSON.parse(event.data);
        const msg64 = message.Data;

        // Safely decode base64 to UTF-8
        const decodedStr = decodeURIComponent(escape(atob(msg64)));

        // Then parse it as JSON
        const decoded = JSON.parse(decodedStr);

        setMessages((prev) => [...prev, decoded]);
        setScroll((prev) => !prev);
      } catch (err) {
        console.error("Failed to parse message:", err);
      }
    };

    ws.addEventListener("message", handleMessage);

    return () => {
      ws.removeEventListener("message", handleMessage);
    };
  }, [ws]);


  /*____________________________return component___________________________*/

  if (!userInfo) return <div>Loading...</div>;

  return (
    <>
      <Header
        id={id}
        first={userInfo.first}
        last={userInfo.last}
        image={userInfo.avatar}
        status={userInfo.online}
        className={styles.header}
      />
      <ChatBody
        id={id} message={message}
        messages={messages}
        className={styles.body}
        scroll={scroll}
        chatBodyRef={chatBodyRef}
      />
      <Footer
        id={id}
        setMessage={setMessage}
        setMessages={setMessages}
        message={message}
        messages={messages}
        className={styles.footer}
        scroll={scroll}
        setScroll={setScroll}
      />
    </>
  );
}

function decodeBase64Utf8(base64) {
  return decodeURIComponent(escape(atob(base64)));
}
export default Chat;
