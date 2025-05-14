import React, { useState, useRef, useEffect } from "react";
import styles from "./NotifPopup.module.css";
import { InvitationResponse, acceptFollowRequest } from "@/services/group";
import Link from "next/link";

export const NotifPopup = ({ onClose, notifs = [] }) => {
  const cardRef = useRef();
  const chatBodyRef = useRef();
  const [notifications, setNotifications] = useState(notifs);

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (cardRef.current && !cardRef.current.contains(event.target)) {
        onClose();
      }
    };

    document.addEventListener("click", handleClickOutside);
    return () => document.removeEventListener("click", handleClickOutside);
  }, [onClose]);

  const removeNotification = (id) => {
    setNotifications((prev) => prev.filter((notif) => notif.id !== id));
  };

  return (
    <div
      className={styles["chat-card"]}
      style={{ position: "absolute", top: 70 }}
      ref={cardRef}
    >
      <div className={styles["chat-body"]} ref={chatBodyRef}>
        {notifications.length > 0 ? (
          notifications.map((msg) => (
            <NotificationItem key={msg.id} msg={msg} onRemove={removeNotification} />
          ))
        ) : (
          <div className={`${styles.message} ${styles.incoming}`}>
            <p>No new notifications</p>
          </div>
        )}
      </div>
    </div>
  );
};

const NotificationItem = ({ msg, onRemove }) => {
  switch (msg.type) {
    case 1:
      return <EventNotification msg={msg} />;
    case 2:
    case 3:
      return <GroupReqNotification msg={msg} onRemove={onRemove} />;
    case 4:
      return <FollowReqPrivitProfile msg={msg} onRemove={onRemove} />;
    default:
      return <BasicNotification msg={msg} />;
  }
};

const BasicNotification = ({ msg }) => (
  <div className={`${styles.message} ${styles.incoming}`}>
    <p>{msg.message}</p>
  </div>
);

const GroupReqNotification = ({ msg, onRemove }) => {
  const handleInvitationResponse = async (status) => {
    try {
      await InvitationResponse({
        id: msg.id,
        type: msg.type,
        group_id: msg.group_id,
        sender_id: msg.sender_id,
        accepted: status,
      });
      onRemove(msg.id);
    } catch (error) {
      console.error("Error handling invitation response:", error);
    }
  };

  return (
    <div className={`${styles.message} ${styles.incoming}`}>
      <p>{msg.message}</p>
      <div className={styles.actionButtons}>
        <button
          className={styles.acceptInvitation}
          onClick={() => handleInvitationResponse(true)}
        >
          Accept
        </button>
        <button
          className={styles.ignoreInvitation}
          onClick={() => handleInvitationResponse(false)}
        >
          Ignore
        </button>
      </div>
    </div>
  );
};

const FollowReqPrivitProfile = ({ msg, onRemove }) => {
  const handleFollowResponse = async (status) => {
    try {
      await acceptFollowRequest({
        id: msg.id,
        type: msg.type,
        sender_id: msg.sender_id,
        receiver_id: msg.receiver_id,
        accepted: status,
      });
      onRemove(msg.id);
    } catch (error) {
      console.error("Error handling follow request:", error);
    }
  };

  return (
    <div className={`${styles.message} ${styles.incoming}`}>
      <p>{msg.message}</p>
      <div className={styles.actionButtons}>
        <button
          className={styles.acceptInvitation}
          onClick={() => handleFollowResponse(true)}
        >
          Accept
        </button>
        <button
          className={styles.ignoreInvitation}
          onClick={() => handleFollowResponse(false)}
        >
          Ignore
        </button>
      </div>
    </div>
  );
};

const EventNotification = ({ msg }) => (
  <div className={`${styles.message} ${styles.incoming}`}>
    <Link href={`/event/${msg.event_id}`}>
      <p>{msg.message}</p>
    </Link>
  </div>
);
