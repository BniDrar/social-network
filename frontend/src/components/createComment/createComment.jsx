"use client";
import React, { useState, useRef, useEffect } from "react";
import styles from "./createComment.module.css";
import Image from "next/image";
import { BiImageAdd } from "react-icons/bi";
import { createComment } from "@/services/comment";
import { useUser } from "@/context/userContext";


export default function CreateComment({ postId, setComments, comments }) {
  const fileInputRef = useRef(null);
  const [image, setImage] = useState(null);
  const [imagePreview, setImagePreview] = useState(null);
  const [successMessage, setSuccessMessage] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [content, setContent] = useState("");
  const { user } = useUser();

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

  const handleSubmit = async (e) => {
    e.preventDefault();

    const commentData = {
      content,
      image,
      postId,
    };

    createComment(commentData).then((response) => {
      if (response.status === 201) {
        setContent("");
        setImage(null);
        setImagePreview(null);
        setSuccessMessage("Comment created successfully!");
        setTimeout(() => setSuccessMessage(""), 2000);
        const newComment = {
          ...response.data,
          avatar: user.avatar,
          nickname: user.nickname,
          content,
          created_at: new Date().toLocaleString(),
        }
        setComments(newComment);
      } else {
        setErrorMessage(response.data.error || "An unexpected error occurred.");
        setTimeout(() => {
          setErrorMessage("");
        }, 5000);
      }
    });
  };

  return (
    <div className={styles.createCommentContainer}>
      <div className={styles.createCommentHeader}>
        <Image
          className={styles.avatar}
          src={user.avatar || "/default-avatar.jpeg"}
          alt="Avatar"
          width={30}
          height={30}
        />
        <form className={styles.createCommentForm} onSubmit={handleSubmit}>
          <textarea
            className={styles.createCommentInput}
            placeholder="Write a comment..."
            value={content}
            onChange={(e) => setContent(e.target.value)}
          />
          <div className={styles.createCommentActions}>
            <input
              type="file"
              accept="image/*"
              ref={fileInputRef}
              style={{ display: "none" }}
              onChange={handleImageChange}
            />
            <button
              type="button"
              className={styles.createCommentActionButton}
              onClick={() => fileInputRef.current.click()}
            >
              <BiImageAdd className={styles.createCommentIcon} />
            </button>
            <button type="submit" className={styles.createCommentSubmitButton}>
              Post
            </button>
          </div>
          {successMessage && (
            <div className={styles.successMessage}>{successMessage}</div>
          )}
        </form>
      </div>
      {imagePreview && (
        <div className={styles.imagePreviewContainer}>
          <Image
            width={100}
            height={50}
            src={imagePreview}
            alt="Preview"
            className={styles.imagePreview}
            onClick={() => setImagePreview(null)}
          />
          <span className={styles.ImageInfo}>{image.name}</span>
        </div>
      )}
      {errorMessage && (
        <div className={styles.errorMessage}>{errorMessage}</div>
      )}
    </div>
  );
}