import React, { useState, useRef, useEffect, useCallback } from "react";
import styles from "./ChatPopup.module.css"; // Ensure you create the CSS file and import it
import { throttle } from "@/utils/throttle";

import { useWebSocket } from "@/context/wsContext";
import { useUser } from "@/context/userContext";

import { postMessages } from "@/services/chat";

const ChatCard = ({ id, first, last, status, position, onClose , isGroup}) => {
  const { user } = useUser();
  const ws = useWebSocket();

  const cardRef = useRef(); // create a ref for the card
  const chatBodyRef = useRef(); // Add ref for chat body
  const imojiRef = useRef();

  const [message, setMessage] = useState(""); // State for input field
  const [messages, setMessages] = useState([]); // Empty initially
  const [showEmojiPicker, setShowEmojiPicker] = useState(false); // Emoji picker state

  const showEmojiPickerRef = useRef();
  showEmojiPickerRef.current = showEmojiPicker;

  function scrollDown() {
    // scrol down
    setTimeout(() => {
      chatBodyRef.current.scrollTo({
        top: chatBodyRef.current.scrollHeight,
        behavior: "smooth",
      });
    }, 100);
  }
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

  /*___________________________________________________________*/

  useEffect(() => {
    const cursor = {
      id: id,
      is_group: isGroup,
      limit: 15,
    };
    const fetchMessages = async () => {
      const data = await postMessages(cursor); // or your cursor
      if (data) setMessages(data);
    };
    fetchMessages();
    scrollDown();
  }, []);

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
        scrollDown();
      } catch (err) {
        console.error("Failed to parse message:", err);
      }
    };

    ws.addEventListener("message", handleMessage);

    return () => {
      ws.removeEventListener("message", handleMessage);
    };
  }, [ws]);

  const handleSend = () => {
    if (message.trim() === "") return; // Do nothing if message is empty
    const pack = {
      to: id,
      text: message,
      created_at: new Date().toISOString(),
    };
    ws.send(JSON.stringify(pack));

    setMessages([
      ...messages,
      {
        created_at: new Date().toISOString(),
        from: user.myID,
        text: message,
        to: id,
      },
    ]);
    scrollDown();
    setMessage(""); // Clear input field
  };

  // Handle emoji click
  const handleEmojiClick = (emoji, event) => {
    event.stopPropagation();
    setMessage((prevMessage) => prevMessage + emoji);
    setShowEmojiPicker(false);
  };

  useEffect(() => {
    function handleClickAnywhere(event) {
      const clickedOutsideCard =
        cardRef.current && !cardRef.current.contains(event.target);

      const clickedInsideEmoji =
        showEmojiPicker &&
        imojiRef.current &&
        imojiRef.current.contains(event.target);

      if (clickedOutsideCard && !clickedInsideEmoji) {
        onClose();
      }
    }

    document.addEventListener("click", handleClickAnywhere);

    return () => {
      document.removeEventListener("click", handleClickAnywhere);
    };
  }, [onClose, showEmojiPicker]);

  return (
    <div
      className={styles["chat-card"]}
      style={{
        position: "absolute",
        top: position.top - 70,
        right: position.right, // `position` could be 'left' or 'right'
      }}
      ref={cardRef}
    >
      {/* Chat Header */}
      <div className={styles["chat-header"]}>
        <div className={styles.h2}>{`${first} ${last}`}</div>
      </div>

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

      {/* Chat Footer */}
      <div className={styles["chat-footer"]}>
        <input
          type="text"
          placeholder="Type your message"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          className={styles["chat-input"]}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleSend();
            }
          }}
        />
        {/* Emoji Button */}
        <button
          type="button"
          className={styles.emojiButton}
          onClick={() => setShowEmojiPicker(!showEmojiPicker)}
        >
          😊
        </button>
        {/* Emoji Picker */}
        {showEmojiPicker && (
          <div className={styles.emojiPicker} ref={imojiRef}>
            <span onClick={(e) => handleEmojiClick("😀", e)}>😀</span>
            <span onClick={(e) => handleEmojiClick("😁", e)}>😁</span>
            <span onClick={(e) => handleEmojiClick("😂", e)}>😂</span>
            <span onClick={(e) => handleEmojiClick("😍", e)}>😍</span>
            <span onClick={(e) => handleEmojiClick("🤔", e)}>🤔</span>
            <span onClick={(e) => handleEmojiClick("😢", e)}>😢</span>
            <span onClick={(e) => handleEmojiClick("🎉", e)}>🎉</span>
            <span onClick={(e) => handleEmojiClick("🙌", e)}>🙌</span>
          </div>
        )}
        <button onClick={handleSend} className={styles.sendButton}>
          Send
        </button>
      </div>
    </div>
  );
};

function decodeBase64Utf8(base64) {
  return decodeURIComponent(escape(atob(base64)));
}
export default ChatCard;
