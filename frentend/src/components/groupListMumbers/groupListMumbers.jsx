import React, { useState, useEffect } from 'react';
import Image from 'next/image';
import styles from './gtoupListMumbers.module.css';
import { GetGroupMembers } from '@/services/group';
import Link from 'next/link';

export default function GroupListMembers({ id }) {
  const [groupMembers, setGroupMembers] = useState([]);
  useEffect(() => {
    const fetchGroupMembers = async () => {
      try {
        const fetchedGroupMembers = await GetGroupMembers(id);
        setGroupMembers(fetchedGroupMembers);
      } catch (error) {
        console.error('Failed to fetch group members:', error);
      }
    };
    if (id) fetchGroupMembers();
  }, [id]);

  return (
    <div className={styles.groupListMumbers}>
      <h3>Group Members</h3>
      <ul className={styles.membersList}>
        {groupMembers.map(member => (
          <Link key={member.id} className={styles.memberItem} href={`/profile/${member.id}`}>
            <div className={styles.avatarWrapper}>
              <Image
                src= {`${process.env.MEDIA_URL}${member.avatar}` || '/default-avatar.png'}
                alt={`${member.nickname}'s avatar`}
                width={40}
                height={40}
                className={styles.avatar}
              />
            </div>
            <div className={styles.memberInfo}>
              <span className={styles.nickname}>{member.nickname}</span>
              {member.is_admin && <span className={styles.adminLabel}>Admin</span>}
            </div>
          </Link>
        ))}
      </ul>
    </div>
  );
}

