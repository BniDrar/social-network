import React, { useEffect } from "react";
import styles from "./chatBody.module.css"; // Ensure you create the CSS file and import it

import { useUser } from "@/context/userContext";
import { useWebSocket } from "@/context/wsContext";

const ChatBody = ({ id, messages, scroll, chatBodyRef }) => {
  const { user } = useUser();
  const ws = useWebSocket();

  function scrollDown() {
    // scrol down
    setTimeout(() => {
      chatBodyRef.current.scrollTo({
        top: chatBodyRef.current.scrollHeight,
        behavior: "smooth",
      });
    }, 100);
  }
  useEffect(() => {
    scrollDown();
  }, [scroll]);

  return (
    <div className={styles["chat-card"]}>
      {/* Chat Body */}
      <div className={styles["chat-body"]} ref={chatBodyRef}>
        {messages.map((msg, index) => (
          <div
            key={index}
            className={`${styles.message} ${
              //outgoing  incoming
              styles[msg.from == user.myID ? "outgoing" : "incoming"]
              }`}
          >
            <p>{msg.text}</p>
            <span className={styles.timestamp}>
              {new Date(msg.created_at).toLocaleString([], {
                year: "numeric",
                month: "short",
                day: "numeric",
                hour: "2-digit",
                minute: "2-digit",
              })}
            </span>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ChatBody;
