import styles from "./footer.module.css";
import Button from "./button/button";

import { useWebSocket } from "@/context/wsContext";
import { useUser } from "@/context/userContext";
import { useRef, useState } from "react"

export default function Footer({ id, setMessage, setMessages, message, messages, scroll, setScroll, isGroup }) {
  const { user } = useUser();
  const ws = useWebSocket();


  const imojiRef = useRef();
  const [showEmojiPicker, setShowEmojiPicker] = useState(false); // Emoji picker state
  const showEmojiPickerRef = useRef();
  showEmojiPickerRef.current = showEmojiPicker;


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
    setScroll(!scroll)
    setMessage(""); // Clear input field
  };
  // Handle emoji click
  const handleEmojiClick = (emoji, event) => {
    event.stopPropagation();
    setMessage((prevMessage) => prevMessage + emoji);
    setShowEmojiPicker(false);
  };

  return (
    <div className={styles.footer}>
      <div className={styles.inputContainer}>
        <input
          className={styles.inputMessage}
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="Type your message..."
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
        <Button onClick={handleSend} />
      </div>
    </div>
  );
}

