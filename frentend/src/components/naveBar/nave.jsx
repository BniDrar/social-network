"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { deleteCookie } from "@/utils/coockie";
import Image from "next/image";
import { BiSearch, BiSun, BiMoon } from "react-icons/bi";
import Notif from "@/components/notifications/notif";
import { useUser } from "@/context/userContext";


import styles from "./nave.module.css";
import { logout } from "@/services/auth";
import Link from "next/link";


export default function NavBar() {
  const router = useRouter();
  const [theme, setTheme] = useState("light");
  const { user, setUser } = useUser();

  useEffect(() => {
    if (user && user.avatar && !user.avatar.startsWith("http")) {
      setUser((prev) => ({
        ...prev,
        avatar: `${process.env.MEDIA_URL}${prev.avatar}`,
      }));
    }

    const savedTheme = localStorage.getItem("theme");
    const systemPrefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    const initialTheme = savedTheme || (systemPrefersDark ? "dark" : "light");
    setTheme(initialTheme);
    document.documentElement.classList.remove("light", "dark");
    document.documentElement.classList.add(initialTheme);
  }, [user?.avatar]);

  const toggleTheme = () => {
    const newTheme = theme === "dark" ? "light" : "dark";
    setTheme(newTheme);
    localStorage.setItem("theme", newTheme);
    document.documentElement.classList.remove("light", "dark");
    document.documentElement.classList.add(newTheme);
  };

  const handleLogout = async () => {
    try {
      // Perform logout actions
      await logout();
      deleteCookie("session");
      localStorage.removeItem("theme");
      router.push("/login");
      setTheme("light");
      document.documentElement.classList.replace("dark", "light");
      return
    } catch (error) {
      console.error("Logout failed:", error);
    }
  };


  const handleProfile = () => {
    router.push("/profile");
  };

  const handleSelect = (e) => {
    const value = e.target.value;
    if (value === "profile") handleProfile();
    else if (value === "logout") handleLogout();
    e.target.value = "profile";
  };

  return (
    <nav className={styles.navbar}>
      <div className={styles.navebarLogo}>
        <Link href="/" className={styles.logoLink}>
          {theme === "light" ? (
            <Image src="/logo-light.svg" alt="logo" width={100} height={70} priority />
          ) : (
            <Image src="/logo-dark.svg" alt="logo" width={100} height={70} priority />
          )}
        </Link>
      </div>

      <div className={styles.navSearch}>
        <form className={styles.searchForm}>
          <input type="text" className={styles.searchInput} placeholder="Search..." />
          <button type="submit" className={styles.searchSubmit}>
            Search
            <BiSearch className={styles.searchIcon} />
          </button>
        </form>
      </div>

      <div className={styles.navActions}>
        <Notif />
        <div className={styles.themeSwitch}>
          <input
            type="checkbox"
            id="theme-toggle"
            className={styles.themeToggle}
            checked={theme === "dark"}
            onChange={toggleTheme}
          />
          <label htmlFor="theme-toggle" className={styles.themeToggleLabel}>
            <BiSun className={`${styles.themeIcon} ${theme === "light" ? styles.themeIconVisible : ""}`} />
            <BiMoon className={`${styles.themeIcon} ${theme === "dark" ? styles.themeIconVisible : ""}`} />
          </label>
        </div>
      </div>

      {user && (
        <div className={styles.profileAvatar}>
          <Image
            className={styles.avatar}
            src={user.avatar || "/default-avatar.jpeg"}
            alt="Avatar"
            width={40}
            height={40}
            priority
            onClick={() => router.push(`/profile/${user.myID}`)}
          />
          <div className={styles.profileDropdown}>
            <select
              onChange={handleSelect}
              defaultValue="profile"
              className={styles.profileSelect}
            >
              <option value="profile">{user.myNickname || `${user.first} ${user.last}`}</option>
              <option value="settings">Settings</option>
              <option value="logout">Logout</option>
            </select>
          </div>
        </div>
      )}
    </nav>
  );
}
