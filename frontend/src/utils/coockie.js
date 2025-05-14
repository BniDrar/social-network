'use server'
 
import { cookies } from 'next/headers'
 async function overiteCookie(name, value) {
  (await cookies()).set(name,value, { maxAge: 0 })
}
async function getCookie(name) {
  const cookie = (await cookies()).get(name);
  return cookie.value;
}
async function setCookie(name, value) {
  (await cookies()).set(name, value, { maxAge: 60 * 60 * 24 * 7 });
}
async function deleteCookie(name) {
  (await cookies()).delete(name);
}

export {
  overiteCookie,
  getCookie,
  setCookie,
  deleteCookie,
}