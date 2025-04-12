"use client"
import { useState } from 'react';
import styles from './Navbar.module.css'
import Image from 'next/image';
import Logo from "@/assets/logo.png"
import HomeIcon from '@/assets/icons/home.png'
import GroupsIcon from '@/assets/icons/groups.png'
import EventIcon from '@/assets/icons/event.png'
import FollowersIcon from '@/assets/icons/followers.png'
import ProfileIcon from '@/assets/icons/profile.png'
import LoginIcon from '@/assets/icons/login.png'
import RegisterIcon from '@/assets/icons/pen.png'
import LogoutIcon from '@/assets/icons/logout.png'
import NotifIcon from '@/assets/icons/notification.png'
import Link from 'next/link';

const Navebar = () => {
    const [isAuth, setIsAuth] = useState(true)
    return (
        <header className={styles.navbar}>
            <div className={styles.navbar_start}>
                {/* [logo] */}
                <Link href={'/'} className={styles.link}>
                    <Image src={Logo} className={styles.logo} width={60} height={50} alt='logo' />
                </Link>
            </div>
            <nav className={styles.navbar_middle}>
                {/* [home - groups - followers - events] */}
                <Link href={'/'}>
                    <Image className={styles.pageIcon} src={HomeIcon} width={36} height={36} alt='home icon link' />
                </Link>
                <Link href={'/groups'}>
                    <Image className={styles.pageIcon} src={GroupsIcon} width={36} height={36} alt='groups icon link' />
                </Link>
                <Link href={'/events'}>
                    <Image className={styles.pageIcon} src={EventIcon} width={36} height={36} alt='events icon link' />
                </Link>
                <Link href={'/followers'}>
                    <Image className={styles.pageIcon} src={FollowersIcon} width={36} height={36} alt='followers icon link' />
                </Link>
            </nav>
            <div className={styles.navbar_end}>
                {/* [login - register] OR [profile - logout] */}
                {isAuth ?
                    <>
                        <Link href={'/notification'}>
                            <Image className={styles.pageIcon} src={NotifIcon} width={38} height={38} alt='notification icon link' />
                        </Link>
                        <Link href={`/profile/yrahhaou`}>
                            <Image className={styles.pageIcon} src={ProfileIcon} width={38} height={38} alt='profile icon link' />
                        </Link>
                        <Link href={'/logout'}>
                            <Image className={styles.pageIcon} src={LogoutIcon} width={38} height={38} alt='logout icon link' />
                        </Link>
                    </> :
                    <>
                        <Link href={'/login'}>
                            <Image className={styles.pageIcon} src={LoginIcon} width={38} height={38} alt='login icon link' />
                        </Link>
                        <Link href={'/register'}>
                            <Image className={styles.pageIcon} src={RegisterIcon} width={38} height={38} alt='register icon link' />
                        </Link>
                    </>}
            </div>
        </header>
    );
}

export default Navebar;