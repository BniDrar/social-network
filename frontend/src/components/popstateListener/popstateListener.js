'use client';

import { useEffect } from 'react';
import  {checkLogin} from "@/services/auth"

export default function PopstateListener() {
  useEffect(() => {
    let lastCheck = 0;
    const throttleInterval = 1000;

    const onPopState = () => {
      const now = Date.now();
      if (now - lastCheck > throttleInterval) {
        lastCheck = now;

        // Call async function inside
        checkLogin().then((loggedIn) => {
          if (!loggedIn) {
            window.location.href = '/login';
          }
        });
      }
    };

    window.addEventListener('popstate', onPopState);
    return () => window.removeEventListener('popstate', onPopState);
  }, []);

  return null;
}
