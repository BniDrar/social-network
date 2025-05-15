'use client';

import React, { useState, useRef, useEffect } from "react";
import styles from "./createPost.module.css";
import Image from "next/image";
import {
  BiImageAdd,
  BiVideo,
  BiPoll,
  BiLocationPlus,
  BiX,
} from "react-icons/bi";
import { createPost } from "@/services/posts";
import { useRouter } from "next/navigation";
import { useUser } from "@/context/userContext";


export default function CreatePost({ groupId }) {
  const popupRef = useRef(null);
  const fileInputRef = useRef(null);
  const [image, setImage] = useState(null);
  const [imagePreview, setImagePreview] = useState(null);
  const [successMessage, setSuccessMessage] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [status, setStatus] = useState(groupId ? 3 : 2);
  const [friends, setFriends] = useState([]);
  const [selectedFriends, setSelectedFriends] = useState([]);
  const [searchQuery, setSearchQuery] = useState("");
  const router = useRouter();
  const { user } = useUser();

  const togglePopup = (show) => {
    if (!popupRef.current) return;
    popupRef.current.classList.toggle(styles.hidden, !show);
    popupRef.current.classList.toggle(styles.active, show);
  };

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result);
      };
      reader.readAsDataURL(file);
      setImage(file);
    }
  };

  const handleFriendSelect = (friendId) => {
    setSelectedFriends((prev) =>
      prev.includes(friendId)
        ? prev.filter((id) => id !== friendId)
        : [...prev, friendId]
    );
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    const content = formData.get("content");
    const group_id = groupId || null;

    const postData = {
      nickname: user.nickname,
      content,
      image,
      status: groupId ? 3 : status,
      group_id,
      allowed_viewers: status === 0 ? selectedFriends : undefined,
    };
    try {
      const response = await createPost(postData);
      if (response.status === 201) {
        form.reset();
        setImage(null);
        setImagePreview(null);
        setSuccessMessage("Post created successfully!");
        setTimeout(() => setSuccessMessage(""), 2000);
        setTimeout(() => togglePopup(false), 1000);
        router.push(`/post/${response.data.post_id}`);
      } else {
        setErrorMessage(response.data.error || "Something went wrong.");
        setTimeout(() => setErrorMessage(""), 5000);
      }
    } catch (err) {
      setErrorMessage("Failed to submit post.");
      setTimeout(() => setErrorMessage(""), 5000);
    }
  };

  return (
    <>
      <div className={styles.createPostContainer}>
        <div className={styles.createPostHeader}>
            <Image
              className={styles.avatar}
              src={user.avatar || "/default-avatar.jpeg"}
              alt="Avatar"
              width={40}
              height={40}
            />
          <button
            type="button"
            className={styles.createPostForm}
            onClick={() => togglePopup(true)}
          >
            <input
              type="text"
              className={styles.createPostInput}
              placeholder="What's on your mind?"
              readOnly
            />
          </button>
        </div>

        <div className={styles.createPostFooter}>
          <div className={styles.createPostActions}>
            <button
              className={styles.createPostActionButton}
              type="button"
              onClick={() => fileInputRef.current.click()}
            >
              <BiImageAdd className={styles.createPostIcon} />
            </button>
            <input
              type="file"
              accept="image/*"
              ref={fileInputRef}
              style={{ display: "none" }}
              onChange={handleImageChange}
            />
          </div>
          {!groupId &&
            <select
              className={styles.createPostSelect}
              value={status}
              onChange={(e) => setStatus(Number(e.target.value))}
            >
              <option value={2}>Public</option>
              <option value={1}>Friends</option>
              <option value={0}>Custom</option>
            </select>
          }
        </div>
      </div>

      {/* Popup */}
      <div ref={popupRef} className={`${styles.createPostPopup} ${styles.hidden}`}>
        <form onSubmit={handleSubmit}>
          <div className={styles.createPosPopuptHeader}>
            <div className={styles.createPostPopupUser}>
              {user.avatar && (
                <Image
                  className={styles.avatar}
                  src={user.avatar}
                  alt="Avatar"
                  width={40}
                  height={40}
                />
              )}
            </div>
            <div className={styles.createPostPopupUserInfo}>
              <h3>{user.nickname}</h3>
              {!groupId &&
                <select
                  name="status"
                  className={styles.postStatus}
                  value={status}
                  onChange={(e) => setStatus(Number(e.target.value))}
                >
                  <option value={2}>Global</option>
                  <option value={1}>Friends</option>
                  <option value={0}>Custom</option>
                </select>
              }
            </div>
            {status === 0 && (
              <div className={styles.customFriendsSection}>
                <input
                  type="text"
                  placeholder="Search friends..."
                  className={styles.friendSearch}
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                />
                <div className={styles.friendList}>
                  {friends
                    .filter(friend =>
                      friend.nickname.toLowerCase().includes(searchQuery.toLowerCase())
                    )
                    .map(friend => (
                      <div
                        key={friend.id}
                        className={`${styles.friendItem} ${selectedFriends.includes(friend.id) ? styles.selected : ""}`}
                        onClick={() => handleFriendSelect(friend.id)}
                      >
                        <Image
                          src={`${process.env.MEDIA_URL}${friend.avatar}` || "/default-avatar.jpeg"}
                          alt="Avatar"
                          width={30}
                          height={30}
                          className={styles.friendAvatar}
                        />
                        <span>{friend.nickname}</span>
                      </div>
                    ))}
                </div>
              </div>
            )}
          </div>

          <div className={styles.createPostPopupContent}>
            <textarea
              name="content"
              rows="12"
              maxLength="500"
              required
              placeholder="What's on your mind?"
            ></textarea>

            {imagePreview && (
              <div className={styles.imagePreviewContainer}>
                <Image
                  width={100}
                  height={100}
                  src={imagePreview}
                  alt="Preview"
                  className={styles.imagePreview}
                  onClick={() => setImagePreview(null)}
                />
                <div className={styles.imageName}>
                  {fileInputRef.current?.files[0]?.name}
                </div>
              </div>
            )}

            <div className={styles.createPostPopupActions}>
              <button
                type="button"
                className={styles.createPostPopupButton}
                onClick={() => fileInputRef.current.click()}
              >
                <BiImageAdd className={styles.createPostPopupIcon} />
                Image
              </button>
            </div>

            {errorMessage && (
              <div className={styles.errorMessage}>{errorMessage}</div>
            )}
            {successMessage && (
              <div className={styles.successMessage}>{successMessage}</div>
            )}

            <button type="submit" className={styles.createPostPopupSubmitButton}>
              Create Post
            </button>
            <button
              type="button"
              className={styles.createPostPopupCancelButton}
              onClick={() => togglePopup(false)}
            >
              <BiX className={styles.createPostPopupIcon} />
            </button>
          </div>
        </form>
      </div>
    </>
  );
}