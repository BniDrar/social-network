import React, { useState, useRef, useEffect } from "react";
import styles from "./NotifPopup.module.css";
import { InvitationResponse, acceptFollowRequest } from "@/services/group";
import Link from "next/link";

export const NotifPopup = ({ onClose, notifs = [], setNotifications }) => {
  const cardRef = useRef();
  const chatBodyRef = useRef();


  // Close the popup when clicking outside
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (cardRef.current && !cardRef.current.contains(event.target)) {
        onClose();
      }
    };

    document.addEventListener("click", handleClickOutside);
    return () => document.removeEventListener("click", handleClickOutside);
  }, [onClose]);

  // Function to remove a notification
  const removeNotification = (id) => {
    setNotifications((prev) => prev.filter((notif) => notif.id !== id));
  };

  return (
    <div
      className={styles["chat-card"]}
      ref={cardRef}
    >
      <div className={styles["chat-body"]} ref={chatBodyRef}>
        {notifs.length > 0 ? (
          notifs.map((msg, index) => (
            <NotificationItem 
              key={msg.id || index} // Using msg.id or fallback to index
              msg={msg} 
              onRemove={removeNotification} // Ensure onRemove is passed correctly
            />
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
  // Check that onRemove is always passed and defined
  if (!onRemove) {
    console.error('onRemove function is not defined');
    return null;
  }

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
  const [loading, setLoading] = useState(false);

  const handleInvitationResponse = async (status) => {
    setLoading(true);
    try {
      await InvitationResponse({
        id: msg.id,
        type: msg.type,
        group_id: msg.group_id,
        sender_id: msg.sender_id,
        accepted: status,
      });
      onRemove(msg.id); // Ensure onRemove is being called properly
    } catch (error) {
      console.error("Error handling invitation response:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={`${styles.message} ${styles.incoming}`}>
      <p>{msg.message}</p>
      <div className={styles.actionButtons}>
        <button
          className={styles.acceptInvitation}
          onClick={() => handleInvitationResponse(true)}
          disabled={loading}
          aria-label="Accept invitation"
        >
          Accept
        </button>
        <button
          className={styles.ignoreInvitation}
          onClick={() => handleInvitationResponse(false)}
          disabled={loading}
          aria-label="Ignore invitation"
        >
          Ignore
        </button>
      </div>
    </div>
  );
};

const FollowReqPrivitProfile = ({ msg, onRemove }) => {
  const [loading, setLoading] = useState(false);

  const handleFollowResponse = async (status) => {
    setLoading(true);
    try {
      await acceptFollowRequest({
        id: msg.id,
        type: msg.type,
        sender_id: msg.sender_id,
        receiver_id: msg.receiver_id,
        accepted: status,
      });
      onRemove(msg.id); // Ensure onRemove is being called properly
    } catch (error) {
      console.error("Error handling follow request:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={`${styles.message} ${styles.incoming}`}>
      <p>{msg.message}</p>
      <div className={styles.actionButtons}>
        <button
          className={styles.acceptInvitation}
          onClick={() => handleFollowResponse(true)}
          disabled={loading}
          aria-label="Accept follow request"
        >
          Accept
        </button>
        <button
          className={styles.ignoreInvitation}
          onClick={() => handleFollowResponse(false)}
          disabled={loading}
          aria-label="Ignore follow request"
        >
          Ignore
        </button>
      </div>
    </div>
  );
};

const EventNotification = ({ msg }) => (
  <div className={`${styles.message} ${styles.incoming}`}>
    <Link href={`/event/${msg.event_id}`} passHref>
      <p>{msg.message}</p>
    </Link>
  </div>
);

export default NotifPopup;
