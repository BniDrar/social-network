"use client";

import React, { useState, useRef, useEffect, useCallback } from "react";
import styles from "./chat.module.css";

import { throttle } from "@/utils/throttle";
import { postMessages } from "@/services/chat";

import Header from "./header/header";
import Footer from "./footer/footer";
import ChatBody from "./chatBody/chatBody";
import { getUserInfoClient } from "@/services/profileClient";
import { getGroupInfoClient } from "@/services/profileClient";

function Chat({ id, isGroup }) {
  console.log('id is group', id , isGroup)
  const bc = new BroadcastChannel("ws"); // Create a broadcast channel for WebSocket messages

  const chatBodyRef = useRef(); // Add ref for chat body

  const [message, setMessage] = useState(""); // State for input field
  const [messages, setMessages] = useState([]); // Empty initially

  const [scroll, setScroll] = useState(false); // Empty initially


  /*________________________fetch user info___________________________*/
  const [userInfo, setUserInfo] = useState(null);

  useEffect(() => {
    const fetchUser = async () => {
      let info
      if (!isGroup) {
        info = await getUserInfoClient(id);
        setUserInfo(info)
      } else if (isGroup) {
        info = await getGroupInfoClient(id)
        setUserInfo(info)
      }
    };
    fetchUser();
    // clean chat body on change of id
    return () => {
      setMessages([]);
    };
  }, [id, isGroup]);


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
          is_group: isGroup,
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
      is_group: isGroup,
    };
    const fetchMessages = async () => {
      const data = await postMessages(cursor); // or your cursor
      if (data) setMessages(data);
    };
    fetchMessages();
    setScroll(!scroll)
  }, [id, isGroup]);

  /*________________________ web socket _______________________*/
  useEffect(() => {
    if (!bc) return;
    const handleMessage = (event) => {
      // try {
        console.log("Received message:", event.data.Type);
        let decoded
        const temporaryMessage = JSON.parse(event.data)
        if (temporaryMessage.Type != 1 && temporaryMessage.Type != 2) {
          // decoded = temporaryMessage
          return
        } else {
          const message = JSON.parse(event.data);
          const msg64 = message.Data;

        // Safely decode base64 to UTF-8
          const decodedStr = decodeURIComponent(escape(atob(msg64)));

        // Then parse it as JSON
           decoded = JSON.parse(decodedStr);
        }

        setMessages((prev) => [...prev, decoded]);
        setScroll((prev) => !prev);
      // } catch (err) {
      //   console.error("Failed to parse message:", err);
      // }
    };

    bc.addEventListener("message", handleMessage);

    return () => {
      bc.removeEventListener("message", handleMessage);
    };
  }, [bc]);


  /*____________________________return component___________________________*/


  let name = ""
  if (!isGroup) {
    name = `${userInfo?.first} ${userInfo?.last}`
  } else {
    name = userInfo?.name
  }


  if (!userInfo) return <div>Loading...</div>;

  return (
    <>
      <Header
        id={id}
        name={name}
        image={userInfo.avatar}
        status={userInfo.online}
        className={styles.header}
        isGroup={isGroup}
      />
      {messages.length > 0 ? (
        <ChatBody
          messages={messages}
          chatBodyRef={chatBodyRef}
          isGroup={isGroup}
        />
      ) : (
        <div className={styles.noMessages}>
          <div className={styles.noMessagesText}>
            <p>No messages yet.</p>
            <p>Send a message to start the conversation.</p>
          </div>
        </div>
      )}
      <Footer
        id={id}
        setMessage={setMessage}
        setMessages={setMessages}
        message={message}
        messages={messages}
        className={styles.footer}
        scroll={scroll}
        setScroll={setScroll}
        isGroup={isGroup}
      />
    </>
  );
}

// function decodeBase64Utf8(base64) {
//   return decodeURIComponent(escape(atob(base64)));
// }
export default Chat;
