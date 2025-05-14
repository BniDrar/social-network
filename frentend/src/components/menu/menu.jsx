"use client";

import { usePathname } from "next/navigation";
import Link from "next/link";
import { BiHomeSmile, BiUser, BiConversation } from "react-icons/bi";
import styles from "./menu.module.css"; // Adjust according to your actual styles file
import { useState, useEffect } from "react";
import { useWebSocket } from "@/context/wsContext";


const Menu = () => {
  const ws = useWebSocket();
  const [count, setCount] = useState(null);
  const notificationSound =
    typeof window !== "undefined" ? new Audio("/sounds/notif.wav") : null;
  useEffect(() => {
    if (!ws) return;

    const handleMessage = (event) => {
      // PrivateMessage MessageType = iota
      // GroupMessage
      // BroadcastMessage
      // NotificationMessage
      try {
        const notif = JSON.parse(event.data);
        const jsonString = atob(notif.Data);
        const message = JSON.parse(jsonString);

        let add = null;

        switch (notif.Type) {
          case 0: // PrivateMessage
            add = { text: `you have message from ${notif.UserID}` };
            setCount((count) => (count ?? 0) + 1);
            // setNotifs((prevNotifs) => [add, ...prevNotifs]);
            break;
          case 1: // GroupMessage
            add = { text: `you have message from Group ${notif.GroupID}` };
            setCount((count) => (count ?? 0) + 1);
            // setNotifs((prevNotifs) => [add, ...prevNotifs]);
            break;
          case 2: // BroadcastMessage
            // console.log("Broadcast message:", message.text);
            // add = {
            //   text: `you have message broadcasted message from ${notif.UserID}`,
            // };
            // setNotifs((prevNotifs) => [add, ...prevNotifs]);
            break;
          case 3: // NotificationMessage
            // console.log("Notification:", message.text);
            // add = { text: `you have  a request chat from ${notif.UserID}` };
            // setNotifs((prevNotifs) => [add, ...prevNotifs]);
            break;
          default:
            console.warn("Unknown message type:", message.type);
        }

        if (notificationSound) {
          notificationSound.pause();
          notificationSound.currentTime = 0;
          notificationSound
            .play()
            .catch((err) => console.error("Failed to play sound:", err));
        }
      } catch (err) {
        console.error("Failed to parse message:", err);
      }
    };

    ws.addEventListener("message", handleMessage);

    return () => {
      ws.removeEventListener("message", handleMessage);
    };
  }, [ws]);
  /*________________________________________________________________________*/
  const pathname = usePathname();
  const isActive = (path) => (pathname === path ? styles.active : "");
  return (
    <div className={styles.menu}>
      <ul className={styles.menuList}>
        <Link href="/" className={`${styles.listItem} ${isActive("/")}`}>
          <BiHomeSmile className={styles.menuIcon} />
          Feed
        </Link>
        <Link
          href="/profile/0"
          className={`${styles.listItem} ${isActive("/profile")}`}
        >
          <BiUser className={styles.menuIcon} />
          Profile
        </Link>
        <Link
          href="/chat"
          className={`${styles.listItem} ${isActive("/chat")}`}
          onClick={() => {
            setCount(null);
          }}
        >
          <BiConversation className={styles.menuIcon} />
          chat
          <span className={count && styles.notificationCount}>{count}</span>
        </Link>
      </ul>
    </div>
  );
};

export default Menu;
