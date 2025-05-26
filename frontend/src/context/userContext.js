"use client";

import { createContext, useContext, useState, useEffect } from "react";
import { checkLogin, getProfile } from "@/services/auth"


const UserContext = createContext(null);


export function UserProvider({ children }) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const [user, setUser] = useState({
    myID: null,
    myNickname: "",
  });

  useEffect(() => {
    checkLogin().then(setIsLoggedIn);
  }, []);

  useEffect(() => {
    async function fetchUser() {
      if (!isLoggedIn) return;
      try {
        const user = await getProfile();
        if (user) {
          setUser({
            myID: user.id,
            myNickname: user.nickname,
            email: user.email,
            avatar: user.avatar ? `${process.env.MEDIA_URL}${user.avatar}` : "/default-avatar.jpeg",
            first: user.first,
            last: user.last,
            birthday: user.birthday,
            about_me: user.about_me,
            status: user.status,
            profile_owner: user.profile_owner,
            is_following: user.is_following,
            is_followed: user.is_followed,
            followers_count: user.followers_count,
            following_count: user.following_count,
            is_admin: user.is_admin,
            online: user.online,
          });
        }
      } catch (error) {
        console.error("Failed to fetch user profile:", error);
      }
    }
    fetchUser();
  }, [isLoggedIn]);

  return (
    <UserContext.Provider value={{ user, setUser,isLoggedIn, setIsLoggedIn }}>
      {children}
    </UserContext.Provider>
  );
}

export function useUser() {
  return useContext(UserContext);
}