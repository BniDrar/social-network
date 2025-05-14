'use client'
import React, { useState } from 'react';
import styles from './createGroup.module.css';
import { BiX, BiListPlus } from "react-icons/bi";
import { CreateGroup } from '@/services/group';
import { useRouter } from 'next/navigation';

const removePopup = () => {
      const popup = document.querySelector(`.${styles.createGroupPopup}`);
      if (popup) {
            popup.classList.remove(styles.active);
            popup.classList.add(styles.hidden);
      }
};

const showPopup = () => {
      const popup = document.querySelector(`.${styles.createGroupPopup}`);
      if (popup) {
            popup.classList.remove(styles.hidden);
            popup.classList.add(styles.active);
      }
};

export default function CreateGroupPopup() {
      const [successMessage, setSuccessMessage] = useState("");
      const [error, setError] = useState("");
      const router = useRouter();
      const handleSubmit = async (e) => {
            e.preventDefault();
            const form = e.target;
            const formData = new FormData(form);
            const name = formData.get("name");
            const description = formData.get("description");
            const type = 0;

            const groupData = {
                  name,
                  description,
                  type,
            };

            try {
                  const res = await CreateGroup(groupData);
                  const errorForm = document.querySelector("#errorForm");
                  if (res.status === 201) {
                        form.reset();
                        setSuccessMessage("Group created successfully!");
                        setTimeout(() => setSuccessMessage(""), 2000);
                        setTimeout(() => {
                              removePopup();
                              router.push(`/group/${res.data.id}`);
                        }, 1000);
                  } else {
                        errorForm.innerHTML = res.error || "An unexpected error occurred.";
                        setTimeout(() => {
                              errorForm.innerHTML = "";
                        }, 5000);
                  }
            } catch (err) {
                  setError("An error occurred while creating the group");
                  setTimeout(() => {
                        setError("");
                  }, 5000);
            }
      };

      return (
            <div className={styles.createGroup}>
                  <button className={styles.createGroupButton} onClick={showPopup}>Create Group <BiListPlus  className={styles.addGroupIcon}/></button>
                  <div className={`${styles.createGroupPopup} ${styles.hidden}`}>
                        <h1>Create Group</h1>
                        <form onSubmit={handleSubmit}>
                              <label htmlFor="name">Group name</label>
                              <input type="text" id="name" name="name" required />
                              <label htmlFor="description">Description</label>
                              <textarea id="description" name="description" required></textarea>
                              <label htmlFor="type">Select group type</label>
                              <button type="submit">Create Group</button>
                              <p id="errorForm" className={styles.errorForm}></p>
                        </form>
                        {successMessage && <p className={styles.successMessage}>{successMessage}</p>}
                        {error && <p className={styles.errorMessage}>{error}</p>}
                        <button className={styles.closeButton} onClick={removePopup}><BiX /></button>
                  </div>
            </div>
      );
}
