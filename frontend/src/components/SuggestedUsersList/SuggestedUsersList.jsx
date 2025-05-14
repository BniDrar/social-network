import Link from 'next/link';
import { useState, useEffect } from 'react';
import { GetSuggestedUsers, InviteToJoinGroup } from '@/services/group';
import { BiX } from 'react-icons/bi';
import styles from './SuggestedUsersList.module.css';
import Image from 'next/image';

export default function SuggestedUsersList({ groupId }) {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [isOpen, setIsOpen] = useState(false);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await GetSuggestedUsers(groupId);
        setUsers(response);
      } catch (error) {
        console.error('Error fetching users:', error);
      } finally {
        setLoading(false);
      }
    };
    fetchUsers();
  }, [groupId]);

  const togglePopup = () => setIsOpen(prev => !prev);
  const handleInvitation = async (userId) => {
    try {
       await InviteToJoinGroup({group_id: parseInt(groupId),invited_id: userId,}).then(() => {
        console.log("Invitation sent successfully!");
      });
      setUsers(prevUsers =>
        prevUsers.map(user =>
          user.id === userId ? { ...user, invited: true } : user
        )
      );
    } catch (error) {
      console.error('Error inviting user:', error);
    }
  };

  return (
    <div className={styles.SuggestedUsersList}>
      <button className={styles.InviteFriends} onClick={togglePopup}>
        Invite Members
      </button>

      {isOpen && (
        <div className={styles.SuggestedUsersListContainer}>
            <h3 className={styles.title}>Suggested Users</h3>
          <button className={styles.closeButton} onClick={togglePopup}>
            <BiX size={24} />
          </button>
            {loading ? (
              <p>Loading...</p>
            ) : users.length === 0 ? (
              <p>No suggested users at the moment.</p>
            ) : (
              <ul className={styles.userList}>
                {users.map(user => (
                  <li key={user.id}>
                    <Link href={`/profile/${user.id}`} className={styles.userLink}>
                      <Image
                        width={50}
                        height={50}
                        src={`${process.env.MEDIA_URL}${user.avatar}` || '/default-avatar.jpeg'}
                        alt= "avatar"
                        className={styles.avatar}
                      />
                      <span className={styles.userName}>{`${user.first} ${user.last}`}</span>

                    </Link>

                    <button
                      className={styles.inviteButton}
                      onClick={() => handleInvitation(user.id)}
                      disabled={user.invited}
                    >
                      {user.invited ? 'Invited' : 'Invite'}
                    </button>
                  </li>
                ))}
              </ul>)}
          </div>
      )}
    </div>
  );
}
