"use client";

import styles from "./notif.module.css";
import { BiBell } from "react-icons/bi";
import { useEffect, useState } from "react";
import { NotifPopup } from "@/components/notifPopup/NotifPopup"
import { useRef } from "react"
import { GetNotifs } from "@/services/notifs";
import notify from "@/utils/notify";

export default function Notif() {
  const [notifs, setNotifs] = useState([])
  const [count, setCount] = useState(null);

  const [isNotifOpen, setIsNotifOpen] = useState(false);
  const contactRef = useRef(null);
  const bc = new BroadcastChannel("ws");

  useEffect(() => {
    async function fetchNotifs() {
      try {
        const result = await GetNotifs();
        if (result) {
          setNotifs(result);
          if (result.length > 0) {
            setCount(result.length)
          }
        }
      } catch (error) {
        console.error('Error fetching notifications:', error);
      }
    }

    fetchNotifs();
  }, [isNotifOpen]);


  const handleClick = (e) => {
    if (contactRef.current) {
      const rect = contactRef.current.getBoundingClientRect();
      const popupHeight = 285; // Example height (if needed)

      let top = rect.top;
      let right = window.innerWidth - rect.left + 10;

      // Optional: adjust top if popup goes off bottom
      if (top + popupHeight > window.innerHeight) {
        top = window.innerHeight - popupHeight;
        if (top < 0) top = 0; // Don't go above the screen
      }
      setChatPosition({ top, right });
    }
    setIsNotifOpen(true);
  };

  const handleCloseNotif = () => {
    setIsNotifOpen(false);
  };

  useEffect(() => {
    const bc = new BroadcastChannel("ws");

    const handleMessage = (event) => {
      try {
        // Check if the data is a JSON string
        if (typeof event.data === 'string' && event.data.startsWith('{') && event.data.endsWith('}')) {
          const notif = JSON.parse(event.data);
          let decoded
          try {
            const jsonString = atob(notif.Data);
            const message = JSON.parse(jsonString);
            const msg64 = notif.Data;

            // Safely decode base64 to UTF-8
            const decodedStr = decodeURIComponent(escape(atob(msg64)));
  
            decoded = JSON.parse(decodedStr);
          } catch (err) {
            decoded = notif; // Fallback to raw data if parsing fails
            console.log("the notification is:", notif);
          }

          switch (notif.Type) {
            case 0: // PrivateMessage
              // console.log("Private message:", message.text);
              // add = { text: `you have message from ${notif.UserID}` }
              // setNotifs(prevNotifs => [add, ...prevNotifs]);
              break;
            case 1: // GroupMessage
              // console.log("Group message:", message.text);
              // add = { text: `you have message from Group ${notif.GroupID}` }
              // setNotifs(prevNotifs => [add, ...prevNotifs]);
              break;
            case 2: // BroadcastMessage --> leave it
              setNotifs(prevNotifs => [decoded, ...prevNotifs]);
              setCount(count => (count ?? 0) + 1);
              break;
            case 3: // NotificationMessage -->
              setNotifs(prevNotifs => [decoded, ...prevNotifs]);
              setCount(count => (count ?? 0) + 1);
              break;
            case 4: // my message from another tab
              console.log("My message from another tab:", decoded);
              break
            default:
              console.warn("Unknown message type:", message?.type);
          }
          if (notif.Type !== 4) notify("/sounds/notif.wav");
        }

      } catch (err) {
        console.error("Failed to parse message:", err);
      }
    };

    bc.onmessage = handleMessage;

    return () => {
      bc.removeEventListener("message", handleMessage);
      bc.close();
    };
  }, [bc]);

  return (
    <>
      <button
        className={styles.notificationButton}
        onClick={() => {
          handleClick();
          setCount(null);
        }}
      >
        <BiBell className={styles.notificationIcon} />
        <span className={count && styles.notificationCount}>{count}</span>
      </button >
      {/* nitif Popup */}
      {
        isNotifOpen && (
          <NotifPopup
            onClose={handleCloseNotif} // Close function
            notifs={notifs}
          />
        )
      }
    </>
  );
}