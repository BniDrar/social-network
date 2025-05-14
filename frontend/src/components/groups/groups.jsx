'use client';

import React, { useState, useEffect } from 'react';
import styles from './groups.module.css';
import { GetAllGroups } from '@/services/group';
import Link from 'next/link';

export default function Groups() {
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const res = await GetAllGroups();
        setGroups(res);
      } catch (err) {
        console.error("Failed to fetch groups:", err);
      }
    };
    fetchGroups();
  }, []);

  const joinGroup = async (id) => {
    try {
      const res = await fetch(`${process.env.BACKEND_URL}/api/group/join/request`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ group_id: id }),
        credentials: "include",
      });

      if (res.ok) {
        const croup = document.querySelector(`#group_${id}`);
        if (croup) {
          croup.innerHTML = "request sent";
          croup.classList.add(styles.joined);
        }
      } else {
        console.error("Failed to join group");
      }
    } catch (err) {
      console.error("Error joining group:", err);
    }
  };

  return (
    <div className={styles.groups}>
      <h3>Groups</h3>
      <div className={styles.groupsList}>
        
        {Array.isArray(groups) &&groups?.map((group,index) =>
          group.is_member ? (
            <Link key={`${group.id}_${index}`} href={`/group/${group.id}`} className={styles.group}>
              {group.name}
            </Link>
          ) : (
            <span key={group.id} className={styles.group}>
              {group.name}
              <button
                className={styles.joinButton}
                id={`group_${group.id}`}
                onClick={() => joinGroup(group.id)}
              >
                Join
              </button>
            </span>
          )
        )}
      </div>
    </div>
  );
}
