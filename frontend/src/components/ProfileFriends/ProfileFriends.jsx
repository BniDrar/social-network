import Link from 'next/link';
import styles from './ProfileFriends.module.css'

const ProfileFriends = () => {
    return (
        <div className={styles.card}>
            <div className={styles.cardHeader}>
                <div className="">
                    <h3 className="">Friends</h3>
                    <span className="">301 Friends</span>
                </div>
                <Link href={'/'}>See all friends</Link>
            </div>
            <div className={styles.cardBody}></div>
        </div>
    );
}

export default ProfileFriends;