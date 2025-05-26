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
  matcher: [
    // Apply middleware to all paths except:
    // - _next (static files)
    // - favicon
    // - login and register pages
    "/((?!_next/static|_next/image|favicon.ico|login|register).*)",
  ],
};
