import { NextResponse } from "next/server";

export async function middleware(request) {
  const session = request.cookies.get("session");

  const isAuthPage = ["/login", "/register"].includes(request.nextUrl.pathname);

  let isLoggedIn = false;

  if (session?.value) {
    try {
      const res = await fetch(`${process.env.BACKEND_URL}/ping/user`, {
        headers: { cookie: `session=${session.value}` },
      });
      isLoggedIn = res.ok;
    } catch (err) {
      isLoggedIn = false;
    }
  }

  if (!isLoggedIn && !isAuthPage) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  if (isLoggedIn && isAuthPage) {
    return NextResponse.redirect(new URL("/", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/", "/group/:path*", "/profile/:path*","/post/:path*", "/event/:path*", "/login", "/register"],
};
